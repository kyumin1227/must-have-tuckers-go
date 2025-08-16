package main

import "fmt"

func main() {
	for i := range 5 {
		for j := 5 - i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}