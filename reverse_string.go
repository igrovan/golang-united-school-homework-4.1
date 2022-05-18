package reverse_string

import "strings"

func ReverseString(input string) (output string) {
	var builder strings.Builder
	inputRunes := []rune(input)

	for i := len(inputRunes) - 1; i >= 0; i-- {
		builder.WriteRune(inputRunes[i])
	}
	output = builder.String()
	return output
}
