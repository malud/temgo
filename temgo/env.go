package temgo

import (
	"os"
	"strings"
)

type EnvVars map[string]string

func NewEnvVars() EnvVars {
	envVars := make(EnvVars)
	for _, e := range os.Environ() {
		idx := strings.IndexByte(e, '=')
		if idx == -1 {
			continue
		}
		envVars[e[:idx]] = e[idx+1:]
	}
	return envVars
}
