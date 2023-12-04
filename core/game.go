package core

import (
	"time"
)

// GameOfLife is the core itself
type GameOfLife struct {
	Speed int // Speed in "generations per second"

	board    *Board
	loader   Loader
	renderer Renderer
}

// SetupGameOfLife creates a new game with predefined loader and renderer
func SetupGameOfLife(fileName string, speed, rows, cols int) (*GameOfLife, error) {
	var loader Loader

	// Validate speed
	if speed <= 0 {
		speed = 1
	}

	// Pick a loader
	if fileName != "" {
		loader = FromFileLoader{FileName: fileName}
	} else {
		loader = RandomLoader{Rows: rows, Cols: cols}
	}

	// Pick a renderer
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
		Speed:    speed,
		loader:   loader,
		renderer: renderer,
	}
}

// Load loads the core using predefined loader
func (g *GameOfLife) Load() error {
	// Load the board
	board, err := g.loader.Load()
	if err != nil {
		return err
	}
	g.board = board
	return nil
}

// Render renders the core using predefined renderer
func (g *GameOfLife) Render() {
	// Render the board
	g.renderer.Render(g)
}

// RunForever runs the core forever (until the board becomes stable or extinct or the user presses Ctrl+C)
func (g *GameOfLife) RunForever() {
	// Calculate sleep time
	sleep := time.Millisecond * time.Duration(1000/g.Speed)

	for {
		// Render the board
		g.Render()

		// Next generation
		g.board.NextGeneration()

		// Check if the board is stable or extinct
		if g.board.IsStable() || g.board.IsExtinct() {
			break
		}

		// Sleep
		time.Sleep(sleep)
	}
}
