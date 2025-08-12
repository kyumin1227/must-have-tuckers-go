package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b)	// Scan 함수의 입력으로 사용할 때는 변수 앞에 &를 붙여 메모리 주소를 입력으로 넘김
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}