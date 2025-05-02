package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("Creating a file")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	_, err := fmt.Fprintln(f, "data")
	if err != nil {
		panic("write error!")
	}
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		panic(err)
	}
}
