package main

import "fmt"

// defer for multi time, the evaluation order will be like items pushed into a stack
// they're reversed
func main() {
	fmt.Println("counting")

	defer fmt.Println("done")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

}
