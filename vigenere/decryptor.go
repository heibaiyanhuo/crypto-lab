package vigenere

import (
	"crypto-lab/utils"
	"strings"
)

func decrypt(key string, ciphertext []byte) string {
	var b strings.Builder
	for i := 0; i < len(ciphertext); i++ {
		messageByte := ciphertext[i]
		if 97 <= messageByte && messageByte < 123 {
			messageByte -= 32
		}
		if 65 <= messageByte && messageByte < 91 {
			b.WriteByte((messageByte - 65 + (91 - key[i % len(key)])) % 26 + 65)
		}
	}
	return b.String()
}

func Decrypt(key string, filename string) string {
	return decrypt(key, utils.GetContentFromFile(filename))
}
