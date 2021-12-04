package main

import "errors"

type Field struct {
	number  int
	checked bool
	X       int
	Y       int
}

type Board struct {
	fields   [25]Field
	byNumber map[int]*Field
}

// The board represenation is a 5x5 matrix. The numbers
// are stored in a 1D array and contain x and y coordinates for
// their position in the matrix.
//
// For easier field access there is a map with the number as key
func NewBoard(numbers []int) (*Board, error) {
	if len(numbers) != 25 {
		return nil, errors.New("invalid number of fields. must be 25")
	}

	board := Board{
		fields:   [25]Field{},
		byNumber: make(map[int]*Field),
	}

	for index, number := range numbers {
		board.fields[index] = Field{number, false, index % 5, index / 5}
	}

	for index, field := range board.fields {
		board.byNumber[field.number] = &board.fields[index]
	}

	return &board, nil
}

// Allows to filter the board by checked or unchecked fields
func (board *Board) Filter(checked bool) []*Field {
	var fields []*Field
	for index, field := range board.fields {
		if field.checked == checked {
			fields = append(fields, &board.fields[index])
		}
	}
	return fields
}

// Tests if the board has a completed row or column
func (board *Board) HasWon() bool {
	checkedFields := board.Filter(true)

	// The xMap and yMap dictonalies are used to count
	// the number of checked fields in a row or column
	xMap := map[int]int{
		0: 0, 1: 0, 2: 0, 3: 0, 4: 0,
	}
	yMap := map[int]int{
		0: 0, 1: 0, 2: 0, 3: 0, 4: 0,
	}

	for _, field := range checkedFields {
		xMap[field.X]++
		yMap[field.Y]++
	}

	// Check if there is a completed row or column
	for _, count := range xMap {
		if count == 5 {
			return true
		}
	}

	for _, count := range yMap {
		if count == 5 {
			return true
		}
	}

	return false
}

func (board *Board) GetScore(winningNumber int) int {
	uncheckedFields := board.Filter(false)
	sum := 0

	for _, field := range uncheckedFields {
		sum += field.number
	}

	return winningNumber * sum
}
