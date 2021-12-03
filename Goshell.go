package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Hello msg
	PrintHello()

	for {
		cmd := ""
		fmt.Print("goShell > ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		cmd = string([]byte(input))
		if cmd == "exit\n" {
			break
		}
		if cmd == "\n" {
			continue
		}

		ParsingCmd(cmd)
	}

}
