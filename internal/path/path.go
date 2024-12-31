package path

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	envVariable = "PATH"
)

type (
	Path struct {
		paths []string
	}
)

func NewFromDefaultEnv() (Path, error) {
	env := os.Getenv(envVariable)
	if env == "" {
		return Path{make([]string, 0)}, nil
	}
	return New(strings.Split(env, ":"))
}

func New(paths []string) (Path, error) {
	return Path{paths}, nil
}

func (p Path) FindPath(cmd string) (string, error) {
	for _, p := range p.paths {
		fp := filepath.Join(p, cmd)
		if _, err := os.Stat(fp); err == nil {
			return fp, nil
		}
	}
	return "", fmt.Errorf("command %s not found", cmd)
}
