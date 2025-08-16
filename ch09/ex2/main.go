package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	// 초기문, 조건문, 후처리 모두 생략 가능
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}
}