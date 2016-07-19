package temgo_test

import (
	"github.com/malud/temgo/temgo"
	"testing"
)

var containsTests = []struct {
	in  string
	out bool
}{
	{"This text contains a {{ VARIABLE }}.", true},
	{"This text contains nothing.", false},
}

var replaceVars = &temgo.EnvVars{
	"VAR_A":      "/etc/ptc",
	"VARIABLE_C": "color:256",
	"S":          "short",
}

var replaceTests = []struct {
	in  string
	out string
}{
	{"This dummy text contains a path to {{ VAR_A }}", "This dummy text contains a path to /etc/ptc"},
	{"Some text with a {{ VARIABLE_C }} variable!", "Some text with a color:256 variable!"},
	{"This is a {{ S }} text.", "This is a short text."},
}

func TestContainsVariable(t *testing.T) {
	for _, test := range containsTests {
		r := temgo.ContainsVariable([]byte(test.in))
		if r != test.out {
			t.Errorf("Failed for case: '%v'. Expected: %v Got: %v", test.in, test.out, r)
		}
	}
}

func TestEnvVars_ReplaceVariables(t *testing.T) {
	for _, test := range replaceTests {
		bytes, err := replaceVars.ReplaceVariables([]byte(test.in))
		if string(bytes) != test.out || err != nil {
			t.Errorf("Failed for case: '%v'. Expected: %v Got: %v", test.in, test.out, string(bytes))
		}
	}
}
