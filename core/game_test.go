package core

import (
	"errors"
	"testing"
)

// BrokenLoader is a Loader that returns an error on Load
type BrokenLoader struct{}

func (l BrokenLoader) String() string {
	return "broken loader"
}
func (l BrokenLoader) Load() (*Board, error) {
	return nil, errors.New("broken loader")
}

var loader = EmptyLoader{Rows: 3, Cols: 3}
var brokenLoader = BrokenLoader{}
var renderer = StdoutRenderer{}

func TestNewGameOfLife(t *testing.T) {
	g := NewGameOfLife(1, loader, renderer)
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
	g := NewGameOfLife(1, loader, renderer)
	if g.board != nil {
		t.Fatal("Board should be nil before Load")
	}
	if err := g.Load(); err != nil {
		t.Fatal("Load should not return error")
	}
	if g.board == nil {
		t.Fatal("Load should set board")
	}

	g = NewGameOfLife(1, brokenLoader, renderer)
	err := g.Load()
	if err == nil {
		t.Fatal("Load should return error")
	}
	if err.Error() != "broken loader" {
		t.Fatalf("Load should return 'broken loader', got %v", err)
	}
}

func TestGameOfLife_Render(_ *testing.T) {
	g := NewGameOfLife(1, loader, renderer)
	g.Load()
	g.Render()
}
