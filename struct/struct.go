package main

import "fmt"

// Vertex X Cord, Y Cord
type Vertex struct {
	X, Y    int
	Comment string
}

func main() {
	vec := Vertex{1, 2, "South"}
	fmt.Println(vec.Comment)
}
