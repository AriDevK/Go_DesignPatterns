package main

func main() {
	cash := &CashPayment{}
	card := &CardPaymentAdapter{
		CardPayment: &CardPayment{},
		BankAccount: "12341234789",
	}

	ProcessPayment(cash)
	ProcessPayment(card)
}
