package command

import (
	"github.com/codecrafters-io/shell-starter-go/internal/model"
	"os"
)

func (c Commands) pwdCMD(command model.Command) model.Output {
	dir, err := os.Getwd()
	if err != nil {
		return model.NewOutput(model.Stderr(commandNotFound(command.Name())), "")
	}
	return model.NewOutput("", model.Stdout(std(dir)))
}
