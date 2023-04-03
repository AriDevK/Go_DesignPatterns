package models

type Computer struct {
	name  string
	stock int
}

func (c *Computer) GetName() string {
	return c.name
}

func (c *Computer) SetName(name string) {
	c.name = name
}

func (c *Computer) GetStock() int {
	return c.stock
}

func (c *Computer) SetStock(stock int) {
	c.stock = stock
}
