package paths

// Unique finds the unique paths between the top left and bottom rigth of a matrix
func Unique(m, n int) int {

	// this is how to make a 2d array in Go
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		mem[i][0] = 1
	}

	for i := 0; i < n; i++ {
		mem[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			mem[i][j] = mem[i-1][j] + mem[i][j-1]
		}
	}
	return mem[m-1][n-1]
}
