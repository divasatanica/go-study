package main

import "fmt"

func main() {
	while()
}

func while() {
	sum := 1
	for sum < 45 {
		sum += sum
	}
	fmt.Printf("Sum is %d", sum);
}