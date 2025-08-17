package main

import "fmt"

func main() {
	nums := [...]int{10, 20, 30, 40, 50}	// ...일 경우 중괄호 안의 값의 개수가 배열의 길이

	nums[2] = 300

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}