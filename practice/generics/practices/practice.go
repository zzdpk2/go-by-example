package main

import "fmt"
import "golang.org/x/exp/constraints"

func MapKeys[K comparable, V any](m map[K]V) []K {
	res := make([]K, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(MapKeys(m))

	fmt.Println(Max(3, 5))         // 返回 5
	fmt.Println(Max("cat", "dog")) // 返回 "dog"
}
