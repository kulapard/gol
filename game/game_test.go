package game

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
	if g == nil {
		t.Error("NewGameOfLife should not return nil")
	}
	if g.Speed != 1 {
		t.Errorf("NewGameOfLife should set Speed to 1, got %d", g.Speed)
	}

	g = NewGameOfLife(-1, loader, renderer)
	if g.Speed != 1 {
		t.Errorf("NewGameOfLife should set Speed to 1, got %d", g.Speed)
	}

}

func TestGameOfLife_Load(t *testing.T) {
	g := NewGameOfLife(1, loader, renderer)
	if g.board != nil {
		t.Error("Board should be nil before Load")
	}
	if err := g.Load(); err != nil {
		t.Error("Load should not return error")
	}
	if g.board == nil {
		t.Error("Load should set board")
	}

	g = NewGameOfLife(1, brokenLoader, renderer)
	err := g.Load()
	if err == nil {
		t.Error("Load should return error")
	}
	if err.Error() != "broken loader" {
		t.Errorf("Load should return 'broken loader', got %v", err)
	}
}
