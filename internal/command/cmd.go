package command

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/model"
	"github.com/codecrafters-io/shell-starter-go/internal/path"
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
}

func NewCommands() (*Commands, error) {
	env, err := path.NewFromDefaultEnv()
	if err != nil {
		return nil, err
	}
	return &Commands{env}, nil
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
