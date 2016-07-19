package temgo

import (
	"errors"
	"regexp"
)

type EnvVars map[string]string

var templatePattern = regexp.MustCompile("({{\\s[A-Z]+\\s}})")

func ContainsVariable(str []byte) bool {
	return templatePattern.Match(str)
}

func (e *EnvVars) ReplaceVariables(str []byte) ([]byte, error) {
	//templatePattern.ReplaceAllFunc()
	return nil, errors.New("not implemented")
}
