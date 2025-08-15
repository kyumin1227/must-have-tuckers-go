package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
		break
	case 3:
		fmt.Println("a == 3")
		fallthrough	// 다음 case도 함께 실행, 코드를 읽을 때 혼돈을 야기하므로 되도록 사용을 지양
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")
	}
}