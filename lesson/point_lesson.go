package main

import "fmt"

//参数值传递
func P1(a int) {
	a++
	fmt.Println(a)
}

//参数应用传递，传的是一个地址，该地址所指向的值用*表示
func P2(a *int) {
	*a++
	fmt.Println(*a)
}

func main() {

	a := 3
	fmt.Println(a)
	P1(a)
	fmt.Println(a)
	P2(&a)
	fmt.Println(a)
}
