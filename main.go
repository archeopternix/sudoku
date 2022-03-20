// sudoku project main.go
package main

import (
	"fmt"
)

type Board [9][9]byte

func check(board *Board) bool {
	// horizontal row
	for row := 0; row < 9; row++ {
		c := []bool{false, false, false, false, false, false, false, false, false}
		for col := 0; col < 9; col++ {
			val := board[row][col]
			if val == 0 {
				continue
			}
			if c[val-1] {
				return false
			}
			c[val-1] = true
		}
	}
	// vertical column
	for col := 0; col < 9; col++ {
		c := []bool{false, false, false, false, false, false, false, false, false}
		for row := 0; row < 9; row++ {
			val := board[row][col]
			if val == 0 {
				continue
			}
			if c[val-1] {
				return false
			}
			c[val-1] = true
		}
	}
	// box
	for boxrow := 0; boxrow < 3; boxrow++ {
		for boxcol := 0; boxcol < 3; boxcol++ {
			c := []bool{false, false, false, false, false, false, false, false, false}
			for row := 0; row < 3; row++ {
				for col := 0; col < 3; col++ {
					val := board[boxrow*3+row][boxcol*3+col]
					if val == 0 {
						continue
					}
					if c[val-1] {
						return false
					}
					c[val-1] = true
				}
			}
		}
	}
	return true
}

func printBoard(board *Board) {
	fmt.Println("-------------------------")
	for i, b := range board {
		fmt.Printf("|%d %d %d|%d %d %d|%d %d %d|\n", b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8])
		if (i % 3) == 2 {
			fmt.Println("-------------------------")
		}
	}
}

func main() {
	raster := Board{
		{0, 0, 9, 0, 3, 0, 5, 0, 0},
		{5, 0, 1, 0, 0, 0, 2, 0, 0},
		{0, 2, 0, 0, 5, 7, 0, 0, 6},
		{0, 0, 0, 7, 0, 0, 0, 0, 0},
		{0, 8, 0, 0, 0, 0, 0, 0, 0},
		{3, 5, 0, 0, 8, 2, 0, 9, 0},
		{0, 0, 0, 3, 0, 0, 0, 7, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 3},
		{0, 6, 0, 0, 9, 0, 0, 0, 1}}

	ok, brd := solve(&raster, 0)
	if ok {
		printBoard(brd)
	} else {
		fmt.Println("Keine Lösung")
	}

}

func solve(board *Board, pos byte) (bool, *Board) {
	var one byte = 1
	var i byte

	for board[pos%9][pos/9] > 0 {
		if pos > 79 {
			return true, board
		}
		pos += 1
	}

	if board[pos%9][pos/9] == 0 {
		for i = 0; i < 9; i++ {
			board[pos%9][pos/9] = i + one
			// if valid number
			if check(board) {
				// until not all numbers are solved
				if pos < 80 {
					brd := new(Board)
					*brd = *board
					ok, retboard := solve(brd, pos+1)
					if ok {

						return true, retboard
					}
				} else {
					return false, nil
				}
			}
		}
	}
	return false, nil
}
