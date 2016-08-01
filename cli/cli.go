package main

import (
	"fmt"
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
	var writeErr error
	input := os.Stdin
	bytes, err := ioutil.ReadAll(input)
	if err != nil {
		_, writeErr = os.Stderr.WriteString(fmt.Sprintf("Could not read: %v", err))
	}
	if temgo.ContainsVariable(bytes) {
		str := envVars.ReplaceVariables(bytes)
		_, err = os.Stdout.Write(str)
		if err != nil {
			_, writeErr = os.Stderr.WriteString(fmt.Sprintf("Could not write: %v", err))
		}
	}
	if writeErr != nil {
		panic(writeErr)
	}
}
