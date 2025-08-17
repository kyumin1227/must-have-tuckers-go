package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a	// 배열의 크기와 타입이 같기 때문에 복사 가능

	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}