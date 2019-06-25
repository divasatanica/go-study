package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	valueN1 := 0
	valueN2 := 0
	return func() int {
		var result int
		if value := valueN1 + valueN2; value == 0 {
			result = value
			valueN2++
		} else {
			result = value
			valueN2 = valueN1
		}
		valueN1 = result
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
