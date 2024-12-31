package command

import (
	"github.com/codecrafters-io/shell-starter-go/internal/model"
	"os"
)

func (c Commands) CdAbsolutePathCMD(command model.Command) model.Output {
	err := os.Chdir(command.Args()[0])
	if err != nil {
		return model.NewOutput(model.Stderr(noSuchFileOrDirectory(command)), "")
	}
	return model.Output{}
}

func (c Commands) cdCMD(command model.Command) model.Output {
	return c.CdAbsolutePathCMD(command)
}
