package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	basedir := "subdir"
	err := os.Mkdir(basedir, 0755)
	check(err)
	defer func(path string) {
		if err := os.RemoveAll(path); err != nil {
			panic(err)
		}
	}(basedir)

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile(basedir + "/file1")
	err = os.MkdirAll(basedir+"/parent/child", 0755)
	check(err)

	createEmptyFile(basedir + "/parent/file2")
	createEmptyFile(basedir + "/parent/file3")
	createEmptyFile(basedir + "/parent/child/file4")

	c, err := os.ReadDir(basedir + "/parent")
	check(err)

	fmt.Println("Listing " + basedir + "/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir(basedir + "/parent/child")
	check(err)

	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing " + basedir + "/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("Visiting " + basedir)
	err = filepath.Walk(basedir, visit)
}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, info.IsDir())
	return nil
}
