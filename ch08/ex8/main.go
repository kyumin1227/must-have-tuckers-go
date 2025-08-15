package main

import "fmt"

// 이 코드는 색상이 추가되면 Switch case 문도 수정해야 하므로
// 열거값에 연관된 Switch case는 최대한 줄이는 것이 좋음
type ColorType int
const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Yellow
}

func main() {
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}