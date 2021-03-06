package main

import (
	"testing"
)

func TestStrPadRight(t *testing.T) {

	type testCase struct {
		Expected string
		Output   string
		Name     string
	}

	cases := []testCase{
		{"", StrPadRight("", 0), "Zero-length no padding"},
		{"          ", StrPadRight("", 10), "Zero-length 10 padding"},
		{"string", StrPadRight("string", 0), "6-length 0 padding"},
	}

	for _, test := range cases {
		if test.Output != test.Expected {
			t.Errorf("Failed case: %s\nExpected %q Found %q", test.Name, test.Expected, test.Output)
		}
	}
}
