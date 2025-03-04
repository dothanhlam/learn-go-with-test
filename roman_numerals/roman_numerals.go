package roman_numerals

func ConvertToRoman(arabic int) string {
	var result string
	switch arabic {
		case 1:
			result = "I"
		case 2:
			result = "II"
		case 3:
			result = "III"
		default:
			result = ""
	}		
	return  result
}
