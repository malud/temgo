package temgo

import (
	"regexp"
	"strings"
)

type EnvVars map[string]string

var templatePattern = regexp.MustCompile("{{\\s([A-Z_]*?)\\s}}")

func ContainsVariable(str []byte) bool {
	return templatePattern.Match(str)
}

func replace(e *EnvVars, bytes []byte) []byte {
	list := templatePattern.FindAllStringSubmatch(string(bytes), -1)
	for _, match := range list {
		bytes = []byte(strings.Replace(string(bytes), match[0], (*e)[match[1]], -1))
	}
	return bytes
}

func (e *EnvVars) ReplaceVariables(str []byte) []byte {
	result := templatePattern.ReplaceAllFunc(str, func(b []byte) []byte {
		return replace(e, b)
	})
	return result
}
