package cli

import (
	"github.com/neilotoole/sq/cli/config/yamlstore"
	"github.com/neilotoole/sq/cli/flag"
	"github.com/spf13/cobra"
)

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Args:  cobra.NoArgs,
		Short: "Manage config",
		Long:  `View and edit config.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		Example: `  # Print config location
  $ sq config location

  # Show base config
  $ sq config get

  # Show base config including unset and default values.
  $ sq config get -v

  # Set option value
  $ sq config set format json

  # Help for an option
  $ sq config set format --help

  # Edit base config
  $ sq config edit

  # Edit config for source
  $ sq config edit @sakila

  # Delete option (reset to default value)
  $ sq config set -D log.level`,
	}

	return cmd
}

func newConfigLocationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "location",
		Aliases: []string{"loc"},
		Short:   "Print config location",
		Long:    "Print config location. Use --verbose for more detail.",
		Args:    cobra.ExactArgs(0),
		RunE:    execConfigLocation,
		Example: `  # Print config location
  $ sq config location
  /Users/neilotoole/.config/sq

  # Print location, also show origin (flag, env, default)
  $ sq config location -v
  /Users/neilotoole/.config/sq
  Origin: env`,
	}

	cmd.Flags().BoolP(flag.JSON, flag.JSONShort, false, flag.JSONUsage)
	cmd.Flags().BoolP(flag.YAML, flag.YAMLShort, false, flag.YAMLUsage)
	return cmd
}

func execConfigLocation(cmd *cobra.Command, _ []string) error {
	rc := RunContextFrom(cmd.Context())
	path := rc.ConfigStore.Location()
	var origin string
	if store, ok := rc.ConfigStore.(*yamlstore.Store); ok {
		origin = store.PathOrigin
	}

	return rc.writers.configw.Location(path, origin)
}
