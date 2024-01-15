// Package yamlstore contains an implementation of config.Store that
// uses YAML files for persistence.
package yamlstore

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/neilotoole/sq/cli/buildinfo"
	"github.com/neilotoole/sq/cli/config"
	"github.com/neilotoole/sq/libsq/core/errz"
	"github.com/neilotoole/sq/libsq/core/ioz"
	"github.com/neilotoole/sq/libsq/core/ioz/lockfile"
	"github.com/neilotoole/sq/libsq/core/lg"
	"github.com/neilotoole/sq/libsq/core/lg/lga"
	"github.com/neilotoole/sq/libsq/core/options"
	"github.com/neilotoole/sq/libsq/source"
)

// Origin of the config path.
// See Store.PathOrigin.
const (
	originFlag    = "flag"
	originEnv     = "env"
	originDefault = "default"
)

var _ config.Store = (*Store)(nil)

// Store provides persistence of config via YAML file.
// It implements config.Store.
type Store struct {
	// Path is the location of the config file
	Path string

	// PathOrigin is one of "flag", "env", or "default".
	PathOrigin string

	// If HookLoad is non-nil, it is invoked by Load
	// on Path's bytes before the YAML is unmarshalled.
	// This allows expansion of variables etc.
	HookLoad func(data []byte) ([]byte, error)

	// ExtPaths holds locations of potential ext config, both dirs and files (with suffix ".sq.yml")
	ExtPaths []string

	// UpgradeRegistry holds upgrade funcs for upgrading the config file.
	UpgradeRegistry UpgradeRegistry

	// OptionsRegistry holds options.
	OptionsRegistry *options.Registry
}

// Lockfile implements Store.Lockfile.
func (fs *Store) Lockfile() (lockfile.Lockfile, error) {
	fp := filepath.Join(filepath.Dir(fs.Path), "config.pid.lock")
	fp, err := filepath.Abs(fp)
	if err != nil {
		return "", errz.Wrap(err, "failed to get abs path for lockfile")
	}
	return lockfile.Lockfile(fp), nil
}

// String returns a log/debug-friendly representation.
func (fs *Store) String() string {
	return fmt.Sprintf("config via %s: %v", fs.PathOrigin, fs.Path)
}

// Location implements Store. It returns the location of the config dir.
func (fs *Store) Location() string {
	return filepath.Dir(fs.Path)
}

// Load reads config from disk. It implements Store.
func (fs *Store) Load(ctx context.Context) (*config.Config, error) {
	log := lg.FromContext(ctx)
	log.Debug("Loading config from file", lga.Path, fs.Path)

	if fs.UpgradeRegistry != nil {
		mightNeedUpgrade, _, err := checkNeedsUpgrade(ctx, fs.Path)
		if err != nil {
			return nil, errz.Wrapf(err, "config: %s", fs.Path)
		}

		if mightNeedUpgrade {
			// The config might need to be upgraded. But, there's an edge case
			// where another process might upgrade the config file before we
			// get a chance to do so. So, we acquire the config lock, and
			// then check again if it still needs upgrade.
			unlock, err := fs.acquireLock(ctx)
			if err != nil {
				return nil, err
			}
			defer unlock()

			// Lock is acquired; check again if config needs upgrade.
			var foundVers string
			mightNeedUpgrade, foundVers, err = checkNeedsUpgrade(ctx, fs.Path)
			if err != nil {
				return nil, errz.Wrapf(err, "config: %s", fs.Path)
			}

			if mightNeedUpgrade {
				log.Info("Upgrade config?", lga.From, foundVers, lga.To, buildinfo.Version)
				if _, err = fs.doUpgrade(ctx, foundVers, buildinfo.Version); err != nil {
					return nil, err
				}

				// We do a cycle of loading and saving the config after the upgrade,
				// because the upgrade may have written YAML via a map, which
				// doesn't preserve order. Loading and saving should fix that.
				cfg, err := fs.doLoad(ctx)
				if err != nil {
					return nil, errz.Wrapf(err, "config: %s: load failed after config upgrade", fs.Path)
				}

				if err = fs.Save(ctx, cfg); err != nil {
					return nil, errz.Wrapf(err, "config: %s: save failed after config upgrade", fs.Path)
				}
			}
		}
	}

	return fs.doLoad(ctx)
}

