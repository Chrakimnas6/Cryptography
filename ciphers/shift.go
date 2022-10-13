package shift

import "fmt"

func Encrypt(msg string, key int) (string, error) {
	if key < 0 || key > 26 {
		return "", fmt.Errorf("key must be between 0 and 26")
	}
	var encrypted string
	for _, c := range msg {
		encrypted += string(c + rune(key)%26)
	}
	return encrypted, nil
}

func Decrypt(msg string, key int) (string, error) {
	var decrypted string
	if key < 0 || key > 26 {
		return "", fmt.Errorf("key must be between 0 and 26")
	}
	for _, c := range msg {
		decrypted += string(c - rune(key)%26)
	}
	return decrypted, nil
}
