package command

import (
	"github.com/codecrafters-io/shell-starter-go/internal/model"
	"os"
	"strings"
)

func (c Commands) cdAbsolutePathCMD(command model.Command) model.Output {
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

	arg := command.Args()[0]
	if arg[0] == '~' {
		arg = c.home + arg[1:]
	}

	if arg[0] == '/' {
		return c.cdAbsolutePathCMD(model.NewCommand(command.Name(), arg))
	}

	out := c.pwdCMD(command)
	if out.Stderr() != "" {
		return out
	}

	stdout := strings.TrimRight(out.Stdout(), "\n")

	stack := strings.Split(stdout, string(os.PathSeparator))

	cd := strings.Split(arg, string(os.PathSeparator))
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
	return c.cdAbsolutePathCMD(model.NewCommand(cdCommand, path))
}
