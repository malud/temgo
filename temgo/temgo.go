package temgo

import (
	"errors"
	"regexp"
	"strings"
)

type Temgo struct {
	envVars EnvVars
	pattern *regexp.Regexp
	strict  bool
}

func New(envs EnvVars, strict bool) *Temgo {
	return &Temgo{
		envVars: envs,
		// following the posix standards
		pattern: regexp.MustCompile(`{{2} ([a-zA-Z_]+[a-zA-Z_0-9]*) }{2}`),
		strict:  strict,
	}
}

func (t *Temgo) ContainsVariable(str []byte) bool {
	return t.pattern.Match(str)
}

func (t *Temgo) replace(bytes []byte) (result []byte, errList []string) {
	result = bytes
	list := t.pattern.FindAllStringSubmatch(string(bytes), -1)
	for _, match := range list {
		str, ok := t.envVars[match[1]]
		if !ok {
			if t.strict {
				errList = append(errList, "variable "+match[1]+" is not set")
			}
			continue
		}
		result = []byte(strings.Replace(string(result), match[0], str, -1))
	}
	return result, errList
}

func (t *Temgo) ReplaceVariables(str []byte) ([]byte, error) {
	var errList []string
	result := t.pattern.ReplaceAllFunc(str, func(b []byte) []byte {
		r, err := t.replace(b)
		if err != nil {
			errList = append(errList, err...)
		}
		return r
	})
	if len(errList) > 0 {
		return nil, errors.New(strings.Join(errList, "\n"))
	}
	return result, nil
}
