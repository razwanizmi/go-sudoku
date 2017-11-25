package main

import (
	"reflect"
	"testing"
)

var boardString = "090000006000960485000581000004000000517200900602000370100804020706000810300090000"
var parsedBoard board

func TestNewBoard(t *testing.T) {
	parsedBoard = newBoard(boardString)
	expectedBoard := board{
		{0, 9, 0, 0, 0, 0, 0, 0, 6},
		{0, 0, 0, 9, 6, 0, 4, 8, 5},
		{0, 0, 0, 5, 8, 1, 0, 0, 0},
		{0, 0, 4, 0, 0, 0, 0, 0, 0},
		{5, 1, 7, 2, 0, 0, 9, 0, 0},
		{6, 0, 2, 0, 0, 0, 3, 7, 0},
		{1, 0, 0, 8, 0, 4, 0, 2, 0},
		{7, 0, 6, 0, 0, 0, 8, 1, 0},
		{3, 0, 0, 0, 9, 0, 0, 0, 0},
	}

	if len(parsedBoard) != 9 {
		t.Errorf("Expected colum length to be %v but got %v", 9, len(parsedBoard))
	}

	if len(parsedBoard[0]) != 9 {
		t.Errorf("Expected row lenght to be %v but got %v", 9, len(parsedBoard[0]))
	}

	if !reflect.DeepEqual(parsedBoard, expectedBoard) {
		t.Errorf("Expected board to be %v but got %v", expectedBoard, parsedBoard)
	}
}

func TestSaveEmptyPositions(t *testing.T) {
	positions := parsedBoard.saveEmptyPositions()
	expectedPositions := [][]int{
		{0, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {0, 1}, {1, 1},
		{2, 1}, {5, 1}, {0, 2}, {1, 2}, {2, 2}, {6, 2}, {7, 2}, {8, 2}, {0, 3},
		{1, 3}, {3, 3}, {4, 3}, {5, 3}, {6, 3}, {7, 3}, {8, 3}, {4, 4}, {5, 4},
		{7, 4}, {8, 4}, {1, 5}, {3, 5}, {4, 5}, {5, 5}, {8, 5}, {1, 6}, {2, 6},
		{4, 6}, {6, 6}, {8, 6}, {1, 7}, {3, 7}, {4, 7}, {5, 7}, {8, 7}, {1, 8},
		{2, 8}, {3, 8}, {5, 8}, {6, 8}, {7, 8}, {8, 8},
	}

	if len(positions) != 51 {
		t.Errorf("Expected number of empty positions to be %v but got %v", 51, len(positions))
	}

	if !reflect.DeepEqual(positions, expectedPositions) {
		t.Errorf("Expected empty positions to be %v but got %v", expectedPositions, positions)
	}
}

func TestCheckXAxis(t *testing.T) {
	if parsedBoard.checkXAxis(0, 2) != true {
		t.Errorf("Expected value 2 to be valid (true) at x-axis 0, but got false")
	}

	if parsedBoard.checkXAxis(0, 9) != false {
		t.Errorf("Expected value 9 to be invalid (false) at x-axis 0, but got true")
	}
}

func TestCheckYAxis(t *testing.T) {
	if parsedBoard.checkYAxis(0, 9) != true {
		t.Errorf("Expected value 9 to be valid (true) at x-axis 0, but got false")
	}

	if parsedBoard.checkYAxis(0, 5) != false {
		t.Errorf("Expected value 5 to be invalid (false) at x-axis 0, but got true")
	}
}

func TestCheckSubSquare(t *testing.T) {
	if parsedBoard.checkSubSquare(2, 2, 1) != true {
		t.Errorf("Expected value 1 to be valid (true) at sub-square (2, 2) but got false")
	}

	if parsedBoard.checkSubSquare(2, 2, 9) != false {
		t.Errorf("Expected value 9 to be invalid (false) at sub-square (2, 2) but got true")
	}
}

func TestCheckValue(t *testing.T) {
	if parsedBoard.checkValue(0, 0, 2) != true {
		t.Errorf("Expected value 2 to be valid (true) at position (0, 0), but got false")
	}

	if parsedBoard.checkValue(0, 0, 9) != false {
		t.Errorf("Expected value 9 to be invalid (false) at position (0, 0), but got true")
	}
}

func TestSolvePuzzle(t *testing.T) {
	solution := parsedBoard.solvePuzzle()
	expectedSolution := board{
		{8, 9, 5, 7, 4, 2, 1, 3, 6},
		{2, 7, 1, 9, 6, 3, 4, 8, 5},
		{4, 6, 3, 5, 8, 1, 7, 9, 2},
		{9, 3, 4, 6, 1, 7, 2, 5, 8},
		{5, 1, 7, 2, 3, 8, 9, 6, 4},
		{6, 8, 2, 4, 5, 9, 3, 7, 1},
		{1, 5, 9, 8, 7, 4, 6, 2, 3},
		{7, 4, 6, 3, 2, 5, 8, 1, 9},
		{3, 2, 8, 1, 9, 6, 5, 4, 7},
	}

	if !reflect.DeepEqual(solution, expectedSolution) {
		t.Errorf("Expected solution to be %v but got %v", expectedSolution, solution)
	}
}
