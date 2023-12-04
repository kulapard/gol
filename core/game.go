package core

import (
	"time"
)

// GameOfLife is the game itself
type GameOfLife struct {
	Speed int // Speed in "generations per second"

	board    *Board
	loader   Loader
	renderer Renderer

	isStable   bool
	generation int
}

// SetupGameOfLife creates a new game with predefined emptyLoader and stdoutRenderer
func SetupGameOfLife(fileName string, speed, rows, cols int) (*GameOfLife, error) {
	var loader Loader

	// Validate speed
	if speed <= 0 {
		speed = 1
	}

	// Pick a emptyLoader
	if fileName != "" {
		loader = FromFileLoader{FileName: fileName}
	} else {
		loader = RandomLoader{Rows: rows, Cols: cols}
	}

	// Pick a stdoutRenderer
	renderer := StdoutRenderer{}

	gol := NewGameOfLife(speed, loader, renderer)
	err := gol.Load()
	if err != nil {
		return nil, err
	}

	return gol, nil
}

// NewGameOfLife creates a new core
func NewGameOfLife(speed int, loader Loader, renderer Renderer) *GameOfLife {
	return &GameOfLife{
		Speed:      speed,
		loader:     loader,
		renderer:   renderer,
		generation: 1,
	}
}

// Load loads the game using predefined emptyLoader
func (g *GameOfLife) Load() error {
	// Load the board
	board, err := g.loader.Load()
	if err != nil {
		return err
	}
	g.board = board
	return nil
}

// Render renders the game using predefined stdoutRenderer
func (g *GameOfLife) Render() {
	// Render the board
	g.renderer.Render(g)
}

// NextGeneration calculates the next generation
func (g *GameOfLife) NextGeneration() {
	// Create a new board to store the next generation
	nextBoard := NewBoard(g.board.Rows, g.board.Cols)

	var hasChanged bool

	// Iterate through the board cells
	for row := range nextBoard.data {
		for col := range nextBoard.data[row] {
			// Get the number of alive neighbors from the current board
			// because the next generation is not calculated yet.
			aliveNeighbours := g.board.CountAliveNeighbours(row, col)

			// Access the cells directly
			oldCell := &g.board.data[row][col]
			newCell := &nextBoard.data[row][col]

			// Apply the rules of the game
			if oldCell.IsAlive() {
				if aliveNeighbours < 2 || aliveNeighbours > 3 {
					// Kill the cell
					newCell.Kill()
					hasChanged = true
				} else {
					// Keep the cell alive
					newCell.Revive()
				}
			} else {
				if aliveNeighbours == 3 {
					// Revive the cell
					newCell.Revive()
					hasChanged = true
				} else {
					// Keep the cell dead
					newCell.Kill()
				}

			}
		}
	}
	// Check if the board is stable
	g.isStable = !hasChanged

	// Copy the next generation to the current board
	g.board = nextBoard

	// Increment the generation counter
	g.generation++

}

// IsExtinct returns true if all cells are dead, false otherwise.
func (g *GameOfLife) IsExtinct() bool {
	return g.board.CountAliveCells() == 0
}

// IsStable returns true if the board is stable, false otherwise.
func (g *GameOfLife) IsStable() bool {
	return g.isStable
}

// RunForever runs the game forever (until the board becomes stable or extinct or the user presses Ctrl+C)
func (g *GameOfLife) RunForever() {
	// Calculate sleep time
	sleep := time.Millisecond * time.Duration(1000/g.Speed)

	for {
		// Render the board
		g.Render()

		// Next generation
		g.NextGeneration()

		// Check if the board is stable or extinct
		if g.IsStable() || g.IsExtinct() {
			break
		}

		// Sleep
		time.Sleep(sleep)
	}
}
