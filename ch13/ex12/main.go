package main

import "fmt"

func main() {
	str1 := "BBB"
	str2 := "aaaaAAA"
	str3 := "BBAD"
	str4 := "ZZZ"

	// 문자열 대소 비교 시 문자열 길이와 상관 없이 앞글자 비교
	fmt.Printf("%s > %s : %v\n", str1, str2, str1 > str2)
	fmt.Printf("%s < %s : %v\n", str1, str3, str1 < str3)
	fmt.Printf("%s <= %s : %v\n", str1, str4, str1 > str4)
}