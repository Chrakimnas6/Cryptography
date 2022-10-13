package shift

func Encrypt(msg string, key int) string {
	var encrypted string
	for _, c := range msg {
		encrypted += string(c + rune(key)%26)
	}
	return encrypted
}

func Decrypt(msg string, key int) string {
	var decrypted string
	for _, c := range msg {
		decrypted += string(c - rune(key)%26)
	}
	return decrypted
}
