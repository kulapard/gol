package core

import "testing"

func TestStdoutRenderer_Render(t *testing.T) {
	r := StdoutRenderer{}
	g := NewGameOfLife(1, EmptyLoader{Rows: 3, Cols: 3}, r)
	if err := g.Load(); err != nil {
		t.Error("Error loading board:", err)
	}
	r.Render(g)

	g = NewGameOfLife(1, RandomLoader{Rows: 3, Cols: 3}, r)
	if err := g.Load(); err != nil {
		t.Error("Error loading board:", err)
	}
	r.Render(g)
}
