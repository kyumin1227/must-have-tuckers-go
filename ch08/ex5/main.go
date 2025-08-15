package main

import "fmt"

func main() {
	temp := 18

	// switch 다음에 비교값을 적지 않는 경우 default 값으로 true를 사용
	switch {
	case temp < 10, temp > 30:
		fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다.")
	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 수 있으니 가벼운 겉옷을 준비하세요.")
	case temp >= 15 && temp < 25:
		fmt.Println("야외 활동하기 좋은 날씨입니다.")
	default:
		fmt.Println("따뜻합니다.")
	} 
}