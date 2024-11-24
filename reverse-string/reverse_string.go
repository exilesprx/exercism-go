package reverse

func Reverse(input string) string {
	reversed := ""
	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		reversed += string(runes[i])
	}

	return reversed
}
