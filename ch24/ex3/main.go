package main

type Integer interface {
	int8 | int16 | int32 | int64 | int // 여기에 ~ 를 추가, ~int
}

func add[T Integer](a, b T) T {
	return a + b
}

type MyInt int

func main() {
	add(1, 2)
	var a MyInt = 3
	var b MyInt = 5
	add(a, b) // 별칭은 Integer의 타입 제한에 포함 되지 않음, 포함 시키고 싶은 경우에는 타입 제한 시 앞에 ~ 를 추가
}
