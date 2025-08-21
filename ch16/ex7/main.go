package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]	// 슬라이싱으로 배열의 1번째 부분의 주소를 가리킴

	fmt.Println("array:", array)							// array: [1 2 3 4 5]
	fmt.Println("slice:", slice, len(slice), cap(slice))	// slice: [2] 1 4

	array[1] = 100

	fmt.Println("After change second element")
	fmt.Println("array:", array)							// array: [1 100 3 4 5]
	fmt.Println("slice:", slice, len(slice), cap(slice))	// slice: [100] 1 4

	slice = append(slice, 500)

	fmt.Println("After append 500")
	fmt.Println("array:", array)							// array: [1 100 500 4 5]
	fmt.Println("slice:", slice, len(slice), cap(slice))	// slice: [100 500] 2 4
}