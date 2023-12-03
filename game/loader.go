package game

import (
	"bufio"
	"fmt"
	"os"
)

type Loader interface {
	Load() (*Board, error)
	String() string
}

type FromFileLoader struct {
	FileName string
}

type RandomLoader struct {
	Rows, Cols int
}

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
	board.LoadData(data)
	return board, nil
}

func (l FromFileLoader) String() string {
	return fmt.Sprintf("file `%s`", l.FileName)
}

func (l RandomLoader) Load() (*Board, error) {
	// Validate the board size
	if l.Rows < 1 || l.Cols < 1 {
		fmt.Println("Invalid board size: rows and cols must be greater than 0")
		return nil, nil
	}
	board := NewBoard(l.Rows, l.Cols)
	board.Randomize()
	return board, nil
}

func (l RandomLoader) String() string {
	return fmt.Sprintf("random %d x %d", l.Rows, l.Cols)
}
