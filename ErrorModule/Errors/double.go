package PointerAndError

import "fmt"

func DoublePointer() {

	var a int = 65     // declaring a variable
	var p1 *int = &a   // assigning pointer to variable
	var p2 **int = &p1 // assigning double pointer

	fmt.Println("Variable is : ", a)
	fmt.Println("Address of variable is : ", p1)
	fmt.Println("Address of the pointer is : ", p2)

}

func PanicMessage() {
	// This is panic function
	fmt.Println("Enter value for a:")
	var a int
	fmt.Scan(&a)
	if a != 22 {
		panic("value assigned is wrong")
	} else {
		fmt.Println("input accepted")
	}
}

// function for nil pointer
func NilPointer() {
	fmt.Println("Input for nil pointer")
	var x *int
	var y *int
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Println(x, y)

}

// example for pointer arthimetic
func PointerArithmetic() {
	fmt.Println("Enter x,y values for pointer arithmetic function")
	var x int
	var y int
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Println(x, y)
	p := &x
	q := &y
	fmt.Println(p) //Gives address of a variable p
	fmt.Println(q) //Gives address of a variable q

}
