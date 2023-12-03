package game

import "testing"

func TestFromFileLoader_Load(t *testing.T) {
	l := FromFileLoader{FileName: "testdata/3x3.txt"}
	board, err := l.Load()
	if err != nil {
		t.Error("Error loading board:", err)
	}

	if board.Rows != 3 || board.Cols != 3 {
		t.Errorf("Expected 3x3 board, got %dx%d", board.Rows, board.Cols)
	}

	if board.data[0][0].IsDead() {
		t.Error("Expected alive cell, got dead")
	}
	if board.data[1][1].IsDead() {
		t.Error("Expected alive cell, got dead")
	}
	if board.data[2][2].IsDead() {
		t.Error("Expected alive cell, got dead")
	}
	if board.CountAliveCells() != 3 {
		t.Error("Expected 3 alive cells, got ", board.CountAliveCells())
	}

	l = FromFileLoader{FileName: "testdata/invalid.txt"}
	_, err = l.Load()
	if err == nil {
		t.Error("Expected error loading board, got nil")
	}

	l = FromFileLoader{FileName: "testdata/not_found.txt"}
	_, err = l.Load()
	if err == nil {
		t.Error("Expected error loading board, got nil")
	}
}

func TestFromFileLoader_String(t *testing.T) {
	l := FromFileLoader{FileName: "testdata/3x3.txt"}
	if l.String() != "file `testdata/3x3.txt`" {
		t.Error("Expected 'file `testdata/3x3.txt`', got ", l.String())
	}
}

func TestRandomLoader_Load(t *testing.T) {
	l := RandomLoader{Rows: 3, Cols: 3}
	_, err := l.Load()
	if err != nil {
		t.Error("Error loading board:", err)
	}

	l = RandomLoader{Rows: -1, Cols: -1}
	_, err = l.Load()
	if err == nil {
		t.Error("Expected error loading board, got nil")
	}
}

func TestRandomLoader_String(t *testing.T) {
	l := RandomLoader{Rows: 3, Cols: 3}
	if l.String() != "random 3 x 3" {
		t.Error("Expected 'random 3 x 3', got ", l.String())
	}
}

func TestEmptyLoader_Load(t *testing.T) {
	l := EmptyLoader{Rows: 3, Cols: 3}
	board, err := l.Load()
	if err != nil {
		t.Error("Error loading board:", err)
	}

	if board.Rows != 3 || board.Cols != 3 {
		t.Errorf("Expected 3x3 board, got %dx%d", board.Rows, board.Cols)
	}

	if board.CountAliveCells() != 0 {
		t.Error("Expected 0 alive cells, got ", board.CountAliveCells())
	}

	l = EmptyLoader{Rows: -1, Cols: -1}
	_, err = l.Load()
	if err == nil {
		t.Error("Expected error loading board, got nil")
	}
}

func TestEmptyLoader_String(t *testing.T) {
	l := EmptyLoader{Rows: 3, Cols: 3}
	if l.String() != "empty 3 x 3" {
		t.Error("Expected 'empty 3 x 3', got ", l.String())
	}
}
