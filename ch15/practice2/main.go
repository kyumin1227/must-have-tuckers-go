package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const (
	initMoney = 1000
	earn = 500
	lose = 100
	victoryPoint = 5000
	gameoverPoint = 0
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

func GetRandomNum() int {
	return rand.Intn(5) + 1
}

func main() {
	money := initMoney
	
	for money > gameoverPoint && money < victoryPoint {
		num := GetRandomNum()
		fmt.Print("1에서 5사이의 값을 입력하세요: ")
		n, err := InputIntValue()

		if err != nil {
			fmt.Println("숫자만 입력하세요.")
			continue
		}

		if n < 1 || n > 5 {
			fmt.Println("1에서 5사이의 값을 입력해주세요.")
			continue
		}

		if num == n {
			money += earn
			fmt.Println("축하합니다. 맞추셨습니다. 현재 돈 :", money)
		} else {
			money -= lose
			fmt.Println("아쉽지만 다음 기회에. 현재 돈 :", money)
		}
	}

	if money >= victoryPoint {
		fmt.Println("게임 승리")
	} else {
		fmt.Println("게임 패배")
	}
}