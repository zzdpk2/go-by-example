package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp: ", s, "len: ", len(s), "cap: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[2:]
	fmt.Println("sl3: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	a := []int{1, 2, 3, 4}
	fmt.Println("len of a1: ", len(a))
	fmt.Println("cap of a1: ", cap(a))
	fmt.Println("a1: ", a)
	b := a[2:]
	fmt.Println("b1: ", b)
	fmt.Println("cap of b1: ", cap(b))
	fmt.Println("len of a1: ", len(a))
	fmt.Println("cap of a1: ", cap(a))
	b = append(b, 100)
	fmt.Println("a2: ", a)
	fmt.Println("b2: ", b)
	fmt.Println("cap of b2: ", cap(b))
}
