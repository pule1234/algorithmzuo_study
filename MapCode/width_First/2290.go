package main

import "math"

func minimumObstacles(grid [][]int) int {
	move := []int{-1, 0, 1, 0, -1}
	n := len(grid)
	m := len(grid[0])

	distance := make([][]int, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]int, m)
		for j := 0; j < m; j++ {
			distance[i][j] = math.MaxInt32
		}
	}

	deque := make([][]int, 0)
	deque = append(deque, []int{0, 0})
	distance[0][0] = 0

	for len(deque) > 0 {
		record := deque[0]
		deque = deque[1:]
		x := record[0]
		y := record[1]

		if x == m-1 && y == n-1 {
			return distance[x][y]
		}

		for i := 0; i < 4; i++ {
			nx := x + move[i]
			ny := y + move[i+1]

			if nx >= 0 && nx < n && ny >= 0 && ny < m && distance[x][y]+grid[nx][ny] < distance[nx][ny] {
				distance[nx][ny] = distance[x][y] + grid[nx][ny]
				if grid[nx][ny] == 0 {
					deque = append([][]int{{nx, ny}}, deque...)
				} else {
					deque = append(deque, []int{nx, ny})
				}
			}
		}
	}

	return -1

}
