package main

import "fmt"

type (
	Payment interface {
		Pay()
	}
	CashPayment struct{}
	BankPayment struct{}
)

func ProcessPayment(p Payment) {
	p.Pay()
}

func (CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using Bankaccount %d\n", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	bpa := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bpa)
}
