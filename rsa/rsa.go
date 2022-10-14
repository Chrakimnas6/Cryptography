// Reference: https://www.khanacademy.org/computing/computer-science/cryptography/modern-crypt/v/rsa-encryption-part-4
package rsa

import (
	"fmt"
	"math/big"
)

type privateKey struct {
	n, d int
}

type publicKey struct {
	n, e int
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func generateKeys(p, q int) (publicKey, privateKey) {
	n := p * q
	phi := (p - 1) * (q - 1)
	e := 3
	for e < phi {
		if gcd(e, phi) == 1 && e%2 != 0 {
			break
		}
		e++
	}
	d := ((e-1)*phi + 1) / e
	return publicKey{n, d}, privateKey{n, e}
}

func encrypt(msg string, publicKey publicKey) string {
	n, e := publicKey.n, publicKey.e
	var encrypted string
	for _, c := range msg {
		bigChar := big.NewInt(int64(c))
		bigN := big.NewInt(int64(n))
		bigE := big.NewInt(int64(e))

		encrypted += string(rune(bigChar.Exp(bigChar, bigE, bigN).Int64()))
	}
	fmt.Println(encrypted)
	return encrypted
}

func decrypt(msg string, privateKey privateKey) string {
	n, d := privateKey.n, privateKey.d
	var decrypted string
	for _, c := range msg {
		bigChar := big.NewInt(int64(c))
		bigN := big.NewInt(int64(n))
		bigD := big.NewInt(int64(d))

		decrypted += string(rune(bigChar.Exp(bigChar, bigD, bigN).Int64()))
	}
	fmt.Println(decrypted)
	return decrypted
}
