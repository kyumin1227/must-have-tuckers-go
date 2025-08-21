package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 10)	// len: 3, cap: 10
	slice3 := make([]int, 10)		// len: 10, cap: 10

	// 복사 시 cap은 영향을 주지 않고 둘 중 작은 len에 맞춤
	cnt1 := copy(slice2, slice1)	// slice1을 slice2에 복사
	cnt2 := copy(slice3, slice1)	// slice1을 slice3에 복사

	fmt.Println(cnt1, slice2)		// 3 [1 2 3]
	fmt.Println(cnt2, slice3)		// 5 [1 2 3 4 5 0 0 0 0 0]

	// append로 코드 개선
	// slice2 := append([]int{}, slice1...)
}