package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var reverse string
	for _, r := range word {
		reverse = string(r) + reverse
	}
	return reverse
}
