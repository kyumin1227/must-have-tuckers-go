package main

import "fmt"

func ChangeArray(arr [5]int) {
	arr[3] = 3000
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	ChangeArray(a)	// 값이 복사되기 때문에 원본 배열은 값이 변하지 않음

	fmt.Println(a[3])	// 4
}