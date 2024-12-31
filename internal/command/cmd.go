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
		return Output{stdout: "", stderr: strings.Join(command, " ") + ": command not found\n"}
	}
	switch command[1] {
	case exitCommand, typeCommand, echoCommand:
		return Output{stdout: fmt.Sprintf("%s is a shell builtin\n", command[1]), stderr: ""}
	default:
		path, err := c.path.FindPath(command[1])
		if err != nil {
			return Output{stdout: "", stderr: strings.Join(command[1:], " ") + ": not found\n"}
		}
		return Output{stdout: fmt.Sprintf("%s is %s\n", command[1], path), stderr: ""}
	}
}

func (c Commands) exitCMD(command []string) Output {
	if len(command) <= 1 {
		return Output{stdout: "", stderr: strings.Join(command, " ") + ": command not found\n"}
	}
	i, err := strconv.Atoi(command[1])
	if err != nil {
		return Output{stdout: "", stderr: err.Error() + "\n"}
	}
	os.Exit(i)
	return Output{}
}

func (c Commands) echoCMD(command []string) Output {
	if len(command) < 1 {
		return Output{stdout: "", stderr: strings.Join(command, " ") + ": command not found\n"}
	}
	return Output{stdout: strings.Join(command[1:], " ") + "\n", stderr: ""}
}

func (c Commands) execCMD(command []string) Output {
	cmd := exec.Command(command[0], command[1:]...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return Output{stdout: "", stderr: fmt.Sprintf("%v: command not found\n", command[0])}
	}
	if cmd.ProcessState.Success() {
		return Output{stdout: stdout.String()}
	}
	return Output{stderr: fmt.Sprintf("%v: command not found\n", command[0])}
}

func (c Commands) Eval(command []string) Output {
	switch command[0] {
	case exitCommand:
		return c.exitCMD(command)
	case echoCommand:
		return c.echoCMD(command)
	case typeCommand:
		return c.typeCMD(command)
	default:
		out := c.execCMD(command)
		return out
	}
}
