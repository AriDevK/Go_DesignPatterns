package models

type Laptop struct {
	Computer
}

func NewLaptop() IProduct {
	return &Laptop{
		Computer{
			name:  "Black HP Laptop",
			stock: 15,
		}}
}
