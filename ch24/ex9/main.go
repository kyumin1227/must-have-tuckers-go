package main

import (
	"fmt"
	"slices"
)

func main() {
	// BinarySearch 찾으면 index 반환 후 true, 못 찾으면 있어야 할 index 반환 후 false
	names := []string{"Alice", "Bob", "Vera"}
	n, found := slices.BinarySearch(names, "Vera")
	fmt.Println("Vera:", n, found)
	n, found = slices.BinarySearch(names, "Bill")
	fmt.Println("Bill:", n, found)
}
