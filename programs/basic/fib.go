package basic

import "fmt"

func Printfib(x int) {

	var myslice = []int{0, 1}

	for i := 0; i < x-2; i++ {
		result := myslice[i] + myslice[i+1]
		myslice = append(myslice, result)
	}
	fmt.Println(myslice)
}
