package main

import "fmt"

// func main() {
// 	CountDown(5)
// }

func CountDown(n int) {
	for i := n; i >= 1; i-- {
		fmt.Println(i)
	}
}