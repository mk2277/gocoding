package main

import (
	"fmt"
	"programs/basic"
)

func main() {

loop:
	for {
		a := 0
		fmt.Println(" ")
		fmt.Println("Enter choice:")
		fmt.Println("1)prime 2)palindrome 3)fib 4)exit")
		fmt.Scanln(&a)
		switch a {
		case 1:
			fmt.Println("enter number :")
			fmt.Scanln(&a)
			fmt.Println(basic.Primeornot(a))

		case 2:
			fmt.Println("enter number :")
			fmt.Scanln(&a)
			fmt.Println(basic.Palinornot(a))

		case 3:
			fmt.Println("enter number upto where to print :")
			fmt.Scanln(&a)
			basic.Printfib(a)

		case 4:
			break loop
		}
	}
}
