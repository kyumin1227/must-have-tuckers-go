package main

import "fmt"

type Data struct {
	value	int
	data	[200]int
}

func ChangeData(arg *Data) {	// Data 포인터를 받음
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(&data)	// data의 주소를 넘김
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}