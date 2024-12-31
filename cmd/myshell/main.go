package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	typeCommand = "type"
	exitCommand = "exit"
	echoCommand = "echo"
)

func typeCMD(command []string) (string, error) {
	if len(command) <= 1 {
		return "", errors.New(strings.Join(command, " ") + ": command not found")
	}
	switch command[1] {
	case exitCommand, typeCommand, echoCommand:
		return fmt.Sprintf("%s is a shell builtin", command[1]), nil
	default:
		return "", errors.New(strings.Join(command[1:], " ") + ": not found")
	}
}

func exitCMD(command []string) error {
	if len(command) <= 1 {
		return errors.New(strings.Join(command, " ") + ": command not found")
	}
	i, err := strconv.Atoi(command[1])
	if err != nil {
		return err
	}
	os.Exit(i)
	return nil
}

func echoCMD(command []string) (string, error) {
	if len(command) < 1 {
		return "", errors.New(strings.Join(command, " ") + ": command not found")
	}
	return strings.Join(command[1:], " "), nil
}

func eval(command []string) (string, error) {
	switch command[0] {
	case exitCommand:
		return "", exitCMD(command)
	case echoCommand:
		return echoCMD(command)
	case typeCommand:
		return typeCMD(command)
	default:
		return "", errors.New(strings.Join(command, " ") + ": command not found")
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

		s, err := eval(strings.Split(command[:len(command)-1], " "))
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(s)
		}
	}
}
