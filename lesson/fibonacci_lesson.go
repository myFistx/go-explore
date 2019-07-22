package main

import "fmt"

//斐波那契数列，即每个数均为前两个数之和
//递归调用，当基数过大，递归调用次数会过多，影响性能

func main() {

	fmt.Println("递归实现斐波那契数列:")
	result := 0
	for i := 0; i < 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}

	fmt.Println("闭包实现斐波那契数列:")
	f := fibonacci2()
	for i := 0; i < 10; i++ {
		res := f()
		fmt.Printf("fibonacci(%d) is: %d\n", i, res)
	}

	r := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(r(i))
	}

}

//闭包实现斐波那契数列
func fibonacci2() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//闭包
func adder() func(v int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

//递归实现斐波那契数列
func fibonacci(n int) (ret int) {
	if n <= 1 {
		ret = 1
	} else {
		ret = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
