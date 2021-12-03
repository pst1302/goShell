package main

import (
	"fmt"
	"strings"
)

func ParsingCmd(cmd string) {
	fmt.Println("Parsing start with " + cmd)

	splitedCmd := strings.Split(cmd, " ")

	fmt.Printf("length : %d\n", len(splitedCmd))

	for i := 0; i < len(splitedCmd); i++ {

		fmt.Println(splitedCmd[i])
	}
}
