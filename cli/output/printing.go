package output

import "github.com/fatih/color"

// Printing describes color and pretty-printing options.
type Printing struct {
	monochrome bool

	// ShowHeader indicates that a header (e.g. a header row) should
	// be printed where applicable.
	ShowHeader bool

	// Verbose indicates that verbose output should be printed where
	// applicable.
	Verbose bool

	// Pretty indicates that output should be pretty-printed.
	// Typically this means indentation, new lines, etc., but
	// varies by output format.
	Pretty bool

	// Indent is the indent string to use when pretty-printing,
	// typically two spaces.
	Indent string

	// Redact indicates that sensitive fields (such as passwords)
	// should be redacted (hidden/masked).
	//
	// TODO: Redact is not being honored by the writers.
	Redact bool

	// Active is the color for an active handle (or group, etc).
	Active *color.Color

	// Bold is the color for bold elements. Frequently Punc will just be color.Bold.
	Bold *color.Color

	// Bool is the color for boolean values.
	Bool *color.Color

	// Bytes is the color for byte / binary values.
	Bytes *color.Color

	// Datetime is the color for time-related values.
	Datetime *color.Color

	// Duration is the color for time duration values.
	Duration *color.Color

	// Error is the color for error elements such as an error message.
	Error *color.Color

	// Faint is the color for faint elements - the opposite of Hilite.
	Faint *color.Color

	// Handle is the color for source handles such as "@sakila"
	Handle *color.Color

	// Header is the color for header elements in a table.
	Header *color.Color

	// Hilite is the color for highlighted elements.
	Hilite *color.Color

	// Key is the color for keys such as a JSON field name.
	Key *color.Color

	// Location is the color for Source.Location values.
	Location *color.Color

	// Null is the color for null.
	Null *color.Color

	// Normal is the default color.
	Normal *color.Color

	// Number is the color for number values, including int, float, decimal etc.
	Number *color.Color

	// Punc is the color for punctuation such as colons, braces, etc.
	Punc *color.Color

	// String is the color for string values.
	String *color.Color

	// Success is the color for success elements.
	Success *color.Color
}

// NewPrinting returns a Printing instance. Color and pretty-print
// are enabled. The default indent is two spaces.
func NewPrinting() *Printing {
	pr := &Printing{
		ShowHeader: true,
		Verbose:    false,
		Pretty:     true,
		Redact:     true,
		monochrome: false,
		Indent:     "  ",
		Active:     color.New(color.FgGreen, color.Bold),
		Bold:       color.New(color.Bold),
		Bool:       color.New(color.FgYellow),
		Bytes:      color.New(color.Faint),
		Datetime:   color.New(color.FgGreen, color.Faint),
		Duration:   color.New(color.FgGreen, color.Faint),
		Error:      color.New(color.FgRed, color.Bold),
		Faint:      color.New(color.Faint),
		Handle:     color.New(color.FgBlue),
		Header:     color.New(color.FgBlue, color.Bold),
		Hilite:     color.New(color.FgHiBlue),
		Key:        color.New(color.FgBlue, color.Bold),
		Location:   color.New(color.FgGreen),
		Normal:     color.New(),
		Null:       color.New(color.Faint),
		Number:     color.New(color.FgCyan),
		Punc:       color.New(color.Bold),
		String:     color.New(color.FgGreen),
		Success:    color.New(color.FgGreen, color.Bold),
	}

	pr.EnableColor(true)
	return pr
}

func (pr *Printing) colors() []*color.Color {
	return []*color.Color{
		pr.Active, pr.Bold, pr.Bold, pr.Bytes, pr.Datetime, pr.Duration,
		pr.Error, pr.Faint, pr.Handle, pr.Header, pr.Hilite,
		pr.Key, pr.Location, pr.Normal, pr.Null, pr.Number,
		pr.Punc, pr.String, pr.Success,
	}
}

// IsMonochrome returns true if in monochrome (no color) mode.
// Default is false (color enabled) for a new instance.
func (pr *Printing) IsMonochrome() bool {
	return pr.monochrome
}

// EnableColor enables or disables all colors.
func (pr *Printing) EnableColor(enable bool) {
	if enable {
		pr.monochrome = false

		for _, clr := range pr.colors() {
			clr.EnableColor()
		}
		return
	}

	pr.monochrome = true
	for _, clr := range pr.colors() {
		clr.DisableColor()
	}
}
