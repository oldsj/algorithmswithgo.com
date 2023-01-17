package module01

import "fmt"

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
func FizzBuzz(n int) {
	var fizzbuzz string
	for i := 1; i <= n; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fizzbuzz += "Fizz Buzz"
		} else if i % 3 == 0 {
			fizzbuzz += "Fizz"
		} else if i % 5 == 0 {
			fizzbuzz += "Buzz"
		} else {
			fizzbuzz += fmt.Sprintf("%d", i)
		}

		if i == n {
			fmt.Println(fizzbuzz)
		} else {
			fizzbuzz += ", "
		}
	}
}
