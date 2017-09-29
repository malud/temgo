package temgo_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/malud/temgo/temgo"
)

var containsTests = []struct {
	in  string
	out bool
}{
	{"This text contains a {{ VARIABLE }}.", true},
	{"This text contains nothing.", false},
	{"This text contains multiple {{ VAR }}{{ S }}.", true},
	{"This text contains a variable with number {{ AUTH0_DOMAIN }}.", true},
}

var replaceVars = temgo.EnvVars{
	"VAR_A":        "/etc/ptc",
	"VARIABLE_C":   "color:256",
	"S":            "short",
	"AUTH0_DOMAIN": "samples.auth0.com",
}

var replaceTests = []struct {
	in   string
	out  string
	fail bool
}{
	{"This dummy text contains a path to {{ VAR_A }}", "This dummy text contains a path to /etc/ptc", false},
	{"Some text with a {{ VARIABLE_C }} variable!", "Some text with a color:256 variable!", false},
	{"This is a {{ S }} text.", "This is a short text.", false},
	{"This is a {{ S }} text with a {{ FAILVAR}} and some vars [\"{{ VAR_A }}\",'{{ VARIABLE_C }}'].", "This is a short text with a {{ FAILVAR}} and some vars [\"/etc/ptc\",'color:256'].", false},
	{"This text contains a variable with number: '{{ AUTH0_DOMAIN }}'.", "This text contains a variable with number: 'samples.auth0.com'.", false},
	{"This text contains a {{ VAR }} which is {{ NOT_SET }} in environment variables.", "variable VAR is not set\nvariable NOT_SET is not set", true},
}

func TestContainsVariable(t *testing.T) {
	tg := temgo.New(nil, false)
	for _, test := range containsTests {
		r := tg.ContainsVariable([]byte(test.in))
		if diff := cmp.Diff(r, test.out); diff != "" {
			t.Errorf("Failed for case: '%v'.\nResult differs:\n%s", test.in, diff)
		}
	}
}

func TestEnvVars_ReplaceVariables(t *testing.T) {
	tg := temgo.New(replaceVars, false)
	for _, test := range replaceTests {
		bytes, err := tg.ReplaceVariables([]byte(test.in))
		if !test.fail && err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(string(bytes), test.out); diff != "" && !test.fail {
			t.Errorf("Failed for case: '%v'.\nResult differs:\n%s", test.in, diff)
		}
	}
}

func TestEnvVars_ReplaceVariablesStrict(t *testing.T) {
	tg := temgo.New(replaceVars, true)
	for _, test := range replaceTests {
		r, err := tg.ReplaceVariables([]byte(test.in))
		if test.fail && err != nil {
			if diff := cmp.Diff(err.Error(), test.out); diff != "" {
				t.Errorf("Failed for case: '%v'.\nResult differs:\n%s", test.in, diff)
			}
		} else if diff := cmp.Diff(string(r), test.out); diff != "" {
			t.Errorf("Failed for case: '%v'.\nResult differs:\n%s", test.in, diff)
		}
	}
}
