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

func RunGameOfLife(FileName string, Speed, Rows, Cols int) {
	var loader Loader
	if FileName != "" {
		loader = FromFileLoader{FileName: FileName}
	} else {
		loader = RandomLoader{Rows: Rows, Cols: Cols}
	}
	renderer := StdoutRenderer{}
	NewGameOfLife(Speed, loader, renderer).Run()
}

func NewGameOfLife(speed int, loader Loader, renderer Renderer) *GameOfLife {
	// Validate speed
	if speed <= 0 {
		speed = 1
	}
	return &GameOfLife{
		Speed:    speed,
		loader:   loader,
		renderer: renderer,
	}
}

func (g *GameOfLife) Run() {
	// Load the board
	board, err := g.loader.Load()
	if err != nil {
		fmt.Print("Error loading board:", err)
		return
	}
	g.board = board

	// Calculate sleep time
	sleep := time.Millisecond * time.Duration(1000/g.Speed)

	// Run the game
	for {
		// Render the board
		g.renderer.Render(g)

		// Next generation
		g.board.NextGeneration()

		// Sleep
		time.Sleep(sleep)

		// Check if the board is stable or extinct
		if g.board.IsStable() || g.board.IsExtinct() {
			break
		}
	}
}
