package basic

import (
	"fmt"
	// "strconv"
)

func Palinornot(x int) bool {

	// a := strconv.Itoa(x)
	y := 0
	z := x

	for {
		z := x % 10
		y = y*10 + z
		x = x / 10
		if x == 0 {
			break
		}
	}
	fmt.Println(z, y)
	if z == y {
		return true
	} else {
		return false
	}
}
