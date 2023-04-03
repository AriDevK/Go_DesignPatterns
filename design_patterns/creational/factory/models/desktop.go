package models

type Desktop struct {
	Computer
}

func NewDesktop() IProduct {
	return &Desktop{
		Computer{
			name:  "DELL Desktop Computer",
			stock: 100,
		}}
}
