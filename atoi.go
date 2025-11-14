// package main

// func Atoi(s string) int {
// 	num := 0
// 	sign := 1

// 		if s[0] == '-' {
// 		sign = -1
// 	}else if s[0] == '+'{
// 		sign = 1
// 	}
// 		for _, r := range s {
// 			if r >= '0' && r <= '9' {
// 				digits := int(r - '0')
// 				num = num*10 + digits
// 			}
// 		}

// return num*sign
// }

package main


func Atoi(s string) int {
	if len(s) == 0 { // empty string → return 0
		return 0
	}

	num := 0
	sign := 1
	start := 0 // start index to read digits

	// Handle optional sign
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	// Loop through characters
	for i := start; i < len(s); i++ {
		r := s[i]
		if r >= '0' && r <= '9' {
			digit := int(r - '0') // convert rune to integer
			num = num*10 + digit  // build the number
		} else {
			// Stop at first non-digit (Z01 allows this)
			break
		}
	}

	return num * sign
}




