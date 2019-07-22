package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "golang社区"
	fmt.Println(s)

	for i, ch := range []byte(s) {
		fmt.Println(i, ch)
	}
	fmt.Println()

	// %X 十六进制表示，字母形式为大写 A-F
	//结果: 67, 6F, 6C, 61, 6E, 67, E7 A4 BE, E5 8C BA
	//utf-8编码,占3个字节
	for _, ch := range []byte(s) {
		fmt.Printf("%X ", ch)
	}
	fmt.Println()

	//结果：(0,67)(1,6F)(2,6C)(3,61)(4,6E)(5,67)(6,793E)(9,533A)
	//unicode 编码，占2个字节
	for i, ch := range s {
		fmt.Printf("(%d,%X)", i, ch)
	}
	fmt.Println()

	bytes := []byte(s)
	//所有字节，结果：[103 111 108 97 110 103 231 164 190 229 140 186]
	fmt.Println(bytes)
	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", r)
	}
	fmt.Println()

	//rune占4个字节，rune相等于golang成char
	for _, ch := range []rune(s) {
		fmt.Printf("%c", ch)
	}
	fmt.Println()

	//字符数 结果：8
	fmt.Println(utf8.RuneCountInString(s))
	//字节长度，结果：12
	fmt.Println(len(s))
}
