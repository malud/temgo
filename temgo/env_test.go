package temgo

import (
	"os"
	"reflect"
	"testing"
)

func TestNewEnvVars(t *testing.T) {
	want := EnvVars{
		"TEMGO_TEST_NORMAL":      "someValue",
		"TEMGO_TEST_EMPTY":       "",
		"TEMGO_TEST_MORE_EQUALS": "https://example.com/?query1=param1&query2=param2&query1=param3#fragment",
	}

	for k, v := range want {
		if err := os.Setenv(k, v); err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := os.Unsetenv(k); err != nil {
				t.Fatal(err)
			}
		}()
	}
	expected := NewEnvVars()

	if got := NewEnvVars(); !reflect.DeepEqual(got, expected) {
		t.Errorf("\nNewEnvVars() = %v,\nwant %v", got, expected)
	}
}
