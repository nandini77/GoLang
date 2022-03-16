package main

import "fmt"

type bank interface {
	Account() int
}
type Deposit struct {
	balance   int
	depAmount int
}
type Withdraw struct {
	balance    int
	withAmount int
}

func (d Deposit) Account() int {
	d.balance = d.balance + d.depAmount
	return d.balance
}
func (w Withdraw) Account() int {
	if w.balance > 0 {
		w.balance = w.balance - w.withAmount
	}
	return w.balance
}
func main() {

	x := Deposit{50, 80}
	y := Withdraw{50, 40}
	var b bank
	b = x
	fmt.Println("Balance in account after deposit : ", b.Account())
	b = y
	fmt.Println("Balance in account after withdraw : ", b.Account())
}
