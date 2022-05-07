package main

import "fmt"

func add2(n int) {
	n += 2
}

func add2ptr(n *int) {
	*n += 2
}

func main() {
	n := 5
	add2(n)
	fmt.Println(n) // 5
	add2ptr(&n)
	fmt.Println(n) // 7
}

/*
	指针带来了
		1. 可以修改原指向的内容
		2. 减少拷贝开销（大对象的时候）
*/
