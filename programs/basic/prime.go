package basic

// import "fmt"

func Primeornot(x int) string {
	count := 0
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			count++
			// fmt.Println("count")
		}
	}
	if count > 2 {
		return "no not a prime"
	} else {
		return "yes it is prime"
	}

}
