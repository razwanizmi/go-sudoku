package main

import (
	"fmt"
	"strconv"
)

type board [][]int

func newBoard(boardString string) [][]int {
	result := board{}

	for i := range boardString {
		rowInt := []int{}
		if i%9 == 0 {
			rowStr := boardString[i : i+9]
			for _, numRune := range rowStr {
				numStr := fmt.Sprintf("%c", numRune)
				num, err := strconv.Atoi(numStr)
				if err == nil {
					rowInt = append(rowInt, num)
				}
			}
			result = append(result, rowInt)
		}
	}

	return result
}

func (b board) saveEmptyPositions() [][]int {
	emptyPositions := [][]int{}

	for y := range b {
		for x := range b[y] {
			if b[y][x] == 0 {
				emptyPositions = append(emptyPositions, []int{y, x})
			}
		}
	}

	return emptyPositions
}

func (b board) checkRow(row int, value int) bool {
	for x := range b[row] {
		if b[row][x] == value {
			return false
		}
	}

	return true
}

func (b board) checkColumn(column int, value int) bool {
	for y := range b {
		if b[y][column] == value {
			return false
		}
	}

	return true
}

func (b board) checkSquare(y int, x int, value int) bool {
	yCorner := 0
	xCorner := 0
	squareSize := 3

	for y >= yCorner+squareSize {
		yCorner += squareSize
	}

	for x >= xCorner+squareSize {
		xCorner += squareSize
	}

	for i := range b[yCorner : yCorner+squareSize] {
		for j := range b[i][xCorner : xCorner+squareSize] {
			if b[i][j] == value {
				return false
			}
		}
	}

	return true
}

func (b board) checkValue(y int, x int, value int) bool {
	return b.checkColumn(y, value) && b.checkRow(x, value) && b.checkSquare(y, x, value)
}
