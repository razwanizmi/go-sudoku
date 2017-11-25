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

func (b board) solvePuzzle() board {
	emptyPositions := b.saveEmptyPositions()
	limit := 9
	i := 0

	for i < len(emptyPositions) {
		x := emptyPositions[i][0]
		y := emptyPositions[i][1]

		value := b[y][x] + 1
		found := false

		for !found && value <= limit {
			if b.checkValue(x, y, value) {
				b[y][x] = value
				found = true
				i++
			} else {
				value++
			}
		}

		if !found {
			b[y][x] = 0
			i--
		}
	}

	return b
}
