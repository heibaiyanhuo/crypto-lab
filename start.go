package main

import (
	"crypto-lab/vigenere"
	"fmt"
	"os"
	"strconv"
)

func vigenereEnc() {
	key := os.Args[1]
	filename := os.Args[2]
	fmt.Println(vigenere.Encrypt(key, filename))
}

func vigenereDec() {
	key := os.Args[1]
	filename := os.Args[2]
	fmt.Println(vigenere.Decrypt(key, filename))
}

func vigenereAnalyze() {
	filename := os.Args[1]
	if len(os.Args) > 2 {
		keyLen, _ := strconv.Atoi(os.Args[2])
		fmt.Println(vigenere.RecoverKey(filename, keyLen))
	} else {
		fmt.Println(vigenere.EstimateKeyLen(filename))
	}
}

func main() {
	//vigenereEnc()
	//vigenereDec()
	vigenereAnalyze()
}
