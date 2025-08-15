package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	// 초기문으로 switch 내부에서 age를 사용
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is", age)
	}

	fmt.Println("age is", age)	// age 변수는 소멸
}