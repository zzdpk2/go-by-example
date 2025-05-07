package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	p := filepath.Join("dir1", "dir2", "file1.txt")
	fmt.Println("p:", p)

}
