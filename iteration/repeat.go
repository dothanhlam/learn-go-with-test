package iteration

func Repeat(character string, n int) string {
	if n < 1 {
		return character
	}
	str:= ""
	for i:= 0; i < n; i ++ {
		str = str + character
	}
	return str
}

func IterationAsWhile(character string, n int) string {
	if n < 1 {
		return character
	}
	str:= ""
	for n > 0 {
		str = str + character
		n--
	}
	return str
}
