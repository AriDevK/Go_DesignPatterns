package main

import "fmt"

type CardPayment struct {
}

func (c CardPayment) Pay(bankAccount string) {
	fmt.Printf("Paying using Card with account %s\n", bankAccount)
}
