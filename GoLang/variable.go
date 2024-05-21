package main

import "fmt"

func add(x, y int) int {
	return x + y
}
func main() {
	fmt.Printf("the sum of two number is %d", add(4, 5))
}
