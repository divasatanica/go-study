package main

import (
	"fmt"
	"math/rand"
	"time"
)

func twoSum(arr []int, target int) (int, int) {
	i := 0
	j := len(arr) - 1

	for i < j {
		if v := arr[i] + arr[j]; v < target {
			i++
		} else if v == target {
			return i, j
		} else {
			j--
		}
	}
	return -1, -1
}

func main() {
	arrToSearch := [8]int{1, 3, 5, 6, 7, 9, 10, 15}
	slice := make([]int, len(arrToSearch))
	for i, v := range arrToSearch {
		slice[i] = v
	}
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(25)
	fmt.Printf("The target number is %d\n", target)
	a, b := twoSum(slice, target)
	if resultFound := (a == -1 && b == -1); resultFound {
		fmt.Println("Target Not Found...")
	} else {
		fmt.Printf("The result is %d, %d", a, b)
	}
}
