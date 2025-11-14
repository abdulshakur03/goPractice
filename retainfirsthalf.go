package main

func RetainFirstHalf(str string) string {
	strLength := len(str)
	if strLength == 0 || strLength < 1{
		return str
	}
	hlfStr := strLength/2
	return str[:hlfStr]	
}
