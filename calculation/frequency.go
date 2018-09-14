package calculation

func calcIC(N int, freq [26]int) float64 {
	IC := 0.0
	for i := 0; i < 26; i++ {
		IC += float64(freq[i] * (freq[i] - 1))
	}
	IC /= float64(N * (N - 1))
	return IC
}

func CalcAvgIC(L int, msg []byte) float64 {
	IC := 0.0
	count := 0
	for i := 0; i < L; i++ {
		var freq [26]int
		count = 0
		for j := i; j <= len(msg) - L; j += L {
			freq[msg[j] - 65]++
			count++
		}
		IC += calcIC(count, freq)
	}
	return IC / float64(L)
}

func FindMostCommonChar(start int, L int, msg []byte) byte {
	var freq [26]int
	var target byte = 65
	maxFreq := 0
	for i := start; i <= len(msg) - L; i += L {
		idx := msg[i] - 65
		freq[idx]++
		if freq[idx] > maxFreq {
			maxFreq = freq[idx]
			target = msg[i]
		}
	}
	return target
}
