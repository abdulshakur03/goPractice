package main

func IsAlpha (str string) bool{
	for _, r := range str {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			return true
		}
	}
	return false 
}