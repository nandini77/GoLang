package PointerAndError

import "fmt"

type Person struct {
	balance    int
	depAmount  int
	withAmount int
}

func Deposit() {
	var p Person
	p.balance = 100
	fmt.Println("Enter amount to deposit")
	fmt.Scan(&p.depAmount)
	p.balance = p.balance + p.depAmount
	fmt.Println(p.balance)
	if p.depAmount > 10000 {
		panic("Amount exceeded")
	}
}

func Withdraw() {
	var p Person
	p.balance = 1000
	fmt.Println("Enter amount to withdraw")
	fmt.Scan(&p.withAmount)
	if p.balance > 0 {
		p.balance = p.balance - p.withAmount
		fmt.Println(p.balance)
	}
	if p.withAmount > p.balance {
		panic("Amount entered is invalid")
	}
}

func DeferDemo() {

	fmt.Println("1.Enter 1 for deposit\n2.Enter 2 for withdrawal")
	var option int
	fmt.Scan(&option)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	switch option {
	case 1:
		Deposit()
	case 2:
		Withdraw()
	case 3:
		defer fmt.Println("Entered option is 3")

	}

}
