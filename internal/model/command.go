package model

import (
	"fmt"
	"strings"
)

type (
	Args []string

	Stdout string
	Stderr string

	Command struct {
		name string
		args Args
	}

	Output struct {
		stderr, stdout string
	}
)

func NewOutput(stderr Stderr, stdout Stdout) Output {
	return Output{string(stderr), string(stdout)}
}

func NewCommand(name string, args ...string) Command {
	return Command{
		name: name,
		args: args,
	}
}

func (args Args) String() string {
	return strings.Join(args, " ")
}

func (cmd *Command) String() string {
	return fmt.Sprintf("%v %v", cmd.name, cmd.args.String())
}

func (cmd *Command) Name() string {
	return cmd.name
}

func (cmd *Command) Args() Args {
	return cmd.args
}

func (args Args) Len() int {
	return len(args)
}

func (o *Output) Stdout() string {
	return o.stdout
}

func (o *Output) Stderr() string {
	return o.stderr
}
