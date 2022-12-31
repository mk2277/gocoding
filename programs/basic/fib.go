package basic

import "fmt"

func Printfib(x int) {
	a := 0
	b := 1
	fmt.Print(a, b, " ")
	for i := 1; i < x; i++ {
		c := a + b
		fmt.Print(c, " ")
		a = b
		b = c
	}
}
