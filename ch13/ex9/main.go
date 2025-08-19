package main

import "fmt"

func main() {
	str := "Hello 월드!"

	// 추가 메모리 할당 없이 한 글자씩 순회 가능
	for _, v := range str {
		fmt.Printf("타입: %T 값: %d 문자: %c\n", v, v, v)
	}
}