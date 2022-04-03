package main

import "fmt"

func test() {
	x := 1
	defer func() {
		fmt.Printf("first %d\n", x)
	}()
	x = 2
	defer func() {
		fmt.Printf("second %d\n", x)
	}()
	defer func() {
		fmt.Printf("third %d\n", x)
	}()
}

func main() {
	test()
}
