package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v\n", b.num)
}

type container struct {
	base
	str string
}

type docker struct {
	container
	namespace string
}

func (d docker) describe() string {
	return fmt.Sprintf("docker with namespace = %v\n", d.namespace)
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num: ", co.base.num)
	fmt.Println("describe: ", co.describe())

	type describer interface {
		describe() string
	}

	d1 := docker{namespace: "nginx"}
	var d11 describer = d1
	fmt.Println("also describe: ", d11.describe())
}
