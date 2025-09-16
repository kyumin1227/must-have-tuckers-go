package main

import "testing"

func TestAtoi1(t *testing.T) {
	n, err := Atoi("0")
	if err != nil {
		t.Fail()
	}
	if n != 0 {
		t.Fail()
	}
}

func TestAtoi2(t *testing.T) {
	n, err := Atoi("-123")
	if err != nil {
		t.Fail()
	}
	if n != -123 {
		t.Fail()
	}
}

func TestAtoi3(t *testing.T) {
	n, err := Atoi("  123  ")
	if err != nil {
		t.Fail()
	}
	if n != 123 {
		t.Fail()
	}
}
