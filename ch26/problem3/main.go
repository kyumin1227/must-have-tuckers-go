package main

import (
	"errors"
	"strings"
)

func Atoi(input string) (int, error) {
	result := 0
	negative := false
	input = strings.TrimSpace(input)
	if input[0] == '-' {
		negative = true
		input = input[1:]
	}
	for _, c := range input {
		if c >= '0' && c <= '9' {
			result *= 10
			result += int(c - '0')
		} else {
			return 0, errors.New("cannot convert to int")
		}
	}
	if negative {
		result *= -1
	}
	return result, nil
}
