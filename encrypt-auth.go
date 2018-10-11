package main

import (
	"bytes"
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func hmacSha256(key, messgae []byte) []byte {
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

func generateIV() []byte {
	token := make([]byte, 16)
	rand.Seed(time.Now().UnixNano())
	rand.Read(token)
	return token
}

func aesCBCEnc(key, IV, M []byte) []byte {
	block, _ := aes.NewCipher(key)
	numOfBlocks := len(M) / 16
	cipherBytes := make([]byte, len(M))
	dst := cipherBytes
	src := M
	iv := IV
	for i := 0; i < numOfBlocks; i++ {
		for j := 0; j < 16; j++ {
			dst[j] = src[j] ^ iv[j]
		}
		block.Encrypt(dst[:16], dst[:16])
		iv = dst[:16]
		dst = dst[16:]
		src = src[16:]
	}
	return cipherBytes
}

func aesCBCDec(key, IV, C []byte) []byte {
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

func Encrypt(encK, macK, M []byte) []byte {
	T := hmacSha256(macK, M)
	var buffer bytes.Buffer
	buffer.Write(M)
	buffer.Write(T)
	n := buffer.Len() % 16
	if n == 0 {
		for i := 0; i < 16; i++ {
			buffer.WriteByte(byte(16))
		}
	} else {
		l := 16 - n
		for i := 0; i < l; i++ {
			buffer.WriteByte(byte(l))
		}
	}
	IV := generateIV()
	C := aesCBCEnc(encK, IV, buffer.Bytes())
	buffer.Reset()
	buffer.Write(IV)
	buffer.Write(C)
	return buffer.Bytes()
}

func Decrypt(encK, macK, C []byte) ([]byte, string) {
	IV := C[:16]
	C = C[16:]
	M0 := aesCBCDec(encK, IV, C)
	l := len(M0)
	n := int(M0[l-1])
	if n == 0 {
		return nil, "INVALID PADDING"
	}
	for i := 2; i <= n; i++ {
		if int(M0[l-i]) != n {
			return nil, "INVALID PADDING"
		}
	}
	M0 = M0[:l-n]
	M := M0[:len(M0)-32]
	T := M0[len(M0)-32:]
	expectedT := hmacSha256(macK, M)
	if bytes.Equal(T, expectedT) {
		return M, "SUCCESS"
	}
	return nil, "INVALID MAC"
}

func splitKey(combKey string) ([]byte, []byte) {
	combKeyBytes, _ := hex.DecodeString(combKey)
	return combKeyBytes[:16], combKeyBytes[16:]
}

func readFrom(filename string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return data
}

func writeIn(filename string, content []byte) {
	ioutil.WriteFile(filename, content, 0644)
}

func main() {
	mode := os.Args[1]
	argSet := flag.NewFlagSet("", flag.ExitOnError)
	key := argSet.String("k", "", "Enc key and MAC key")
	inputFile := argSet.String("i", "", "input file")
	outputFile := argSet.String("o", "", "output file")
	argSet.Parse(os.Args[2:])
	encK, macK := splitKey(*key)
	if mode == "encrypt" {
		data := Encrypt(encK, macK, readFrom(*inputFile))
		writeIn(*outputFile, data)
	} else {
		data, e := Decrypt(encK, macK, readFrom(*inputFile))
		if data == nil {
			writeIn(*outputFile, []byte(e))
		} else {
			writeIn(*outputFile, data)
		}
	}
}
