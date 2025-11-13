package main

func ToUpper(str string) string {
	result := ""

	for _, r := range str {
		if r >= 'a' && r <= 'z' {
			r = r - 32
		}
		result += string(r)
	}
	return result
}