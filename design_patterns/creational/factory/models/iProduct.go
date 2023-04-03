package models

type IProduct interface {
	GetStock() int
	GetName() string
	SetStock(stock int)
	SetName(name string)
}
