package main

import "fmt"

type Payment interface {
	Pay()
}
type CashPayment struct {
}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash")
}
func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct {
}

func (BankPayment) Pay(bankAccount int) {
	fmt.Println("Payment using bank account ", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (b *BankPaymentAdapter) Pay() {
	b.BankPayment.Pay(b.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	/*bank :=&BankPayment{}
	ProcessPayment(bank)*/
	bpa := &BankPaymentAdapter{BankPayment: &BankPayment{}, bankAccount: 123}
	ProcessPayment(bpa)
}
