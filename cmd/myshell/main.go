package main

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/command"
	"log"
	"os"
	"strings"
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

		sp := strings.Split(command[:len(command)-1], " ")
		output := cmd.Eval(sp)
		if output.Stdout() != "" {
			fmt.Fprint(os.Stdout, output.Stdout())
		}
		if output.Stderr() != "" {
			fmt.Fprint(os.Stdout, output.Stderr())
		}
	}
}
