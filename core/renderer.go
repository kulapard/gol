package core

import (
	"fmt"
)

// Renderer is an interface for rendering the game
type Renderer interface {
	Render(*GameOfLife)
}

// StdoutRenderer is a Renderer that renders the game to stdout
type StdoutRenderer struct {
}

// Render renders the game
func (s StdoutRenderer) Render(g *GameOfLife) {
	// Clear the screen
	clearScreen()

	// TODO: call the Print method on GameOfLife
	fmt.Print(g.board.String())
	fmt.Printf("Generation: %d\n", g.generation)
	fmt.Printf("Population: %d (%d%%) \n", g.board.CountAliveCells(), g.board.AlivePercentage())
	fmt.Printf("Size: %d x %d \n", g.board.Rows, g.board.Cols)
	fmt.Printf("Speed: %d generations per second\n", g.Speed)
	fmt.Println("Loader:", g.loader)

	if g.IsExtinct() {
		fmt.Println("Extinct!")
	} else if g.IsStable() {
		fmt.Println("Stable!")
	}
	fmt.Println()
	fmt.Println("[Press Ctrl+C to exit]")
}

// clearScreen clears the screen
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
