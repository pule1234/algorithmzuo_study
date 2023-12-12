package main

import "math"

func minCost(grid [][]int) int {
	move := [][]int{{}, {0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m := len(grid)
	n := len(grid[0])
	distance := make([][]int, m)
	for i := 0; i < m; i++ {
		distance[i] = make([]int, n)
		for j := 0; j < n; j++ {
			distance[i][j] = math.MaxInt32
		}
	}

	q := make([][]int, 0)
	q = append(q, []int{0, 0})
	distance[0][0] = 0

	for len(q) > 0 {
		record := q[0]
		q = q[1:]
		x := record[0]
		y := record[1]

		if x == m-1 && y == n-1 {
			return distance[x][y]
		}
		for i := 1; i <= 4; i++ {
			nx := x + move[i][0]
			ny := y + move[i][1]
			var weight int
			if grid[x][y] == i {
				weight = 0
			} else {
				weight = 1
			}
			if nx >= 0 && nx < m && ny >= 0 && ny < n && distance[x][y]+weight < distance[nx][ny] {
				distance[nx][ny] = distance[x][y] + weight
				if weight == 0 {
					q = append([][]int{{nx, ny}}, q...)
				} else {
					q = append(q, []int{nx, ny})
				}
			}
		}
	}

	return -1

}
