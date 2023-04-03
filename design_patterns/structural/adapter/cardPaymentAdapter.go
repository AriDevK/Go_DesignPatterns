package main

type CardPaymentAdapter struct {
	CardPayment *CardPayment
	BankAccount string
}

func (cpa CardPaymentAdapter) Pay() {
	cpa.CardPayment.Pay(cpa.BankAccount)
}
