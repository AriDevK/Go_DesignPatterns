package main

type IPayment interface {
	Pay()
}

func ProcessPayment(p IPayment) {
	p.Pay()
}
