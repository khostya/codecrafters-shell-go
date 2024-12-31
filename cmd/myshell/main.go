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

		s, err := cmd.Eval(strings.Split(command[:len(command)-1], " "))
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(s)
		}
	}
}
