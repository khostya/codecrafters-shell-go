package command

type Command struct {
	name string
	args []string
}

func (cmd *Command) Name() string {
	return cmd.name
}

func (cmd *Command) Args() []string {
	return cmd.args
}

type Output struct {
	stderr, stdout string
}

func (o *Output) Stdout() string {
	return o.stdout
}

func (o *Output) Stderr() string {
	return o.stderr
}
