package editdistance

// MinDistance returns total changes to be made between words
func MinDistance(w1, w2 string) int {
	m := len(w1) + 1
	n := len(w2) + 1

	// make the dp table
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
	}

	// when second string is empty, difference is all first
	for i := 0; i < m; i++ {
		mem[i][0] = i
	}

	// when first string is empty, difference is all second
	for i := 0; i < n; i++ {
		mem[0][i] = i
	}

	// loop through checking string if it is the same, else 3 possible operations
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if w1[i-1] == w2[j-1] {
				mem[i][j] = mem[i-1][j-1]
			} else {
				mem[i][j] = 1 + min(mem[i-1][j], mem[i-1][j-1], mem[i][j-1])
			}
		}
	}

	return mem[m-1][n-1]
}

func min(a, b, c int) int {
	if a < b {
		return minHelper(a, c)
	}
	return minHelper(b, c)
}

func minHelper(a, b int) int {
	if a < b {
		return a
	}
	return b
}
