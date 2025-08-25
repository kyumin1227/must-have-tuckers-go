package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 0
	m[2] = 2
	m[3] = 3

	delete(m, 3)
	delete(m, 4)
	fmt.Println(m[3])	// 삭제된 요소를 출력할 경우 기본값 반환
	fmt.Println(m[1])

	// 값이 기본 값이 경우랑 없는 경우를 구분하기 위해서는 반환을 둘로 받아 존재 여부를 확인
	// v, ok := m[3]
}