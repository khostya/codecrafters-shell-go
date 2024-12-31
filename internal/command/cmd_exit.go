package command

import (
	"github.com/codecrafters-io/shell-starter-go/internal/model"
	"os"
	"strconv"
)

func (c Commands) exitCMD(command model.Command) model.Output {
	if command.Args().Len() < 1 {
		return model.NewOutput(model.Stderr(std(commandNotFound(command.String()))), "")
	}
	i, err := strconv.Atoi(command.Args()[0])
	if err != nil {
		return model.NewOutput(model.Stderr(std(err.Error())), "")
	}
	os.Exit(i)
	return model.Output{}
}
