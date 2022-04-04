package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		podata string
		msg    string
	}{
		{"", "\n"},
		{".", "\t"},
		{"\t", "NaN\null\n\n"},
		{"Data for test", "Number 99999 to data test"},
		{"Yes, no", "No, or, yes, _, ops"},
	}

	var prevpodata string
	for _, test := range tests {
		if test.podata != prevpodata {
			fmt.Printf("\n%s\n", test.podata)
			prevpodata = test.podata
		}
	}

	var prevmsg string
	for _, test := range tests {
		if test.msg != prevmsg {
			fmt.Printf("\n%s\n", test.msg)
			prevmsg = test.msg
		}
	}
}
