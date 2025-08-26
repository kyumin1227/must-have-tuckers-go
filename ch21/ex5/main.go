package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다")	// 프로그램 종료
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(9, 0)
}