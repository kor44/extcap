package extcap

import (
	"fmt"
	"regexp"
	"strings"
)

// Config represents config option which will be shown in Wireshark GUI
// Output examples
// arg {number=0}{call=--delay}{display=Time delay}{tooltip=Time delay between packages}{type=integer}{range=1,15}{required=true}
// arg {number=1}{call=--message}{display=Message}{tooltip=Package message content}{placeholder=Please enter a message here ...}{type=string}
// arg {number=2}{call=--verify}{display=Verify}{tooltip=Verify package content}{type=boolflag}
// arg {number=3}{call=--remote}{display=Remote Channel}{tooltip=Remote Channel Selector}{type=selector}
// arg {number=4}{call=--server}{display=IP address for log server}{type=string}{validation=\\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\b}
// value {arg=3}{value=if1}{display=Remote1}{default=true}
// value {arg=3}{value=if2}{display=Remote2}{default=false}

// ConfigOption
type ConfigOption interface {
	call() string
	display() string
	tooltip() string
	setNumber(int)
}

// common for all options
type cfg struct {
	number     int
	callValue  string
	displayVal string
	tooltipVal string
	group      string
	required   bool
}

func (c *cfg) call() string {
	return c.callValue
}
func (c *cfg) display() string {
	return c.displayVal
}
func (c *cfg) tooltip() string {
	return c.tooltipVal
}

func (c *cfg) string(optType string, params [][2]string) string {
	w := new(strings.Builder)
	fmt.Fprintf(w, "arg {number=%d}{call=--%s}{display=%s}{type=%s}", c.number, c.callValue, c.displayVal, optType)

	if c.tooltipVal != "" {
		fmt.Fprintf(w, "{tooltip=%s}", c.tooltipVal)
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
	opt.callValue = call
	opt.displayVal = display

	return opt
}

// WithRange sets min and max value for option
func (c *ConfigIntegerOpt) Range(min, max int) *ConfigIntegerOpt {
	if min >= max {
		panic("in range max value should be greater min value")
	}

	c.min = min
	c.max = max

	c.rangeSet = true

	return c
}

// WithDefault sets default value for INTEGER option
func (c *ConfigIntegerOpt) Default(val int) *ConfigIntegerOpt {
	c.defaultValue = val
	c.defaultSet = true
	return c
}

// WithRequired sets option required
func (c *ConfigIntegerOpt) Required(val bool) *ConfigIntegerOpt {
	c.required = val
	return c
}

// WithGroup sets option's group
func (c *ConfigIntegerOpt) Group(group string) *ConfigIntegerOpt {
	c.group = group
	return c
}

// SetTooltip sets option tooltip
func (c *ConfigIntegerOpt) Tooltip(tooltip string) *ConfigIntegerOpt {
	c.tooltipVal = tooltip
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
	opt.callValue = call
	opt.displayVal = display

	return opt
}

// Default sets default value for STRING option
func (c *ConfigStringOpt) Default(val string) *ConfigStringOpt {
	c.defaultValue = val
	c.defaultSet = true
	return c
}

// SetTooltip sets option tooltip
func (c *ConfigStringOpt) Placeholder(str string) *ConfigStringOpt {
	c.placeholder = str
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

// SetTooltip sets option tooltip
func (c *ConfigStringOpt) Tooltip(tooltip string) *ConfigStringOpt {
	c.tooltipVal = tooltip
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
	opt.callValue = call
	opt.displayVal = display

	return opt
}

// Default sets default value option
func (c *ConfigBoolOpt) Default(val bool) *ConfigBoolOpt {
	c.defaultValue = val
	c.defaultSet = true
	return c
}

// SetTooltip sets option tooltip
func (c *ConfigBoolOpt) Tooltip(tooltip string) *ConfigBoolOpt {
	c.tooltipVal = tooltip
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

// Need implement
// fileselect
// selector
// radio
// multicheck
