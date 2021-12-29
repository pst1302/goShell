package main

import (
	"fmt"
	"os"
)

func Init() {

	curDir, err := os.UserHomeDir()
	if err != nil {

	}

	curLocal = curDir

	fmt.Printf("Cur dir is %s\n", curDir)
}
