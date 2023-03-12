package source

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/neilotoole/sq/libsq/core/errz"
	"github.com/neilotoole/sq/libsq/core/stringz"
)

const (
	msgUnknownSrc  = `unknown data source %s`
	msgNoActiveSrc = "no active data source"
)

// Set is a set of sources. Typically it is loaded from config
// at a start of a run.
type Set struct {
	mu   sync.Mutex
	data setData
}

// setData holds Set's for the purposes of serialization
// to YAML etc. (we don't want to expose setData's exported
// fields directly on Set.)
//
// This seemed like a good idea t the time, but probably wasn't.
type setData struct {
	ActiveSrc  string    `yaml:"active" json:"active"`
	ScratchSrc string    `yaml:"scratch" json:"scratch"`
	Items      []*Source `yaml:"items" json:"items"`
}

// Data returns the internal representation of the set data.
// This is a filthy hack so that the internal data can be passed
// directly to sq's colorizing json encoder (it can't handle colorization
// of values that implement json.Marshaler).
//
// There are two long-term solutions here:
//  1. The color encoder needs to be able to handle json.RawMessage.
//  2. Refactor source.Set so that it doesn't have this weird internal
//     representation.
func (s *Set) Data() any {
	if s == nil {
		return nil
	}

	return s.data
}

// MarshalJSON implements json.Marshaler.
func (s *Set) MarshalJSON() ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return json.Marshal(s.data)
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *Set) UnmarshalJSON(b []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return json.Unmarshal(b, &s.data)
}

// MarshalYAML implements yaml.Marshaler.
func (s *Set) MarshalYAML() (any, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data, nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (s *Set) UnmarshalYAML(unmarshal func(any) error) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return unmarshal(&s.data)
}

// Items returns the sources as a slice.
func (s *Set) Items() []*Source {
	return s.data.Items
}

func (s *Set) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	return stringz.SprintJSON(s)
}

// Add adds src to s.
func (s *Set) Add(src *Source) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if i, _ := s.indexOf(src.Handle); i != -1 {
		return errz.Errorf("data source with name %s already exists", src.Handle)
	}

	s.data.Items = append(s.data.Items, src)
	return nil
}

// Exists returns true if handle already exists loc the set.
func (s *Set) Exists(handle string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	i, _ := s.indexOf(handle)
	return i != -1
}

func (s *Set) indexOf(handle string) (int, *Source) {
	for i, src := range s.data.Items {
		if src.Handle == handle {
			return i, src
		}
	}

	return -1, nil
}

// Active returns the active source, or nil.
func (s *Set) Active() *Source {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.active()
}

func (s *Set) active() *Source {
	if s.data.ActiveSrc == "" {
		return nil
	}

	i, src := s.indexOf(s.data.ActiveSrc)
	if i == -1 {
		return nil
	}

	return src
}

// Scratch returns the scratch source, or nil.
func (s *Set) Scratch() *Source {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.data.ScratchSrc == "" {
		return nil
	}

	i, src := s.indexOf(s.data.ScratchSrc)
	if i == -1 {
		return nil
	}

	return src
}

// Get gets the src with handle, or returns an error.
func (s *Set) Get(handle string) (*Source, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	handle = strings.TrimSpace(handle)
	if handle == "" {
		return nil, errz.Errorf(msgUnknownSrc, handle)
	}

	if !strings.HasPrefix(handle, "@") {
		handle = "@" + handle
	}

	// Special handling for "@active", which is the reserved
	// handle for the active source.
	if handle == ActiveHandle {
		activeSrc := s.active()
		if activeSrc == nil {
			return nil, errz.New(msgNoActiveSrc)
		}
		return activeSrc, nil
	}

	i, src := s.indexOf(handle)
	if i == -1 {
		return nil, errz.Errorf(msgUnknownSrc, handle)
	}

	return src, nil
}

// SetActive sets the active src, or unsets any active
// src if handle is empty (and thus returns nil,nil).
// If handle does not exist, an error is returned.
func (s *Set) SetActive(handle string) (*Source, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handle == "" {
		s.data.ActiveSrc = ""
		return nil, nil //nolint:nilnil
	}

	for _, src := range s.data.Items {
		if src.Handle == handle {
			s.data.ActiveSrc = handle
			return src, nil
		}
	}

	return nil, errz.Errorf(msgUnknownSrc, handle)
}

