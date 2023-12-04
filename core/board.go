package core

import (
	"math"
)

// Board represents the game board
type Board struct {
	Rows, Cols int
	data       [][]Cell
}

// NewBoard creates a new board with the given number of rows and columns.
func NewBoard(rows, cols int) *Board {
	data := make([][]Cell, rows)
	for row := range data {
		data[row] = make([]Cell, cols)
	}
	return &Board{Rows: rows, Cols: cols, data: data}
}

// GetNeighbours returns the neighbors of the given cell.
func (b *Board) GetNeighbours(row, coll int) []*Cell {
	var neighbors []*Cell
	for r := row - 1; r <= row+1; r++ {
		for c := coll - 1; c <= coll+1; c++ {
			// Skip if the cell is out of bounds
			if b.IsOutside(r, c) {
				continue
			}
			// Skip the current cell
			if r == row && c == coll {
				continue
			}
			// Access the cell directly
			cell := &b.data[r][c]

			// Append the neighbors
			neighbors = append(neighbors, cell)
		}
	}
	return neighbors
}

// CountAliveNeighbours returns the number of alive neighbors for a given cell.
func (b *Board) CountAliveNeighbours(row, coll int) int {
	var count int
	for _, n := range b.GetNeighbours(row, coll) {
		if n.IsAlive() {
			count++
		}
	}
	return count
}

// CountAliveCells returns the number of alive cells.
// TODO: calculate this only when the board changes
func (b *Board) CountAliveCells() int {
	// Calculate the alive cells only for the first generation
	var count int
	for i := range b.data {
		for j := range b.data[i] {
			// Access the cell directly
			cell := &b.data[i][j]
			if cell.IsAlive() {
				count++
			}
		}
	}
	return count
}

// TotalCells returns the total number of cells.
func (b *Board) TotalCells() int {
	return b.Rows * b.Cols
}

// AlivePercentage returns the percentage of alive cells.
func (b *Board) AlivePercentage() int {
	percentage := float64(b.CountAliveCells()) / float64(b.TotalCells()) * 100
	return int(math.Round(percentage))
}

func (b *Board) String() string {
	var s string
	for _, row := range b.data {
		for _, cell := range row {
			// Add the cell string to the board string
			s += cell.String()
		}
		// Add a new line
		s += "\n"
	}
	return s

}

// IsOutside returns true if the given Row and column are outside the board.
func (b *Board) IsOutside(row, col int) bool {
	return row < 0 || row >= b.Rows || col < 0 || col >= b.Cols
}
