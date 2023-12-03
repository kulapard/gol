package game

import (
	"fmt"
)

type Renderer interface {
	Render(*GameOfLife)
}

type StdoutRenderer struct {
}

func (s StdoutRenderer) Render(g *GameOfLife) {
	// Clear the screen
	clearScreen()

	// TODO: call the Print method on GameOfLife
	fmt.Print(g.board.String())
	fmt.Printf("Generation: %d\n", g.board.generation)
	fmt.Printf("Population: %d (%d%%) \n", g.board.CountAliveCells(), g.board.AlivePercentage())
	fmt.Printf("Size: %d x %d \n", g.board.Rows, g.board.Cols)
	fmt.Printf("Speed: %d generations per second\n", g.Speed)
	fmt.Println("Loader:", g.loader)

	if g.board.IsExtinct() {
		fmt.Println("Extinct!")
	} else if g.board.IsStable() {
		fmt.Println("Stable!")
	}
}

// clearScreen clears the screen
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
