package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	slice = append(slice, 0)

	idx := 2

	for i := len(slice) - 2; i >= idx; i-- {
		slice[i + 1] = slice[i]
	}

	slice[idx] = 100

	fmt.Println(slice)

	// append로 코드 개선
	// slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)
}
