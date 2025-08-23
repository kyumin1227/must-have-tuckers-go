package main

import (
	"ch18/fedex"
	koreapost "ch18/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &koreapost.PostSender{}
	SendBook("어린 왕자", sender)	// fedex와 koreaPost는 서로 타입이 달라서 에러 발생
	SendBook("그리스인 조르바", sender)
}