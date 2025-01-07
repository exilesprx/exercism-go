package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	cipher := ""
	for _, r := range plain {
		if r >= 'a' && r <= 'z' {
			cipher += string(rune((int(r-'a')+shiftKey)%26 + 'a'))
			continue
		}

		if r >= 'A' && r <= 'Z' {
			cipher += string(rune((int(r-'A')+shiftKey)%26 + 'A'))
			continue
		}

		cipher += string(r)
	}

	return cipher
}
