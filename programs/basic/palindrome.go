package basic

import (
	"fmt"
	"strconv"
)

func Palinornot(x int) string {

	a := strconv.Itoa(x)
	y := 0
	z := x

	for i := 0; i < len(a); i++ {
		z := x % 10
		y = y*10 + z
		x = x / 10
	}
	fmt.Println(z, y)
	if z == y {
		return "yes its palindrome"
	} else {
		return "not a palindrome"
	}
}
