package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Hello msg
	PrintHello()

	Init()

	for {
		cmd := ""
		fmt.Printf("goShell %s > ", curLocal)

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		//cmd = string(input)
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
