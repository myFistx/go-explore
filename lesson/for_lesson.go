package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//省略初始条件
//数字转成0和1，13/2=6余1，6/2=3余0，3/2=1余1，1/2=0余1，从后往前排1101
func convertToBin(n int) string {
	res := ""
	for ; n > 0; n /= 2 {
		lab := n % 2
		res = strconv.Itoa(lab) + res
	}
	return res
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func p1(a int) {
	a++
	fmt.Println(a)
}

func p2(a *int) {
	*a++
	fmt.Println(*a)
}

func main() {
	fmt.Println(convertToBin(13))
	fmt.Println(convertToBin(344222333333))

	a := 4
	p1(a)
	p2(&a)
	fmt.Println(a)

	printFile("abc.txt")
}