func (fs *Store) doLoad(ctx context.Context) (*config.Config, error) {
	bytes, err := os.ReadFile(fs.Path)
	if err != nil {
		return nil, errz.Wrapf(err, "config: failed to load file: %s", fs.Path)
	}

	loadHookFn := fs.HookLoad
	if loadHookFn != nil {
		bytes, err = loadHookFn(bytes)
		if err != nil {
			return nil, err
		}
	}

	cfg := config.New()
	if err = ioz.UnmarshallYAML(bytes, cfg); err != nil {
		return nil, errz.Wrapf(err, "config: %s: failed to unmarshal config YAML", fs.Path)
	}

	if cfg.Version == "" {
		cfg.Version = buildinfo.Version
	}

	if cfg.Options == nil {
		cfg.Options = options.Options{}
	}

	if cfg.Options, err = fs.OptionsRegistry.Process(cfg.Options); err != nil {
		return nil, errz.Wrapf(err, "config: %s", fs.Path)
	}

	if cfg.Collection == nil {
		cfg.Collection = &source.Collection{}
	}

	repaired, err := source.VerifyIntegrity(cfg.Collection)
	if err != nil {
		if repaired {
			// The config was repaired. Save the changes.
			err = errz.Combine(err, fs.Save(ctx, cfg))
		}
		return nil, errz.Wrapf(err, "config: %s", fs.Path)
	}

	if err = fs.loadExt(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Save writes config to disk. It implements Store.
func (fs *Store) Save(ctx context.Context, cfg *config.Config) error {
	if fs == nil {
		return errz.New("config file store is nil")
	}

	if err := canonicalizeConfig(fs.OptionsRegistry, cfg); err != nil {
		return err
	}

	data, err := ioz.MarshalYAML(cfg)
	if err != nil {
		return err
	}

	return fs.write(ctx, data)
}

// Write writes the config bytes to disk.
func (fs *Store) write(ctx context.Context, data []byte) error {
	// It's possible that the parent dir of fs.Path doesn't exist.
	if err := ioz.RequireDir(filepath.Dir(fs.Path)); err != nil {
		return errz.Wrapf(err, "failed to make parent dir of config file: %s", filepath.Dir(fs.Path))
	}

	if err := os.WriteFile(fs.Path, data, ioz.RWPerms); err != nil {
		return errz.Wrap(err, "failed to save config file")
	}

	lg.FromContext(ctx).Info("Wrote config file", lga.Path, fs.Path)
	return nil
}

// fileExists returns true if the backing file can be accessed, false if it doesn't
// exist or on any error.
func (fs *Store) fileExists() bool {
	_, err := os.Stat(fs.Path)
	return err == nil
}

// acquireLock acquires the config lock, and returns an unlock func.
func (fs *Store) acquireLock(ctx context.Context) (unlock func(), err error) {
	lock, err := fs.Lockfile()
	if err != nil {
		return nil, errz.Wrap(err, "failed to get config lock")
	}

	// We use the default timeout because config isn't loaded yet,
	// so we don't know what the value is.
	lockTimeout := config.OptConfigLockTimeout.Default()
	if err = lock.Lock(ctx, lockTimeout); err != nil {
		return nil, errz.Wrap(err, "acquire config lock")
	}

	return func() {
		lg.WarnIfFuncError(lg.FromContext(ctx), "Release config lock", lock.Unlock)
	}, nil
}

// canonicalizeConfig checks cfg's validity, and patches cfg to the canonical
// form,cfg's validity. For example, an unknown or nil value in an
// options.Options is deleted.
func canonicalizeConfig(optsReg *options.Registry, cfg *config.Config) error {
	var err error
	if err = config.Valid(cfg); err != nil {
		return err
	}

	cfg.Options, err = optsReg.Process(cfg.Options)
	if err != nil {
		return errz.Wrapf(err, "processing 'config.options'")
	}

	cfg.Options = options.DeleteNil(cfg.Options)
	if cfg.Collection == nil {
		return nil
	}

	if err = cfg.Collection.Visit(func(src *source.Source) error {
		if src.Options, err = optsReg.Process(src.Options); err != nil {
			return errz.Wrapf(err, "processing source options for %s", src.Handle)
		}

		src.Options = options.DeleteNil(src.Options)
		return nil
	}); err != nil {
		return err
	}

	return config.Valid(cfg)
}
