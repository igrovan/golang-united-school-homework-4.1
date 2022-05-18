package reverse_string

func ReverseString(input string) (output string) {
	for i := range input {
		output += string(input[len(input)-i-1])
	}
	return output
}
