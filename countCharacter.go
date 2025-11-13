package main

func CountChar(str string, c rune) int {
    var count int

		for _, r := range str {
			if r == c {
				count++
			}
		}
		return count
}