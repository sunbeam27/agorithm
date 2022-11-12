package main

func main() {
	uniquePathsWithObstacles([][]int{{0, 0}})
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}
	for i := 0; i < m && obstacleGrid[i][0] != 1; i++ {
		f[i][0] = 1
	}
	for j := 0; j < n && obstacleGrid[0][j] != 1; j++ {
		f[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				f[i][j] = f[i][j-1] + f[i-1][j]
			}
		}
	}
	return f[m-1][n-1]
}
