package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	number := 17

	if isPrime(number) {
		fmt.Printf("%d является простым числом.", number)
	} else {
		fmt.Printf("%d не является простым числом.", number)
	}
}
