package main

import "fmt"

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},	// 닫는 중괄호와 다른 줄이기 때문에 쉼표
	}

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}