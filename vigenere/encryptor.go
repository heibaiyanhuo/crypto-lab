package vigenere

import (
	"crypto-lab/utils"
	"strings"
)

func encrypt(key string, plaintext []byte) string {
	var b strings.Builder
	for i := 0; i < len(plaintext); i++ {
		messageByte := plaintext[i]
		if 97 <= messageByte && messageByte < 123 {
			messageByte -= 32
		}
		if 65 <= messageByte && messageByte < 91 {
			b.WriteByte((messageByte - 65 + key[i % len(key)] - 65) % 26 + 65)
		}
	}
	return b.String()
}

func Encrypt(key string, filename string) string {
	return encrypt(key, utils.GetContentFromFile(filename))
}
