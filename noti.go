package main

import (
	"fmt"
)

// Program information about version, build date,
const version string = "0.0.0.1"
const last_release_date string = "2021-12-02"

// Print instruction message of program. this fuction is printed when user program start
// If you add notification of program. you can add at this fuction
// mady by Gilbok (21/12/01)
func PrintHello() {

	fmt.Println("Hello go shell")
	fmt.Println("This is made by Gilbok.")
	fmt.Printf("Latest version is %s. It is updated at %s\n", version, last_release_date)
}
