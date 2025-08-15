package main

import "fmt"

func main() {
	const C int = 10

	var b int = C * 10
	C = 20	// 상수는 대입문 좌변에 올수 없음
	fmt.Println(&C)	// 상수의 주소는 출력할 수 없음
}