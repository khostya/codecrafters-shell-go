package command

import (
	"bytes"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/path"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	typeCommand = "type"
	exitCommand = "exit"
	echoCommand = "echo"
	pwdCommand  = "pwd"
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

func (c Commands) typeCMD(command []string) Output {
	if len(command) <= 1 {
		return Output{stderr: std(commandNotFound(strings.Join(command, " ")))}
	}
	switch command[1] {
	case exitCommand, typeCommand, echoCommand, pwdCommand:
		return Output{stdout: std(fmt.Sprintf("%s is a shell builtin", command[1]))}
	default:
		path, err := c.path.FindPath(command[1])
		if err != nil {
			return Output{stderr: std(strings.Join(command[1:], " ") + ": not found")}
		}
		return Output{stdout: std(fmt.Sprintf("%s is %s", command[1], path))}
	}
}

func (c Commands) exitCMD(command []string) Output {
	if len(command) <= 1 {
		return Output{stderr: std(commandNotFound(strings.Join(command, " ")))}
	}
	i, err := strconv.Atoi(command[1])
	if err != nil {
		return Output{stderr: std(err.Error())}
	}
	os.Exit(i)
	return Output{}
}

func (c Commands) echoCMD(command []string) Output {
	if len(command) < 1 {
		return Output{stderr: commandNotFound(strings.Join(command, " "))}
	}
	return Output{stdout: std(strings.Join(command[1:], " "))}
}

func (c Commands) execCMD(command []string) Output {
	cmd := exec.Command(command[0], command[1:]...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return Output{stderr: commandNotFound(command[0])}
	}
	if cmd.ProcessState.Success() {
		return Output{stdout: stdout.String()}
	}
	return Output{stderr: commandNotFound(command[0])}
}

func (c Commands) Eval(command []string) Output {
	switch command[0] {
	case exitCommand:
		return c.exitCMD(command)
	case echoCommand:
		return c.echoCMD(command)
	case typeCommand:
		return c.typeCMD(command)
	case pwdCommand:
		return c.pwdCMD(command)
	default:
		out := c.execCMD(command)
		return out
	}
}

func (c Commands) pwdCMD(command []string) Output {
	dir, err := os.Getwd()
	if err != nil {
		return Output{stderr: commandNotFound(command[0])}
	}
	return Output{stdout: std(dir)}
}

func commandNotFound(command string) string {
	return std(fmt.Sprintf("%v: command not found", command))
}

func std(std string) string {
	return std + "\n"
}
