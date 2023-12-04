package core

import (
	"errors"
	"fmt"
	"testing"
)

// BrokenLoader is a Loader that returns an error on Load
type BrokenLoader struct{}

func (l BrokenLoader) String() string {
	return "broken emptyLoader"
}
func (l BrokenLoader) Load() (*Board, error) {
	return nil, errors.New("broken emptyLoader")
}

var emptyLoader = EmptyLoader{Rows: 3, Cols: 3}
var brokenLoader = BrokenLoader{}
var stdoutRenderer = StdoutRenderer{}

func TestNewGameOfLife(t *testing.T) {
	g := NewGameOfLife(1, emptyLoader, stdoutRenderer)
	if g.Speed != 1 {
		t.Fatalf("NewGameOfLife should set Speed to 1, got %d", g.Speed)
	}
}

func TestSetupGameOfLife(t *testing.T) {
	t.Run("ok random", func(t *testing.T) {
		g, err := SetupGameOfLife("", 1, 3, 3)
		if err != nil {
			t.Fatal("SetupGameOfLife should not return nil")
		}
		if g.Speed != 1 {
			t.Fatalf("SetupGameOfLife should set Speed to 1, got %d", g.Speed)
		}
	})

	t.Run("ok from file", func(t *testing.T) {
		_, err := SetupGameOfLife("testdata/3x3.txt", 1, 3, 3)
		if err != nil {
			t.Fatal("SetupGameOfLife should not return nil")
		}
	})
	t.Run("ok fixing speed", func(t *testing.T) {
		g, err := SetupGameOfLife("", -1, 3, 3)
		if g == nil && err != nil {
			t.Fatal("SetupGameOfLife should fix speed and not return error:", err)
		}
		if g.Speed != 1 {
			t.Fatalf("SetupGameOfLife should set Speed to 1, got %d", g.Speed)
		}
	})

	t.Run("invalid Rows", func(t *testing.T) {
		g, err := SetupGameOfLife("", 1, 0, 3)
		if g != nil && err == nil {
			t.Fatal("SetupGameOfLife should return nil")
		}
	})
	t.Run("invalid board size", func(t *testing.T) {
		g, err := SetupGameOfLife("", 1, 0, 3)
		if g != nil && err == nil && err.Error() != "invalid board size" {
			t.Fatal("SetupGameOfLife should return nil")
		}
	})
	t.Run("invalid fileName", func(t *testing.T) {
		g, err := SetupGameOfLife("invalid", 1, 0, 3)
		if g != nil && err == nil {
			t.Fatal("SetupGameOfLife should return nil")
		}
	})

}

func TestGameOfLife_Load(t *testing.T) {
	g := NewGameOfLife(1, emptyLoader, stdoutRenderer)
	if g.board != nil {
		t.Fatal("Board should be nil before Load")
	}
	if err := g.Load(); err != nil {
		t.Fatal("Load should not return error")
	}
	if g.board == nil {
		t.Fatal("Load should set board")
	}

	g = NewGameOfLife(1, brokenLoader, stdoutRenderer)
	err := g.Load()
	if err == nil {
		t.Fatal("Load should return error")
	}
	if err.Error() != "broken emptyLoader" {
		t.Fatalf("Load should return 'broken emptyLoader', got %v", err)
	}
}

func TestGameOfLife_Render(_ *testing.T) {
	g := NewGameOfLife(1, emptyLoader, stdoutRenderer)
	g.Load()
	g.Render()
}

func TestGameOfLife_NextGeneration(t *testing.T) {
	g := NewGameOfLife(1, emptyLoader, stdoutRenderer)
	g.Load()

	// Set the initial states
	/*
		1 1 1
		0 0 0
		0 0 0
	*/
	g.board.data[0][0].Revive()
	g.board.data[0][1].Revive()
	g.board.data[0][2].Revive()

	fmt.Println("Initial state:")
	fmt.Println(g.board)

	g.NextGeneration()
	// Expected state
	/*
		0 1 0
		0 1 0
		0 0 0
	*/
	fmt.Println("Next generation:")
	fmt.Println(g.board)

	// 0 1 0
	if g.board.data[0][0].IsAlive() {
		t.Error("Expected cell to be dead")
	}
	if g.board.data[0][1].IsDead() {
		t.Error("Expected cell to be alive")
	}
	if g.board.data[0][2].IsAlive() {
		t.Error("Expected cell to be dead")
	}

	// 0 1 0
	if g.board.data[1][0].IsAlive() {
		t.Error("Expected cell to be dead")
	}
	if g.board.data[1][1].IsDead() {
		t.Error("Expected cell to be alive")
	}
	if g.board.data[1][2].IsAlive() {
		t.Error("Expected cell to be dead")
	}

	// 0 0 0
	if g.board.data[2][0].IsAlive() {
		t.Error("Expected cell to be dead")
	}
	if g.board.data[2][1].IsAlive() {
		t.Error("Expected cell to be dead")
	}
	if g.board.data[2][2].IsAlive() {
		t.Error("Expected cell to be dead")
	}
}

func TestGameOfLife_IsStable(t *testing.T) {
	g := NewGameOfLife(1, RandomLoader{Rows: 3, Cols: 3}, stdoutRenderer)
	g.Load()

	if g.HasChanged() {
		t.Error("Expected false, got true")
	}
	g.NextGeneration()
	if !g.HasChanged() {
		t.Error("Expected true, got false")
	}

	g = NewGameOfLife(1, emptyLoader, stdoutRenderer)
	g.Load()
	g.NextGeneration()
	if g.HasChanged() {
		t.Error("Expected false, got true")
	}

}

func TestBoard_IsExtinct(t *testing.T) {
	g := NewGameOfLife(1, emptyLoader, stdoutRenderer)
	g.Load()

	if !g.IsExtinct() {
		t.Error("Expected true, got false")
	}
	g = NewGameOfLife(1, RandomLoader{Rows: 3, Cols: 3}, stdoutRenderer)
	g.Load()
	if g.IsExtinct() {
		t.Error("Expected false, got true")
	}
}
