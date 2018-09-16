package main

import (
	"crypto-lab/vigenere"
	"fmt"
	"os"
)

func vigenereEnc() {
	key := os.Args[1]
	filename := os.Args[2]
	fmt.Println(vigenere.Encrypt(key, filename))
}

func main() {
	vigenereEnc()
}
