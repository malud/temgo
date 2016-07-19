package main

import (
	"github.com/malud/temgo/temgo"
	"io/ioutil"
	"os"
	"strings"
)

var envVars = make(temgo.EnvVars)

// Initialize and filter all non upper case variables.
func init() {
	for _, e := range os.Environ() {
		string := strings.Split(e, "=")
		if strings.Compare(string[0], strings.ToUpper(string[0])) == 0 {
			envVars[string[0]] = string[1]
		}
	}
}

func main() {
	input := os.Stdin
	bytes, _ := ioutil.ReadAll(input)
	str := string(bytes)
	if temgo.ContainsVariable(bytes) {
		//str, err := temgo.ReplaceVariables(bytes)
	}
	strings.Replace(str, "{{ TERM }}", envVars["TERM"], -1)
	//t.Execute(os.Stdout, data)
}