// SetScratch sets the scratch src to handle. If handle
// is empty string, the scratch src is unset, and nil,nil
// is returned.
func (s *Set) SetScratch(handle string) (*Source, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handle == "" {
		s.data.ScratchSrc = ""
		return nil, nil //nolint:nilnil
	}
	for _, src := range s.data.Items {
		if src.Handle == handle {
			s.data.ScratchSrc = handle
			return src, nil
		}
	}

	return nil, errz.Errorf(msgUnknownSrc, handle)
}

// Remove removes from the set the src having handle.
func (s *Set) Remove(handle string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.data.Items) == 0 {
		return errz.Errorf(msgUnknownSrc, handle)
	}

	i, _ := s.indexOf(handle)
	if i == -1 {
		return errz.Errorf(msgUnknownSrc, handle)
	}

	if s.data.ActiveSrc == handle {
		s.data.ActiveSrc = ""
	}

	if s.data.ScratchSrc == handle {
		s.data.ScratchSrc = ""
	}

	if len(s.data.Items) == 1 {
		s.data.Items = s.data.Items[0:0]
		return nil
	}

	pre := s.data.Items[:i]
	post := s.data.Items[i+1:]

	s.data.Items = pre
	s.data.Items = append(s.data.Items, post...)
	return nil
}

// Handles returns the set of source handles.
func (s *Set) Handles() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	handles := make([]string, len(s.data.Items))
	for i := range s.data.Items {
		handles[i] = s.data.Items[i].Handle
	}

	return handles
}

// Clone returns a deep copy of s. If s is nil, nil is returned.
func (s *Set) Clone() *Set {
	if s == nil {
		return nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	data := setData{
		ActiveSrc:  s.data.ActiveSrc,
		ScratchSrc: s.data.ScratchSrc,
		Items:      make([]*Source, len(s.data.Items)),
	}

	for i, src := range s.data.Items {
		data.Items[i] = src.Clone()
	}

	return &Set{
		mu:   sync.Mutex{},
		data: data,
	}
}

// VerifySetIntegrity verifies the internal state of s.
// Typically this func is invoked after s has been loaded
// from config, verifying that the config is not corrupt.
// If err is returned non-nil, repaired may be returned true
// to indicate that ss has been repaired and modified. The
// caller should save the config to persist the repair.
func VerifySetIntegrity(ss *Set) (repaired bool, err error) {
	if ss == nil {
		return false, errz.New("source set is nil")
	}

	ss.mu.Lock()
	defer ss.mu.Unlock()

	handles := map[string]*Source{}
	for i := range ss.data.Items {
		src := ss.data.Items[i]
		if src == nil {
			return false, errz.Errorf("source set item %d is nil", i)
		}

		err := verifyLegalSource(src)
		if err != nil {
			return false, errz.Wrapf(err, "source set item %d", i)
		}

		if _, exists := handles[src.Handle]; exists {
			return false, errz.Errorf("source set item %d duplicates handle %s", i, src.Handle)
		}

		handles[src.Handle] = src
	}

	if strings.TrimSpace(ss.data.ActiveSrc) != "" {
		if _, exists := handles[ss.data.ActiveSrc]; !exists {
			// The active source doesn't exist. We'll unset the active source.
			activeSrc := ss.data.ActiveSrc
			ss.data.ActiveSrc = ""
			repaired = true
			// Note that the caller will still need to save the source set
			// to the config file for the repair to take effect.
			return repaired, errz.Errorf("active source {%s} does not exist in source set: config has been repaired: please try again", activeSrc) //nolint:lll
		}
	}

	return repaired, nil
}

// verifyLegalSource performs basic checking on source s.
func verifyLegalSource(s *Source) error {
	if s == nil {
		return errz.New("source is nil")
	}

	err := VerifyLegalHandle(s.Handle)
	if err != nil {
		return err
	}

	if strings.TrimSpace(s.Location) == "" {
		return errz.New("source location is empty")
	}

	if s.Type == TypeNone {
		return errz.Errorf("source type is empty or unknown: %q", s.Type)
	}

	return nil
}
