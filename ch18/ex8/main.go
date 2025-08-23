package main

type Stringer interface {
	String() string
}

type Student struct {}

func main() {
	var stringer Stringer
	stringer.(*Student)	// Student 구조체가 String 메소드를 구현하고 있지 않기에 컴파일 에러
}