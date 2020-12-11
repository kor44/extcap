package extcap

import "errors"

var (
	// ErrNoInterfaceSpecified is returned when start capture or query configuration options or query supported DLTs
	ErrNoInterfaceSpecified = errors.New("No interface specified")

	// ErrNoPipeProvided is returned when start capture and not provide pipe name to write
	ErrNoPipeProvided = errors.New("No FIFO pipe provided")
)
