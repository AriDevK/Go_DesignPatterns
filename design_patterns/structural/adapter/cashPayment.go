package main

import "fmt"

type CashPayment struct {
}

func (c CashPayment) Pay() {
	fmt.Println("Paying with cash money")
}
