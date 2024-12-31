package command

import (
	"github.com/codecrafters-io/shell-starter-go/internal/model"
	"strings"
)

func (c Commands) echoCMD(command model.Command) model.Output {
	if command.Args().Len() < 1 {
		return model.NewOutput(model.Stderr(commandNotFound(command.String())), "")
	}
	return model.NewOutput("", model.Stdout(std(strings.Join(command.Args(), " "))))
}
