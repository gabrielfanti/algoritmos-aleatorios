package exercicio32

import (
	"fmt"
)

func validaSudoku() bool {
	board := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	for i := 0; i < 9; i++ {
		seen := make(map[int]bool)
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num != 0 {
				if seen[num] {
					return false
				}
				seen[num] = true
			}
		}
	}

	for i := 0; i < 9; i++ {
		seen := make(map[int]bool)
		for j := 0; j < 9; j++ {
			num := board[j][i]
			if num != 0 {
				if seen[num] {
					return false
				}
				seen[num] = true
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			seen := make(map[int]bool)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					num := board[i+k][j+l]
					if num != 0 {
						if seen[num] {
							return false
						}
						seen[num] = true
					}
				}
			}
		}
	}

	return true
}

func Sudoku() {
	if validaSudoku() {
		fmt.Println("O tabuleiro de Sudoku é válido.")
	} else {
		fmt.Println("O tabuleiro de Sudoku contém valores incorretos.")
	}
}
