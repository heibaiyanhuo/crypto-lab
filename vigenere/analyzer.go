package vigenere

import (
	"crypto-lab/calculation"
	"crypto-lab/utils"
	"math"
	"strings"
)

const MAX_KEY_SIZE int = 20
const IC_ENGILISH float64 = 0.066

func estimateKeyLength(cipher []byte) int {
	L := 0
	diff := 10.0
	for i := 1; i <= MAX_KEY_SIZE; i++ {
		currIC := calculation.CalcAvgIC(i, cipher)
		currDiff := math.Abs(currIC - IC_ENGILISH)
		if currDiff < diff && math.Abs(currDiff - diff) > 0.001 {
			diff = currDiff
			L = i
		}
	}
	return L
}

func recoverKey(cipher []byte, keyLen int) string {
	var b strings.Builder
	for i := 0; i < keyLen; i++ {
		cipherE := calculation.FindMostCommonChar(i, keyLen, cipher)
		if cipherE < 69 {
			b.WriteByte(cipherE + 22)
		} else {
			b.WriteByte(cipherE - 4)
		}
	}
	return b.String()
}

func EstimateKeyLen(filename string) int{
	return estimateKeyLength(utils.GetUpperCaseContentFromFile(filename))
}

func RecoverKey(filename string, keyLen int) string {
	return recoverKey(utils.GetUpperCaseContentFromFile(filename), keyLen)
}
