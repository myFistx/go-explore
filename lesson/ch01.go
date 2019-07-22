package main

import "fmt"

//寻找最长不包含重复字符串的子串
//用一个map记录该字符串中每个字符最后出现的位置，每出现一个不重复的字符，最大长度就加1

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int) //记录字符最后出现的位置
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		fmt.Println(i, ch)
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func deal(s string) int {
	item := make(map[rune]int)
	start := 0
	maxL := 0
	for i, ch := range []rune(s) {
		//fmt.Println(i, ch)
		if v, ok := item[ch]; ok && v >= start {
			start = v + 1
		}
		if i+1-start > maxL {
			maxL = i + 1 - start
		}
		item[ch] = i
	}
	return maxL
}

func main() {

	fmt.Println(deal("abcabcbb"))
	fmt.Println(deal("golang社区"))
	fmt.Println(deal("一二三二一"))

}
