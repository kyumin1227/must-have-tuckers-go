package main

import "fmt"

func main() {
	i := 0
	
	f := func() {
		// 함수 리터럴에서 외부 변수를 내부 상태로 가져올 때 값 복사가 아닌 인스턴스 참조로 가져옴
		i += 10
	}

	i++

	f()

	fmt.Println(i)
}