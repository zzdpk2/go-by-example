package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	_, err = exec.Command("date", "-x").Output()
	// if err != nil {
	// 	switch e := err.(type) {
	// 	case *exec.Error:
	// 		fmt.Println("failed executing:", err)
	//
	// 	case *exec.ExitError:
	// 		fmt.Println("command exit rc =", e.ExitCode())
	// 	default:
	// 		panic(err)
	// 	}
	// }
	if err != nil {
		var exitErr *exec.ExitError
		var execErr *exec.Error

		switch {
		case errors.As(err, &execErr):
			fmt.Println("failed executing:", execErr)

		case errors.As(err, &exitErr):
			fmt.Println("command exit rc =", exitErr.ExitCode())

		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	// grepCmd.Start()

	if err := grepCmd.Start(); err != nil {
		log.Fatalf("failed to start grep: %v", err)
	}

	if _, err := grepIn.Write([]byte("hello grep\ngoodbye grep")); err != nil {
		log.Fatalf("failed to write into grep: %v", err)
	}
	if err := grepIn.Close(); err != nil {
		log.Fatalf("failed to close grep: %v", err)
	}
	grepBytes, _ := io.ReadAll(grepOut)
	if err := grepCmd.Wait(); err != nil {
		log.Fatalf("failed to wait for grep: %v", err)
	}

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Since it is pipe ops, bash -c should be added
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
