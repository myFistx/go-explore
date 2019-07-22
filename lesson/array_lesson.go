package main

import "fmt"

func main() {
	//数组定义
	var arr1 [3]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(arr1, arr2, arr3)

	printArr(&arr2)

	for i, v := range arr2 {
		fmt.Println(i, v)
	}
}

func printArr(arr *[5]int) {
	arr[0] = 12
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
