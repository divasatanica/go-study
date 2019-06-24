package main

import "fmt"

func invoke() string {
	fmt.Println("OJBK")
	return "coma"
}

// the function invoke will be immediately invoked,
// but defer statement will be delayed
func main() {
	defer fmt.Println("world, ", invoke())

	fmt.Println("hello")
}
