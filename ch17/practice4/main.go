package main

import "time"

type Courier struct {
	Name string
}

type Product struct {
	Name 	string
	Price 	int
	ID 		int
}

type Parcel struct {
	Pdt			*Product
	SippedTime	time.Time
	DeliveredTime	time.Time
}

func (c *Courier) SendProduce(pdt *Product) *Parcel {
	parcel := new(Parcel)
	parcel.Pdt = pdt
	parcel.SippedTime = time.Now()

	return parcel
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}