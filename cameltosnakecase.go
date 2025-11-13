package main

func CamelToSnakeCase(str string) string {
	result := ""
	for _, r := range str {
		if r >= 'A' && r <= 'Z' {
			result += "_"
			r = r + 32
		}
		 result += string(r)
	}

	return result
}
