package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	rArr := []rune(input)
	i := 0
	j := len(rArr) - 1
	for i < j {
		rArr[i], rArr[j] = rArr[j], rArr[i]
		i++
		j--
	}
	output = string(rArr)
	return output
}
