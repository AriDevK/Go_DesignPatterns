package main

import (
	m "Go_DesignPatterns_Platzi/design_patterns/creational/factory/models"
	"fmt"
)

func GetComputerFactory(computerType string) (m.IProduct, error) {
	switch computerType {
	case "Laptop":
		return m.NewLaptop(), nil
	case "Desktop":
		return m.NewDesktop(), nil
	default:
		return nil, fmt.Errorf("invalid @computerType:%s", computerType)
	}
}

func PrintProductNameAndStock(p m.IProduct) {
	fmt.Printf("Product name: %s \t Stock: %d", p.GetName(), p.GetStock())
}
