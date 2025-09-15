package main

import "fmt"

type NodeType1 struct {
	val  interface{}
	next *NodeType1
}

type NodeType2[T any] struct {
	val  T
	next *NodeType2[T]
}

func main() {
	node1 := &NodeType1{val: 1}
	node2 := &NodeType2[int]{val: 2}

	var v1 int = node1.val // 빈 인터페이스의 경우 값을 사용할 때 타입 변환이 필요
	fmt.Println(v1)
	var v2 int = node2.val
	fmt.Println(v2)
}
