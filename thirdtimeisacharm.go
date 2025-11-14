package main

func ThirdTimeIsACharm(str string) string {
	result := ""
	if str == "" || len(str) < 3{ 
			return "\n"
		}
	for i := 0; i < len(str); i++ {
		if (i+1)%3 == 0 {
			result += string(str[i])
		}
	}
	return result + "\n"
}