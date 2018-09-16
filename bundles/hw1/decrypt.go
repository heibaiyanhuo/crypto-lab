package main

import (
	"crypto-lab/vigenere"
	"fmt"
	"os"
)

func vigenereDec() {
	key := os.Args[1]
	filename := os.Args[2]
	fmt.Println(vigenere.Decrypt(key, filename))
}

func main() {
	vigenereDec()
}
