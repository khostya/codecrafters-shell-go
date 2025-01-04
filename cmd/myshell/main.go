package main

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/command"
	"github.com/codecrafters-io/shell-starter-go/internal/model"
	"github.com/codecrafters-io/shell-starter-go/internal/split"
	"log"
	"os"
)

func main() {
	cmd, err := command.NewCommands()
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		args := split.Split(command[:len(command)-1])
		output := cmd.Eval(model.NewCommand(args[0], args[1:]...))
		if output.Stdout() != "" {
			fmt.Fprint(os.Stdout, output.Stdout())
		}
		if output.Stderr() != "" {
			fmt.Fprint(os.Stdout, output.Stderr())
		}
	}
}
