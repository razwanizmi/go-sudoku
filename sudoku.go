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
