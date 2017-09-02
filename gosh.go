package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func printPrompt(wd string) {
	fmt.Printf("%s %s ", wd, ">")
}

func readInput(inputReader *bufio.Reader, inputBuffer *string) {
	*inputBuffer, _ = inputReader.ReadString('\n')
	*inputBuffer = strings.TrimSpace(*inputBuffer)
}

func tokenizeInput(input string) (cmd string, args []string) {
	tokens := strings.Split(input, " ")
	if len(tokens) > 0 {
		cmd = tokens[0]
	}
	if len(tokens) > 1 {
		args = tokens[1:]
	}
	return cmd, args
}

func cd(wd *string, path string) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Printf("%s%s", err, "\n")
	} else {
		*wd = path
	}
}

func exc(wd *string, cmd string, args ...string) {
	if len(cmd) > 0 {
		if cmd == "cd" && len(args) > 0 {
			cd(wd, args[0])
		} else {

			command := exec.Command(cmd, args...)
			var out bytes.Buffer
			command.Stdout = &out
			err := command.Run()
			if err != nil {
				fmt.Printf("%s%s", err, "\n")
			}
			fmt.Print(out.String())
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	exit := false
	var input string
	wd, _ := os.Getwd()

	for !exit {
		printPrompt(wd)
		readInput(reader, &input)
		cmd, args := tokenizeInput(input)

		exit = input == "exit"
		if !exit {
			exc(&wd, cmd, args...)
		}
	}
}
