package module01

import "fmt"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var reverse string
	for i := range word {
		reverse += string(word[len(word)-i-1])
	}
	return reverse
}
