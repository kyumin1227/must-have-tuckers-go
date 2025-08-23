package main

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker	// 기본값 nil
	att.Attack()		// att가 nil이기 때문에 런타임 에러 발생
}