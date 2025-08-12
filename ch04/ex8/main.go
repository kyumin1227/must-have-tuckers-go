package main

import (
	"fmt"
	"math"
)

// 가장 작은 비트값만큼의 오차만 인정
func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b	// x 에서 y를 향해 1비트를 증가시키거나 감소 시킴
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a + b)
	fmt.Printf("%0.18f + %0.18f = %v\n", c, a + b, equal(a + b, c))

	a = 0.0000000000004
	b = 0.0000000000002
	c = 0.0000000000007

	fmt.Printf("%g == %g : %v\n", c, a + b, equal(a + b, c))
}