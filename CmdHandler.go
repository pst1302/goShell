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
		Ls_cmd_handler("")
	} else if splitedCmd[0] == "ls" {
		if len(splitedCmd) == 2 {
			Ls_cmd_handler(splitedCmd[1])
		}
	}
}

func Ls_cmd_handler(addedString string) {

	if addedString != "" {
		if strings.HasPrefix(addedString, "-") {
			// arg is option
			fmt.Println("arg is option")
		} else {
			// arg is path
			fmt.Println("arg is path : " + addedString)

			fmt.Println(GetFilesFromDir(addedString))
		}
	} else {
		ex, err := os.Executable()
		if err != nil {
			//
		}

		fmt.Println(GetFilesFromDir(ex))
	}
}

func GetFilesFromDir(dir string) string {

	dir = strings.ReplaceAll(dir, "\n", "")

	fileInfo, err := os.Stat(dir)
	if err != nil {
		// error handling
		fmt.Println("There is something problem")
		if os.IsNotExist(err) {
			dir += "/"
			fileInfo, err := os.Stat(dir)

			fmt.Println("dir ::: " + dir)
			if err != nil {
				return "There is some problem while reading directory"
			}

			if !fileInfo.IsDir() {
				return "There is some problem while reading directory"
			}
		}
	}

	if fileInfo != nil {
		fmt.Println("file Info isn't nil")
		dir += "/"
	}

	exPath := filepath.Dir(dir)

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
