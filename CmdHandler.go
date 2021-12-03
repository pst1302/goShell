package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ParsingCmd(cmd string) {
	//fmt.Println("Parsing start with " + cmd)

	splitedCmd := strings.Split(cmd, " ")

	if len(splitedCmd) == 0 {
		return
	}

	if splitedCmd[0] == "ls\n" {
		Ls_cmd_handler()
	}
}

func Ls_cmd_handler() {
	ex, err := os.Executable()
	if err != nil {
		//
	}

	exPath := filepath.Dir(ex)
	files, err := ioutil.ReadDir(exPath)
	if err != nil {
		//
	}

	for _, file := range files {
		fmt.Print(file.Name() + " ")
	}
	fmt.Println()
}
