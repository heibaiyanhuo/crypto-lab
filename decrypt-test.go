package main

import (
	"bytes"
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
)

var encK = initKey()
var macK = initKey()

func initKey() []byte {
	ks := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	keyBytes, _ := hex.DecodeString(ks)
	return keyBytes
}

func hs(key, messgae []byte) []byte {
	inner := sha256.New()
	outer := sha256.New()
	ipad := make([]byte, 64)
	opad := make([]byte, 64)
	copy(ipad, key)
	copy(opad, key)
	for i := range ipad {
		ipad[i] ^= 0x36
	}
	for i := range opad {
		opad[i] ^= 0x5c
	}
	inner.Write(ipad)
	inner.Write(messgae)
	in := inner.Sum(nil)
	outer.Reset()
	outer.Write(opad)
	outer.Write(in)
	return outer.Sum(nil)
}

func cbc(key, IV, C []byte) []byte {
	block, _ := aes.NewCipher(key)
	numOfBlocks := len(C) / 16
	plaintextBytes := make([]byte, len(C))
	dst := plaintextBytes
	src := C
	iv := IV
	for i := 0; i < numOfBlocks; i++ {
		block.Decrypt(dst[:16], src[:16])
		for j := 0; j < 16; j++ {
			dst[j] ^= iv[j]
		}
		iv = src[:16]
		dst = dst[16:]
		src = src[16:]
	}
	return plaintextBytes
}

func aesDecResult(C []byte) string {
	IV := C[:16]
	C = C[16:]
	M0 := cbc(encK, IV, C)
	l := len(M0)
	n := int(M0[l-1])
	if n == 0 || l < n {
		return "INVALID PADDING"
	}
	for i := 2; i <= n; i++ {
		if int(M0[l-i]) != n {
			return "INVALID PADDING"
		}
	}
	M0 = M0[:l-n]
	if len(M0) < 32 {
		return "INVALID MAC"
	}
	M := M0[:len(M0)-32]
	T := M0[len(M0)-32:]
	expectedT := hs(macK, M)
	if bytes.Equal(T, expectedT) {
		return "SUCCESS"
	}
	return "INVALID MAC"
}

func main() {
	inputFile := flag.String("i", "", "")
	flag.Parse()
	data, _ := ioutil.ReadFile(*inputFile)
	fmt.Println(aesDecResult(data))
}
