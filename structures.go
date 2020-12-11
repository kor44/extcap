package extcap

import (
	"fmt"
)

//
type VersionInfo struct {
	Info string
	Help string
}

// Format to string in format
// extcap {version=0.1.0}{help=<some help or URL}
func (ver VersionInfo) String() string {
	return fmt.Sprintf("extcap {version=%s}{help=%s}", ver.Info, ver.Help)
}

// CaptureInterface represents single network interface for capture
type CaptureInterface struct {
	Value   string
	Display string
}

// Format to string in format
// interface {value=example1}{display=Example interface 1 for extcap}
func (iface CaptureInterface) String() string {
	return fmt.Sprintf("interface {value=%s}{display=%s}", iface.Value, iface.Display)
}

// DLT represents link type supported by interface
type DLT struct {
	Number  int
	Name    string
	Display string
}

// Format to string in format
// dlt {number=147}{name=USER1}{display=Demo Implementation for Extcap}
func (dlt DLT) String() string {
	return fmt.Sprintf("dlt {number=%d}{name=%s}{display=%s}", dlt.Number, dlt.Name, dlt.Display)
}

// Config represents config option which will be shown in Wireshark GUI
// Output examples
// arg {number=0}{call=--delay}{display=Time delay}{tooltip=Time delay between packages}{type=integer}{range=1,15}{required=true}
// arg {number=1}{call=--message}{display=Message}{tooltip=Package message content}{placeholder=Please enter a message here ...}{type=string}
// arg {number=2}{call=--verify}{display=Verify}{tooltip=Verify package content}{type=boolflag}
// arg {number=3}{call=--remote}{display=Remote Channel}{tooltip=Remote Channel Selector}{type=selector}
// arg {number=4}{call=--server}{display=IP address for log server}{type=string}{validation=\\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\b}
// value {arg=3}{value=if1}{display=Remote1}{default=true}
// value {arg=3}{value=if2}{display=Remote2}{default=false}
// type Config struct {
// 	Number  int
// 	Call    string
// 	Display string
// 	Tooltip string
// 	Type    string
// }

// // StringConfig impement ConfigOption interface
// type IntegerConfig struct {
// 	cfg
// 	Min     int
// 	Max     int
// 	Default int
// }

// // StringConfig impement ConfigOption interface
// type StringConfig struct {
// 	cfg
// 	Validation *regexp.Regexp
// }

// // boolflag
// type BoolConfig struct {
// 	cfg
// }

// // fileselect
// type FileSelectConfig struct {
// 	cfg
// 	MustExist bool
// }

// // selector
// type SelectorConfig struct {
// 	cfg
// 	Values []interface{}
// }

// // radio

// // multicheck
