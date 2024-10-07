package main

import "fmt"

func main() {

	a := 10
	b := 5

	result := if_func(a, b)

	fmt.Println(result)

	result = for_func(10)
	fmt.Println(result)
}

func for_func(a int) int {
	for i := 0; i < a; i++ {
		fmt.Println(i)
	}
	return a
}

func if_func(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
