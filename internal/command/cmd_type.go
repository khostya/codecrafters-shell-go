package command

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/model"
)

func (c Commands) typeCMD(command model.Command) model.Output {
	if command.Args().Len() < 1 {
		return model.NewOutput(model.Stderr(std(commandNotFound(command.String()))), "")
	}
	switch command.Args()[0] {
	case exitCommand, typeCommand, echoCommand, pwdCommand:
		return model.NewOutput("", model.Stdout(std(fmt.Sprintf("%s is a shell builtin", command.Args()[0]))))
	default:
		path, err := c.path.FindPath(command.Args()[0])
		if err != nil {
			return model.NewOutput(model.Stderr(std(command.Args().String()+": not found")), "")
		}
		return model.NewOutput("", model.Stdout(std(fmt.Sprintf("%s is %s", command.Args()[0], path))))
	}
}
