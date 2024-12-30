package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func exitCMD(command []string) error {
	if len(command) < 1 {
		return errors.New("no command given")
	}
	i, err := strconv.Atoi(command[1])
	if err != nil {
		return err
	}
	os.Exit(i)
	return nil
}

func eval(command []string) error {
	switch command[0] {
	case "exit":
		return exitCMD(command)
	default:
		return errors.New("unknown command: " + command[0])
	}
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		err = eval(strings.Split(command[:len(command)-1], " "))
		if err != nil {
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}
