package main

import "fmt"

// func main() {
// 	isEven(24)
// }

func isEven(n int) {
	if n%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
