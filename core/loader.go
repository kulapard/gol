package core

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Loader is an interface for loading a board
type Loader interface {
	Load() (*Board, error)
	String() string
}

// FromFileLoader is a Loader that loads a board from a file
type FromFileLoader struct {
	FileName string
}

// RandomLoader is a Loader that returns a random board
type RandomLoader struct {
	Rows, Cols int
}

// EmptyLoader is a Loader that returns an empty board
type EmptyLoader struct {
	Rows, Cols int
}

func (l FromFileLoader) String() string {
	return fmt.Sprintf("file `%s`", l.FileName)
}

func (l RandomLoader) String() string {
	return fmt.Sprintf("random %d x %d", l.Rows, l.Cols)
}

func (l EmptyLoader) String() string {
	return fmt.Sprintf("empty %d x %d", l.Rows, l.Cols)
}

// Load returns a board loaded from a file
func (l FromFileLoader) Load() (*Board, error) {
	file, err := os.Open(l.FileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]Cell
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var row []Cell
		for _, char := range line {
			// Ignore spaces
			if char == ' ' {
				continue
			}
			if char != '0' && char != '.' {
				return nil, fmt.Errorf("invalid character: '%c'", char)
			}
			cell := Cell{}
			if char == '0' {
				cell.Revive()
			} else {
				cell.Kill()
			}
			row = append(row, cell)
		}
		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	board := NewBoard(len(data), len(data[0])) // Assuming all data have equal length
	board.data = data
	return board, nil
}

// Load returns a board with random data
func (l RandomLoader) Load() (*Board, error) {
	// Validate the board size
	if l.Rows < 1 || l.Cols < 1 {
		return nil, errors.New("invalid board size")
	}
	board := NewBoard(l.Rows, l.Cols)

	// Randomize the board
	for row := range board.data {
		for col := range board.data[row] {
			// Access the cell directly
			cell := &board.data[row][col]
			if randBool() {
				cell.Kill()
			} else {
				cell.Revive()
			}

		}
	}
	return board, nil
}

// Load returns an empty board (for testing)
func (l EmptyLoader) Load() (*Board, error) {
	// Validate the board size
	if l.Rows < 1 || l.Cols < 1 {
		return nil, errors.New("invalid board size")
	}
	board := NewBoard(l.Rows, l.Cols)
	return board, nil
}
