package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10	// 패닉 발생
	fmt.Println(slice)
}

var array = [...]int{1, 2, 3}	// 배열 선언
var slice = []int{1, 2, 3}		// 슬라이스 선언