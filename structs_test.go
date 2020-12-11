package extcap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringerInterface(t *testing.T) {
	testCases := []struct {
		name     string
		str      fmt.Stringer
		expected string
	}{
		{"Interface",
			CaptureInterface{"example1", "Example interface 1 for extcap"},
			"interface {value=example1}{display=Example interface 1 for extcap}",
		},

		{"DLT",
			DLT{147, "USER1", "Demo Implementation for Extcap"},
			"dlt {number=147}{name=USER1}{display=Demo Implementation for Extcap}",
		},

		{"Config Integer option",
			NewConfigIntegerOpt("delay", "Time delay").WithRange(1, 15).WithRequired(true).WithTooltip("Time delay between packages"),
			"arg {number=0}{call=--delay}{display=Time delay}{type=integer}{tooltip=Time delay between packages}{range=1,15}{required=true}",
		},

		// {"Config Integer option",
		// 	NewConfigIntegerOpt("delay", "Time delay").Range(1, 15).Required(true).SetTooltip("Time delay between packages"),
		// 	"arg {number=0}{call=--delay}{display=Time delay}{tooltip=Time delay between packages}{type=integer}{range=1,15}{required=true}",
		// },

		{"Config String option",
			NewConfigStringOpt("server", "IP address for log server").Validation("\\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\b"),
			"arg {number=0}{call=--server}{display=IP address for log server}{type=string}{validation=\\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\b}",
		},

		{"Config String option",
			NewConfigStringOpt("message", "Message").SetTooltip("Package message content").Placeholder("Please enter a message here ..."),
			"arg {number=0}{call=--message}{display=Message}{type=string}{tooltip=Package message content}{placeholder=Please enter a message here ...}",
		},

		{"Config Bool option",
			NewConfigBoolOpt("verify", "Verify").SetTooltip("Verify package content"),
			"arg {number=0}{call=--verify}{display=Verify}{tooltip=Verify package content}{type=boolflag}",
		},

		// arg {number=3}{call=--remote}{display=Remote Channel}{tooltip=Remote Channel Selector}{type=selector}
		//
		// value {arg=3}{value=if1}{display=Remote1}{default=true}
		// value {arg=3}{value=if2}{display=Remote2}{default=false}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := fmt.Sprint(tc.str)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
