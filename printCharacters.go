package main

import "fmt"

func PrintCharacters(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]))
	}
}