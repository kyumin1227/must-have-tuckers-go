package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt ++
	return cnt
}

func main() {
	// 쇼트서킷에 의해 IncreaseAndReturn 함수가 호출되지 않음
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 increase")
	}
	if true || IncreaseAndReturn() < 5 {
		fmt.Println("2 increase")
	}

	fmt.Println("cnt:", cnt)
}