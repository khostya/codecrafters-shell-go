package command

import (
	"github.com/codecrafters-io/shell-starter-go/internal/model"
	"os"
	"strings"
)

func (c Commands) CdAbsolutePathCMD(command model.Command) model.Output {
	err := os.Chdir(command.Args()[0])
	if err != nil {
		return model.NewOutput(model.Stderr(noSuchFileOrDirectory(command)), "")
	}
	return model.Output{}
}

func (c Commands) cdCMD(command model.Command) model.Output {
	if command.Args()[0] == "" {
		return model.Output{}
	}
	if command.Args()[0][0] == '/' {
		return c.CdAbsolutePathCMD(command)
	}

	out := c.pwdCMD(command)
	if out.Stderr() != "" {
		return out
	}

	stdout := strings.TrimRight(out.Stdout(), "\n")

	stack := strings.Split(stdout, string(os.PathSeparator))

	cd := strings.Split(command.Args()[0], string(os.PathSeparator))
	for _, s := range cd {
		if s == "" {
			continue
		}
		if s == "." {
			continue
		}
		if s == ".." {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s)
		}
	}

	path := strings.Join(stack, string(os.PathSeparator))
	return c.CdAbsolutePathCMD(model.NewCommand(cdCommand, path))
}
