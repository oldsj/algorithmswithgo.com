package module01

import (
	"strconv"
	"fmt"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	var result string
	var rem string
	for i := dec; i >= 0; i = i/base {
		rem = strconv.Itoa(i%base)
		fmt.Printf("i: %d rem: %s\n", i, rem)

		switch rem {
		case "10":
			rem = "A"
		case "11":
			rem = "B"
		case "12":
			rem = "C"
		case "13":
			rem = "D"
		case "14":
			rem = "E"
		case "15":
			rem = "F"
		}

		if i == 0 {
			return result
		} else {
			result = rem + result
		}
	}
	return result
}
