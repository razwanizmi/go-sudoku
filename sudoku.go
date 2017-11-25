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
				emptyPositions = append(emptyPositions, []int{x, y})
			}
		}
	}

	return emptyPositions
}

func (b board) checkXAxis(y int, value int) bool {
	for _, currentValue := range b[y] {
		if currentValue == value {
			return false
		}
	}

	return true
}

func (b board) checkYAxis(x int, value int) bool {
	for _, currentAxis := range b {
		if currentAxis[x] == value {
			return false
		}
	}

	return true
}

func (b board) checkSubSquare(x int, y int, value int) bool {
	xBorder := 0
	yBorder := 0
	interval := 3

	for x >= xBorder+interval {
		xBorder += interval
	}

	for y >= yBorder+interval {
		yBorder += interval
	}

	for i := range b[yBorder : yBorder+interval] {
		for j := range b[i][xBorder : xBorder+interval] {
			if b[i][j] == value {
				return false
			}
		}
	}

	return true
}

func (b board) checkValue(x int, y int, value int) bool {
	return b.checkYAxis(x, value) && b.checkXAxis(y, value) && b.checkSubSquare(x, y, value)
}

// func (b board) saveEmptyPositions() [][]int {
// 	emptyPositions := [][]int{}

// 	for y := range b {
// 		for x := range b[y] {
// 			if b[y][x] == 0 {
// 				emptyPositions = append(emptyPositions, []int{y, x})
// 			}
// 		}
// 	}

// 	return emptyPositions
// }

// func (b board) checkRow(row int, value int) bool {
// 	for i := range b[row] {
// 		if b[row][i] == value {
// 			return false
// 		}
// 	}

// 	return true
// }

// func (b board) checkColumn(column int, value int) bool {
// 	for i := range b {
// 		if b[i][column] == value {
// 			return false
// 		}
// 	}

// 	return true
// }

// func (b board) checkSquare(col int, row int, value int) bool {
// 	colCorner := 0
// 	rowCorner := 0
// 	squareSize := 3

// 	for col >= colCorner+squareSize {
// 		colCorner += squareSize
// 	}

// 	for row >= rowCorner+squareSize {
// 		rowCorner += squareSize
// 	}

// 	for i := range b[rowCorner : rowCorner+squareSize] {
// 		for j := range b[i][colCorner : colCorner+squareSize] {
// 			if b[i][j] == value {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

// func (b board) checkValue(col int, row int, value int) bool {
// 	return b.checkColumn(col, value) && b.checkRow(row, value) && b.checkSquare(col, row, value)
// }

// func (b board) solvePuzzle() board {
// 	emptyPositions := b.saveEmptyPositions()
// 	i := 0

// 	for i < len(emptyPositions) {
// 		limit := 9
// 		row := emptyPositions[i][0]
// 		col := emptyPositions[i][1]

// 		value := emptyPositions[row][col] + 1
// 		found := false

// 		for !found && value <= limit {
// 			if b.checkValue(col, row, value) {
// 				b[row][col] = value
// 				found = true
// 				i++
// 			} else {
// 				value++
// 			}
// 		}

// 		if !found {
// 			b[row][col] = 0
// 			i--
// 		}
// 	}

// 	return b
// }
