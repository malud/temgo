package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReplace(t *testing.T) {
	const urlEnvExample = "https://example.com/?query=param"
	os.Setenv("TEMGO_TEST", urlEnvExample)
	defer os.Unsetenv("TEMGO_TEST")

	wd, _ := os.Getwd()
	testCases := []struct {
		in       []byte
		expected string
	}{
		{[]byte("foo bar"), "foo bar"},
		{[]byte("foo {{ NOT_SET }} bar"), "foo {{ NOT_SET }} bar"},
		{[]byte("foo {{invalid }} bar"), "foo {{invalid }} bar"},
		{[]byte("foo {{ PWD }}!"), fmt.Sprintf("foo %s!", wd)},
		{[]byte(`foo "{{ TEMGO_TEST }}"!`), fmt.Sprintf("foo %q!", urlEnvExample)},
	}

	for _, testCase := range testCases {
		r := bytes.NewBuffer(testCase.in)
		w := bytes.NewBuffer(make([]byte, 0))
		rw := bufio.NewReadWriter(bufio.NewReader(r), bufio.NewWriter(w))
		replace(rw, nil)
		out := w.Bytes()
		if diff := cmp.Diff(string(out), testCase.expected); diff != "" {
			t.Errorf("Failed for case: '%v'.\nResult differs:\n%s", string(testCase.in), diff)
		}
	}
}
