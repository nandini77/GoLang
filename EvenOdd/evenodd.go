package main

import "fmt"

func main() {

	var a [10]int
	var i int
	for i = 0; i < 10; i++ {
		fmt.Scan(&a[i])
		if a[i]%2 == 0 {
			fmt.Println("Number is Even")
		} else {
			fmt.Println("Number is Odd")
		}
	}

	// for i = 0; i < 10; i++ {
	// 	if a[i]%2 == 0 {
	// 		fmt.Println("Number is Even")
	// 	} else {
	// 		fmt.Println("Number is Odd")
	// 	}

	// }
}
