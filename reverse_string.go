package reverse_string

func ReverseString(input string) (output string) {
	for i := range input {
		output += string(rune(input[len(input)-i-1]))
	}
	return output
}
