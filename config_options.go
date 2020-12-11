package extcap

import (
	"fmt"
	"regexp"
	"strings"
)

// ConfigOption
type ConfigOption interface {
	Call() string
	Display() string
	Tooltip() string
	setNumber(int)
}

// common for all options
type cfg struct {
	number   int
	call     string
	display  string
	tooltip  string
	group    string
	required bool
}

func (c *cfg) Call() string {
	return c.call
}
func (c *cfg) Display() string {
	return c.display
}
func (c *cfg) Tooltip() string {
	return c.tooltip
}

func (c *cfg) string(optType string, params [][2]string) string {
	w := new(strings.Builder)
	fmt.Fprintf(w, "arg {number=%d}{call=--%s}{display=%s}{type=%s}", c.number, c.call, c.display, optType)

	if c.tooltip != "" {
		fmt.Fprintf(w, "{tooltip=%s}", c.tooltip)
	}

	if c.required {
		fmt.Fprintf(w, "{required=true}")
	}

	if c.group != "" {
		fmt.Fprintf(w, "{group=%s", c.group)
	}

	for i := range params {
		fmt.Fprintf(w, "{%s=%s}", params[i][0], params[i][1])
	}

	return w.String()
}

func (c *cfg) setNumber(i int) {
	c.number = i
}

// Integer option
type ConfigIntegerOpt struct {
	cfg
	min          int
	max          int
	defaultValue int

	rangeSet   bool
	defaultSet bool
}

// Create new integer option
func NewConfigIntegerOpt(call, display string) *ConfigIntegerOpt {
	opt := &ConfigIntegerOpt{}
	opt.call = call
	opt.display = display

	return opt
}

// SetTooltip sets option tooltip
func (c *ConfigIntegerOpt) SetTooltip(tooltip string) *ConfigIntegerOpt {
	c.tooltip = tooltip
	return c
}

// WithRange sets min and max value for option
func (c *ConfigIntegerOpt) SetRange(min, max int) *ConfigIntegerOpt {
	if min >= max {
		panic("in range max value should be greater min value")
	}

	c.min = min
	c.max = max

	c.rangeSet = true

	return c
}

// WithDefault sets default value for INTEGER option
func (c *ConfigIntegerOpt) WithDefault(val int) *ConfigIntegerOpt {
	c.defaultValue = val
	c.defaultSet = true
	return c
}

// WithRequired sets option required
func (c *ConfigIntegerOpt) SetRequired(val bool) *ConfigIntegerOpt {
	c.required = val
	return c
}

// WithGroup sets option's group
func (c *ConfigIntegerOpt) WithGroup(group string) *ConfigIntegerOpt {
	c.group = group
	return c
}

// String implement stringer interface
// Example output
//    arg {number=0}{call=--delay}{display=Time delay}{tooltip=Time delay between packages}{type=integer}{range=1,15}{required=true}
func (c *ConfigIntegerOpt) String() string {
	params := [][2]string{}
	if c.rangeSet {
		params = append(params, [2]string{"range", fmt.Sprintf("%d,%d", c.min, c.max)})
	}

	return c.string("integer", params)
}

// ConfigStringOpt impplement ConfigOption interface
type ConfigStringOpt struct {
	cfg
	placeholder  string
	validation   *regexp.Regexp
	required     bool
	defaultValue string
	defaultSet   bool
}

// Create new STRING option
func NewConfigStringOpt(call, display string) *ConfigStringOpt {
	opt := &ConfigStringOpt{}
	opt.call = call
	opt.display = display

	return opt
}

// SetTooltip sets option tooltip
func (c *ConfigStringOpt) SetTooltip(tooltip string) *ConfigStringOpt {
	c.tooltip = tooltip
	return c
}

// SetTooltip sets option tooltip
func (c *ConfigStringOpt) Placeholder(str string) *ConfigStringOpt {
	c.placeholder = str
	return c
}

// Default sets default value for STRING option
func (c *ConfigStringOpt) Default(val string) *ConfigStringOpt {
	c.defaultValue = val
	c.defaultSet = true
	return c
}

// Required sets option required
func (c *ConfigStringOpt) Required(val bool) *ConfigStringOpt {
	c.required = val
	return c
}

// Validation sets option validation
func (c *ConfigStringOpt) Validation(str string) *ConfigStringOpt {
	c.validation = regexp.MustCompile(str)
	return c
}

// String implements string interface
// arg {number=0}{call=--server}{display=IP address for log server}{type=string}{validation=\\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\b}
func (c *ConfigStringOpt) String() string {
	params := [][2]string{}

	if c.placeholder != "" {
		params = append(params, [2]string{"placeholder", c.placeholder})
	}

	if c.validation != nil {
		params = append(params, [2]string{"validation", c.validation.String()})
	}

	return c.string("string", params)
}

// ConfigBoolOpt impplement ConfigOption interface
type ConfigBoolOpt struct {
	cfg
	validation   *regexp.Regexp
	required     bool
	defaultValue bool
	defaultSet   bool
}

// Create new BOOL option
func NewConfigBoolOpt(call, display string) *ConfigBoolOpt {
	opt := &ConfigBoolOpt{}
	opt.call = call
	opt.display = display

	return opt
}

// SetTooltip sets option tooltip
func (c *ConfigBoolOpt) SetTooltip(tooltip string) *ConfigBoolOpt {
	c.tooltip = tooltip
	return c
}

// Default sets default value option
func (c *ConfigBoolOpt) Default(val bool) *ConfigBoolOpt {
	c.defaultValue = val
	c.defaultSet = true
	return c
}

// Required sets option required
func (c *ConfigBoolOpt) Required(val bool) *ConfigBoolOpt {
	c.required = val
	return c
}

// String implements string interface
// arg {number=2}{call=--verify}{display=Verify}{tooltip=Verify package content}{type=boolflag}
func (c *ConfigBoolOpt) String() string {
	params := [][2]string{}

	if c.defaultSet {
		params = append(params, [2]string{"default", fmt.Sprintf("%t", c.defaultValue)})
	}

	return c.string("boolflag", params)
}
