package main

import "fmt"

type Student struct {
	Age		int
	No		int
	Score	float64
}

func PrintStudent(s Student) {
	fmt.Printf("나이: %d 번호: %d 점수: %.2f\n", s.Age, s.No, s.Score)
}

func main() {
	var student = Student{15, 23, 88.2}

	student2 := student	// 모든 필드 복사

	PrintStudent(student2)	// 함수 호출 시에도 구조체 복사
}