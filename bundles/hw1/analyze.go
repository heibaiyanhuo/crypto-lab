package main

import (
	"crypto-lab/vigenere"
	"fmt"
	"os"
	"strconv"
)

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
	vigenereAnalyze()
}
