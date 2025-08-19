package main

import (
	"fmt"
	"strings"
)

func ToUpper1(str string) string {
	// 합 연산 마다 매번 메모리 공간이 버려짐
	var rst string
	for _, c := range str {
		if c >= 'a' && 'c' <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

func ToUpper2(str string) string {
	// 기존 메모리 공간의 빈자리에 더함
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	var str string = "Hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}