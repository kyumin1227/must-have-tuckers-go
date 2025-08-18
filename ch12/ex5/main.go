package main

import "fmt"

type User struct {
	Name string
	Age int
}

// 함수 외부로 공개되는 인스턴스는 함수가 종료되어도 사라지지 않음
func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u	// u는 탈출하여 메모리가 사라지지 않음
}

func main() {
	userPointer := NewUser("AAA", 23)

	fmt.Println(userPointer)
}