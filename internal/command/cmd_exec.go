package command

import (
	"bytes"
	"github.com/codecrafters-io/shell-starter-go/internal/model"
	"os/exec"
)

func (c Commands) execCMD(command model.Command) model.Output {
	cmd := exec.Command(command.Name(), command.Args()...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return model.NewOutput(model.Stderr(commandNotFound(command.Name())), "")
	}
	if cmd.ProcessState.Success() {
		return model.NewOutput("", model.Stdout(stdout.String()))
	}

	return model.NewOutput(model.Stderr(commandNotFound(command.Name())), "")
}
