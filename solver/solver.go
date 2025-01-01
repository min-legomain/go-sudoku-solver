package solver

import (
	"errors"
	"fmt"
)

func Execute(grid [][]int) ([][]int, error) {
	for _, row := range grid {
		fmt.Println(row)
	}

	if !validateGrid(grid) {
		return nil, errors.New("invalid input form")
	}

	if solve(grid) {
		return grid, nil
	}
	return nil, errors.New("cannot solve")
}

func solve(grid [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if isAvailable(grid, row, col, num) {
						grid[row][col] = num

						if solve(grid) {
							return true
						}

						grid[row][col] = 0
					}
				}

				return false
			}
		}
	}

	return true
}

func isAvailable(grid [][]int, row, col, num int) bool {
	// 行を確認
	for x := 0; x < 9; x++ {
		if grid[row][x] == num {
			return false
		}
	}

	// 列を確認
	for x := 0; x < 9; x++ {
		if grid[x][col] == num {
			return false
		}
	}

	// 3*3のマスを確認
	// 起点となる行と列 = 行(列) - 行(列)の値を3で割ったあまり
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

func validateGrid(grid [][]int) bool {
	if len(grid) != 9 {
		return false
	}
	for _, row := range grid {
		if len(row) != 9 {
			return false
		}
	}
	return true
}