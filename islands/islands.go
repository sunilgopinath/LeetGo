package islands

// Count gets the number of islands in the grid
func Count(grid [][]byte) int {
	// calculate grid dimensions
	tr := len(grid)
	tc := len(grid[0])

	// create 2d slice
	visited := make([][]bool, tr)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, tc)
	}

	dx := []int{0, -1, 1, 0}
	dy := []int{-1, 0, 0, 1}
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				dfs(grid, visited, i, j, tr, tc, dx, dy)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, visited [][]bool, row, col, tr, tc int, dx, dy []int) {
	visited[row][col] = true
	for i := 0; i < 4; i++ {
		r := row + dx[i]
		c := col + dy[i]
		if isSafe(r, c, tr, tc) && !visited[r][c] && grid[r][c] == '1' {
			dfs(grid, visited, r, c, tr, tc, dx, dy)
		}
	}
}

func isSafe(row, col, tr, tc int) bool {
	return row >= 0 && row < tr && col >= 0 && col < tc
}
