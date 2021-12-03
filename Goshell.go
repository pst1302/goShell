package main

import "fmt"

func main() {

	// Hello msg
	PrintHello()

	for {
		cmd := ""
		fmt.Print("goShell > ")
		fmt.Scanln(&cmd)

		if cmd == "exit" {
			break
		}
		if cmd == "" {
			continue
		}

		ParsingCmd(cmd)
	}

}
