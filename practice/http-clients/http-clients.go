package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Print("err: ", err)
		}
	}()

	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
