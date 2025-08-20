package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	randomNum := rand.Intn(100)
	var count int
	
	for {
		count++
		fmt.Print("숫자값을 입력하세요: ")
		n, err := InputIntValue()

		if err != nil {
			fmt.Println("숫자만 입력하세요.")
			continue
		}
		
		if n > randomNum {
			fmt.Println("입력하신 숫자가 더 큽니다.")
		} else if n < randomNum {
			fmt.Println("입력하신 숫자가 더 작습니다.")
		} else {
			fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수:", count)
			break
		}
	}
}