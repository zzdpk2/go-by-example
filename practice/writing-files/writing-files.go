package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}(f)

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Write into disk in hard way
	// defer func() {
	// 	if r := f.Sync(); r != nil {
	// 		panic(r)
	// 	}
	// }()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	defer func(w *bufio.Writer) {
		if r := w.Flush(); r != nil {
			panic(r)
		}
	}(w)
}
