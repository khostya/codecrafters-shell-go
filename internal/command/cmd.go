package command

import (
	"errors"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/path"
	"os"
	"strconv"
	"strings"
)

const (
	typeCommand = "type"
	exitCommand = "exit"
	echoCommand = "echo"
)

type Commands struct {
	path path.Path
}

func NewCommands() (*Commands, error) {
	env, err := path.NewFromDefaultEnv()
	if err != nil {
		return nil, err
	}
	return &Commands{env}, nil
}

func (c Commands) typeCMD(command []string) (string, error) {
	if len(command) <= 1 {
		return "", errors.New(strings.Join(command, " ") + ": command not found")
	}
	switch command[1] {
	case exitCommand, typeCommand, echoCommand:
		return fmt.Sprintf("%s is a shell builtin", command[1]), nil
	default:
		path, err := c.path.FindType(command[1])
		if err != nil {
			return "", errors.New(strings.Join(command[1:], " ") + ": not found")
		}
		return fmt.Sprintf("%s is %s", command[1], path), nil
	}
}

func (c Commands) exitCMD(command []string) error {
	if len(command) <= 1 {
		return errors.New(strings.Join(command, " ") + ": command not found")
	}
	i, err := strconv.Atoi(command[1])
	if err != nil {
		return err
	}
	os.Exit(i)
	return nil
}

func (c Commands) echoCMD(command []string) (string, error) {
	if len(command) < 1 {
		return "", errors.New(strings.Join(command, " ") + ": command not found")
	}
	return strings.Join(command[1:], " "), nil
}

func (c Commands) Eval(command []string) (string, error) {
	switch command[0] {
	case exitCommand:
		return "", c.exitCMD(command)
	case echoCommand:
		return c.echoCMD(command)
	case typeCommand:
		return c.typeCMD(command)
	default:
		return "", errors.New(strings.Join(command, " ") + ": command not found")
	}
}
