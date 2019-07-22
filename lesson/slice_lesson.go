package main

import "fmt"

func main() {
	var s []int

	for i := 0; i < 10; i++ {
		printSlice(s)
		s = append(s, i)
	}

	fmt.Println(s)

	s1 := []int{2, 3, 4, 5, 6, 7}
	printSlice(s1)

	s2 := make([]int, 10)
	s3 := make([]int, 5, 10)
	printSlice(s2)
	printSlice(s3)

	copy(s2, s1)
	printSlice(s2)

	//删除中间某个元素
	s2 = append(s2[:5], s2[6:]...)
	printSlice(s2)

}

func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d\n", s, len(s), cap(s))
}
