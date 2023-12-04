package core

import (
	"math"
	"math/rand"
)

// Board represents the core board
type Board struct {
	Rows, Cols int
	data       [][]Cell
	generation int // Generation counter
	isStable   bool
	aliveCount int
}

// randBool returns a random boolean value
func randBool() bool {
	return rand.Intn(2) == 0 // nolint
}

// NewBoard creates a new board with the given number of rows and columns.
func NewBoard(rows, cols int) *Board {
	data := make([][]Cell, rows)
	for row := range data {
		data[row] = make([]Cell, cols)
	}
	return &Board{Rows: rows, Cols: cols, data: data, generation: 1}
}

// Copy returns a copy of the board.
func (b *Board) Copy() Board {
	data := make([][]Cell, b.Rows)
	for row := range data {
		data[row] = make([]Cell, b.Cols)
		for col := range data[row] {
			// Copy the cell
			data[row][col] = b.data[row][col]
		}
	}
	return Board{Rows: b.Rows, Cols: b.Cols, data: data}
}

// Randomize Randomly populate the board with alive cells.
// TODO: move this to a RandomLoader
func (b *Board) Randomize() {
	for row := range b.data {
		for col := range b.data[row] {
			// Access the cell directly
			cell := &b.data[row][col]
			if randBool() {
				cell.Kill()
			} else {
				cell.Revive()
			}

		}
	}
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
func (b *Board) CountAliveCells() int {
	// Calculate the alive cells only for the first generation
	if b.generation == 1 {
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
	return b.aliveCount
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

// NextGeneration iterates through the board cells and apply the rules of the core.
func (b *Board) NextGeneration() {
	// Create a new board to store the next generation
	nextBoard := b.Copy()

	var hasChanged bool
	var aliveCount int

	// Iterate through the board cells
	for row := range nextBoard.data {
		for col := range nextBoard.data[row] {
			// Get the number of alive neighbors from the current board
			// because the next generation is not calculated yet.
			aliveNeighbours := b.CountAliveNeighbours(row, col)

			// Access the cell directly
			cell := &nextBoard.data[row][col]

			// Apply the rules of the core
			if cell.IsAlive() {
				if aliveNeighbours < 2 || aliveNeighbours > 3 {
					// Kill the cell
					cell.Kill()
					hasChanged = true
				} else {
					aliveCount++
				}
			} else if aliveNeighbours == 3 {
				// Revive the cell
				cell.Revive()
				hasChanged = true
				aliveCount++
			}
		}
	}
	// Check if the board is stable
	b.isStable = !hasChanged
	b.aliveCount = aliveCount

	// Copy the next generation to the current board
	b.data = nextBoard.data
	// Increment the generation counter
	b.generation++

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

// LoadData loads the given data into the board.
func (b *Board) LoadData(data [][]Cell) {
	b.data = data
}

// IsExtinct returns true if all cells are dead, false otherwise.
func (b *Board) IsExtinct() bool {
	return b.CountAliveCells() == 0
}

// IsStable returns true if the board is stable, false otherwise.
func (b *Board) IsStable() bool {
	return b.isStable
}
