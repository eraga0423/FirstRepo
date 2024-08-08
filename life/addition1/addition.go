package addition

import (
	"fmt"
)

// greenday shows the grid; # - alive cell, . - dead cell

func Greenday(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == '#' {
				fmt.Print("x ")
			} else {
				fmt.Print("· ")
			}
		}
		fmt.Println()
	}
}

// thenbhd counts alive neighbours
func Thenbhd(grid [][]rune, row, col int) int {
	// Массив смещений для проверки вокруг клетки (всего 8 направлений).
	directions := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	count := 0
	for _, dir := range directions {
		r, c := row+dir[0], col+dir[1]
		if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == '#' {
			count++
		}
	}
	return count
}

// mindlessselfindulgence generates new state of grid and returns new state of grid
func MindlessSelfIndulgence(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for i := range grid {
		newGrid[i] = make([]rune, len(grid[i]))
		for j := range grid[i] {
			liveNeighbors := Thenbhd(grid, i, j)
			if grid[i][j] == '#' {
				if liveNeighbors < 2 || liveNeighbors > 3 {
					newGrid[i][j] = '.'
				} else {
					newGrid[i][j] = '#'
				}
			} else { // if cell is dead .
				if liveNeighbors == 3 {
					newGrid[i][j] = '#'
				} else {
					newGrid[i][j] = '.'
				}
			}
		}
	}
	return newGrid
}

// radiohead checks alive cells on the grid
func Radiohead(grid [][]rune) bool {
	for _, row := range grid {
		for _, cell := range row {
			if cell == '#' {
				return true
			}
		}
	}
	return false
}

// deftones counts alive cells
func Deftones(grid [][]rune) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == '#' {
				count++
			}
		}
	}
	return count
}
