package path

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	envVariable = "PATH"
)

type (
	command string
	path    string

	Path struct {
		commands map[command]path
	}
)

func NewFromDefaultEnv() (Path, error) {
	env := os.Getenv("PATH")
	if env == "" {
		return Path{make(map[command]path, 0)}, nil
	}
	return New(strings.Split(env, ":"))
}

func New(paths []string) (Path, error) {
	res := make(map[command]path, len(paths))
	for _, p := range paths {
		err := filepath.Walk(p, func(_ string, info os.FileInfo, err error) error {
			var pathError *fs.PathError
			if errors.As(err, &pathError) {
				return nil
			}

			if err != nil {
				return err
			}

			res[command(info.Name())] = path(p)
			return nil
		})
		if err != nil {
			return Path{}, err
		}
	}
	return Path{res}, nil
}

func (p Path) FindType(cmd string) (string, error) {
	path, ok := p.commands[command(cmd)]
	if !ok {
		return "", fmt.Errorf("command %s not found", cmd)
	}
	return filepath.Join(string(path), cmd), nil
}
