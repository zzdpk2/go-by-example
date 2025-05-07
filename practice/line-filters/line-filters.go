package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, "error:", err)
		if err != nil {
			log.SetPrefix("[ERROR] ")
			log.Println("This is an error")
		}
		os.Exit(1)
	}
}
