package main

import "fmt"

func getPt(i, j int) (*int, *int) {

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	q := &j        // 指向 j
	*q = *q / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值

	return p, q
}

func main() {
	p, q := getPt(42, 2701)
	fmt.Println(*p, *q)
}
