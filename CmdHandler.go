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
		Ls_cmd_handler(curLocal)
	} else if splitedCmd[0] == "ls" {
		if len(splitedCmd) == 2 {
			Ls_cmd_handler(splitedCmd[1])
		}
	}
}

func Ls_cmd_handler(arg string) {

	if arg != "" {
		if strings.HasPrefix(arg, "-") {
			// arg is option
			fmt.Println("arg is option")
		} else {
			// arg is path
			//fmt.Println("arg is path : " + arg)

			fmt.Println(GetFilesFromDir(arg))
		}
	} else {

		fmt.Println(GetFilesFromDir(curLocal))
	}
}

func GetFilesFromDir(dir string) string {

	dir = strings.ReplaceAll(dir, "\n", "")

	fileInfo, err := os.Stat(dir)
	if err != nil {
		// Error handling
	}

	exPath := ""
	if fileInfo != nil {
		fmt.Println("fileInfo is not null")
		if fileInfo.IsDir() {
			fmt.Println("It is dir")
			exPath = dir + "/"
		} else {
			fmt.Println("It is file")
			exPath = filepath.Dir(dir)
		}
	} else {
		fmt.Print(dir + " is not file or directory.")
		//exPath = filepath.Dir(dir)
	}

	files, err := ioutil.ReadDir(exPath)
	if err != nil {
		//
	}

	filesList := ""

	for _, file := range files {
		filesList += fmt.Sprint(file.Name() + " ")
	}

	return filesList
}
