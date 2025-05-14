package main

import (
	"fmt"
	"os"
)

func main() {
	// Defer will not be called cuz the OS is exited
	defer fmt.Println("!")
	os.Exit(3)
}
