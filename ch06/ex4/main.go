package main

import "fmt"

// 상수는 컴파일 과정에서 리터럴로 변환

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	var a int = PI * 100
	var b int = FloatPI * 100	// 타입 있는 상수의 경우 타입 오류 발생

	fmt.Println(a)
	fmt.Println(b)
}