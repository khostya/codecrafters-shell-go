package command

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/model"
	"github.com/codecrafters-io/shell-starter-go/internal/path"
	"os"
)

const (
	typeCommand = "type"
	exitCommand = "exit"
	echoCommand = "echo"
	pwdCommand  = "pwd"
	cdCommand   = "cd"
)

type Commands struct {
	path path.Path
	home string
}

func NewCommands() (*Commands, error) {
	env, err := path.NewFromDefaultEnv()
	if err != nil {
		return nil, err
	}
	home := os.Getenv("HOME")
	if home == "" {
		return nil, fmt.Errorf("$HOME environment variable not set")
	}
	return &Commands{path: env, home: home}, nil
}

func (c Commands) Eval(command model.Command) model.Output {
	switch command.Name() {
	case exitCommand:
		return c.exitCMD(command)
	case echoCommand:
		return c.echoCMD(command)
	case typeCommand:
		return c.typeCMD(command)
	case pwdCommand:
		return c.pwdCMD(command)
	case cdCommand:
		return c.cdCMD(command)
	default:
		out := c.execCMD(command)
		return out
	}
}

func commandNotFound(command string) string {
	return std(fmt.Sprintf("%v: command not found", command))
}

func noSuchFileOrDirectory(command model.Command) string {
	return std(fmt.Sprintf("%v: %v: No such file or directory", command.Name(), command.Args().String()))
}

func std(std string) string {
	return std + "\n"
}
