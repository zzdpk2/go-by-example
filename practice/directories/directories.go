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
	if err := os.RemoveAll(basedir); err != nil {
		check(err)
	}
	err := os.Mkdir(basedir, 0755)
	check(err)
	defer func() {
		if err := os.RemoveAll(basedir); err != nil {
			panic(err)
		}
	}()

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile(filepath.Join(basedir, "/file1"))
	err = os.MkdirAll(filepath.Join(basedir, "/parent/child"), 0755)
	check(err)

	createEmptyFile(filepath.Join(basedir, "/parent/file2"))
	createEmptyFile(filepath.Join(basedir, "/parent/file3"))
	createEmptyFile(filepath.Join(basedir, "/parent/child/file4"))

	c, err := os.ReadDir(filepath.Join(basedir, "/parent"))
	check(err)

	fmt.Println("Listing " + filepath.Join(basedir, "/parent"))
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir(
		filepath.Join(basedir,
			filepath.Join("parent",
				filepath.Join("child"))))
	// err = os.Chdir(filepath.Dir(filepath.Dir(
	// 	filepath.Dir(
	// 		filepath.Join(basedir, "parent", "child")))))
	check(err)

	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing " + filepath.Join(basedir, "/parent/child"))
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// err = os.Chdir("../../..")
	err = os.Chdir(filepath.Join("..", filepath.Join("..", filepath.Join(".."))))
	check(err)

	fmt.Println("Visiting " + basedir)
	err = filepath.Walk(basedir, visit)
	check(err)
}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, info.IsDir())
	return nil
}
