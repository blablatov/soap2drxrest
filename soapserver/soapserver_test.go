package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		Bar string
		Doc string
	}{
		{"", "\n"},
		{".", "\t"},
		{"\t", "NaN\null\n\n"},
		{"Data for test", "Number 99999 to data test"},
		{"Yes, no", "No, or, yes, _, ops"},
	}

	var prevBar string
	for _, test := range tests {
		if test.Bar != prevBar {
			fmt.Printf("\n%s\n", test.Bar)
			prevBar = test.Bar
		}
	}

	var prevDoc string
	for _, test := range tests {
		if test.Doc != prevDoc {
			fmt.Printf("\n%s\n", test.Doc)
			prevDoc = test.Doc
		}
	}
}
