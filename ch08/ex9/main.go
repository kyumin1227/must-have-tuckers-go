package main

import "fmt"

func main() {
	a := 3

	// go에서는 기본적으로 case 하나를 실행 후 switch를 빠져나가 break가 필요 없음
	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
		break
	case 3:
		fmt.Println("a == 3")
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")
	}
}