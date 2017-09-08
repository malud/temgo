package temgo

import (
	"regexp"
	"strings"
)

type EnvVars map[string]string

// following the posix standards
var templatePattern = regexp.MustCompile(`\{{2} ([a-zA-Z_]+[a-zA-Z_0-9]*) }{2}`)

func ContainsVariable(str []byte) bool {
	return templatePattern.Match(str)
}

func replace(e *EnvVars, bytes []byte) []byte {
	list := templatePattern.FindAllStringSubmatch(string(bytes), -1)
	for _, match := range list {
		str, ok := (*e)[match[1]]
		if !ok {
			continue // variable not set
		}
		bytes = []byte(strings.Replace(string(bytes), match[0], str, -1))
	}
	return bytes
}

func (e *EnvVars) ReplaceVariables(str []byte) []byte {
	result := templatePattern.ReplaceAllFunc(str, func(b []byte) []byte {
		return replace(e, b)
	})
	return result
}
