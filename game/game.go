package game

import (
	"fmt"
	"time"
)

type GameOfLife struct {
	Speed int // Speed in "generations per second"

	board    *Board
	loader   Loader
	renderer Renderer
}

func SetupGameOfLife(FileName string, Speed, Rows, Cols int) *GameOfLife {
	var loader Loader

	// Validate speed
	if Speed <= 0 {
		Speed = 1
	}

	// Pick a loader
	if FileName != "" {
		loader = FromFileLoader{FileName: FileName}
	} else {
		loader = RandomLoader{Rows: Rows, Cols: Cols}
	}

	// Pick a renderer
	renderer := StdoutRenderer{}

	gol := NewGameOfLife(Speed, loader, renderer)
	gol.Load()
	return gol
}

func NewGameOfLife(speed int, loader Loader, renderer Renderer) *GameOfLife {
	return &GameOfLife{
		Speed:    speed,
		loader:   loader,
		renderer: renderer,
	}
}

func (g *GameOfLife) Load() error {
	// Load the board
	board, err := g.loader.Load()
	if err != nil {
		fmt.Print("Error loading board:", err)
		return err
	}
	g.board = board
	return nil
}
func (g *GameOfLife) Render() {
	// Render the board
	g.renderer.Render(g)
}

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
