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
