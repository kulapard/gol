// Description: Test for board.go
package core

import (
	"fmt"
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard(3, 3)
	if b.Rows != 3 {
		t.Errorf("Expected 3, got %d", b.Rows)
	}
	if b.Cols != 3 {
		t.Errorf("Expected 3, got %d", b.Cols)
	}
	// Check if the cells are initialized
	for row := range b.data {
		for col := range b.data[row] {
			// Access the cell directly
			cell := &b.data[row][col]
			if cell.IsAlive() {
				t.Error("Expected the cell to be dead")
			}
		}
	}
}
func TestBoard_Copy(t *testing.T) {
	b := NewBoard(3, 3)
	b.Randomize()
	c := b.Copy()
	if b.Rows != c.Rows {
		t.Errorf("Expected %d, got %d", b.Rows, c.Rows)
	}
	if b.Cols != c.Cols {
		t.Errorf("Expected %d, got %d", b.Cols, c.Cols)
	}
	for i, row := range b.data {
		for j, cell := range row {
			if cell != c.data[i][j] {
				t.Errorf("Expected %v, got %v", cell, c.data[i][j])
			}
		}
	}
}

func TestBoard_Randomize(t *testing.T) {
	b := NewBoard(3, 3)
	b.Randomize()
	for _, row := range b.data {
		for _, cell := range row {
			if cell.IsAlive() {
				return
			}
		}
	}
	t.Error("Expected at least one cell to be alive")
}

func TestBoard_GetNeighbours(t *testing.T) {
	b := NewBoard(3, 3)
	neighbors := b.GetNeighbours(1, 1)
	if len(neighbors) != 8 {
		t.Errorf("Expected 8 neighbors, got %d", len(neighbors))
	}
	// Check if the neighbors are correct
	if neighbors[0] != &b.data[0][0] {
		t.Errorf("Expected %v, got %v", &b.data[0][0], neighbors[0])
	}
	if neighbors[1] != &b.data[0][1] {
		t.Errorf("Expected %v, got %v", &b.data[0][1], neighbors[1])
	}
	if neighbors[2] != &b.data[0][2] {
		t.Errorf("Expected %v, got %v", &b.data[0][2], neighbors[2])
	}
	if neighbors[3] != &b.data[1][0] {
		t.Errorf("Expected %v, got %v", &b.data[1][0], neighbors[3])
	}
	if neighbors[4] != &b.data[1][2] {
		t.Errorf("Expected %v, got %v", &b.data[1][2], neighbors[4])
	}
	if neighbors[5] != &b.data[2][0] {
		t.Errorf("Expected %v, got %v", &b.data[2][0], neighbors[5])
	}
	if neighbors[6] != &b.data[2][1] {
		t.Errorf("Expected %v, got %v", &b.data[2][1], neighbors[6])
	}
	if neighbors[7] != &b.data[2][2] {
		t.Errorf("Expected %v, got %v", &b.data[2][2], neighbors[7])
	}
}
func TestBoard_CountAliveCells(t *testing.T) {
	b := NewBoard(5, 5)

	// Set the initial states
	b.data[0][0].Revive()
	b.data[1][1].Revive()
	b.data[2][2].Revive()

	count := b.CountAliveCells()
	if count != 3 {
		t.Errorf("Expected 3 alive cells, got %d", count)
	}

	b.NextGeneration()
	count = b.CountAliveCells()
	if count != 1 {
		t.Errorf("Expected 1 alive cells, got %d", count)
	}
}

func TestBoard_CountAliveNeighbours(t *testing.T) {
	b := NewBoard(5, 5)

	// Set the initial states
	/*
		1 0 0 0 0
		0 0 1 0 0
		0 0 * 1 0
		1 0 1 0 1
		0 0 0 0 1
	*/
	b.data[0][0].Revive()
	b.data[1][2].Revive()
	b.data[2][3].Revive()
	b.data[3][0].Revive()
	b.data[3][2].Revive()
	b.data[3][4].Revive()
	b.data[4][4].Revive()

	count := b.CountAliveNeighbours(2, 2)
	if count != 3 {
		t.Errorf("Expected 3 alive neighbors, got %d", count)
	}

}

func TestBoard_NextGeneration(t *testing.T) {
	b := NewBoard(3, 3)

	// Set the initial states
	/*
		1 1 1
		0 0 0
		0 0 0
	*/
	b.data[0][0].Revive()
	b.data[0][1].Revive()
	b.data[0][2].Revive()

	fmt.Println("Initial state:")
	fmt.Println(b)

	b.NextGeneration()
	// Expected state
	/*
		0 1 0
		0 1 0
		0 0 0
	*/
	fmt.Println("Next generation:")
	fmt.Println(b)

	// 0 1 0
	if b.data[0][0].IsAlive() {
		t.Error("Expected cell to be dead")
	}
	if b.data[0][1].IsDead() {
		t.Error("Expected cell to be alive")
	}
	if b.data[0][2].IsAlive() {
		t.Error("Expected cell to be dead")
	}

	// 0 1 0
	if b.data[1][0].IsAlive() {
		t.Error("Expected cell to be dead")
	}
	if b.data[1][1].IsDead() {
		t.Error("Expected cell to be alive")
	}
	if b.data[1][2].IsAlive() {
		t.Error("Expected cell to be dead")
	}

	// 0 0 0
	if b.data[2][0].IsAlive() {
		t.Error("Expected cell to be dead")
	}
	if b.data[2][1].IsAlive() {
		t.Error("Expected cell to be dead")
	}
	if b.data[2][2].IsAlive() {
		t.Error("Expected cell to be dead")
	}
}

func TestBoard_IsOutside(t *testing.T) {
	b := NewBoard(3, 3)
	if !b.IsOutside(-1, 0) {
		t.Error("Expected true, got false")
	}
	if !b.IsOutside(0, -1) {
		t.Error("Expected true, got false")
	}
	if !b.IsOutside(3, 0) {
		t.Error("Expected true, got false")
	}
	if !b.IsOutside(0, 3) {
		t.Error("Expected true, got false")
	}
	if b.IsOutside(0, 0) {
		t.Error("Expected false, got true")
	}
	if b.IsOutside(1, 1) {
		t.Error("Expected false, got true")
	}
}

func TestBoard_LoadData(t *testing.T) {
	b := NewBoard(3, 3)
	b.Randomize()

	b1 := NewBoard(3, 3)
	b1.Randomize()

	b.LoadData(b1.data)
	for i, row := range b.data {
		for j, cell := range row {
			if cell != b1.data[i][j] {
				t.Errorf("Expected %v, got %v", cell, b1.data[i][j])
			}
		}
	}
}

func TestBoard_IsExtinct(t *testing.T) {
	b := NewBoard(3, 3)
	if !b.IsExtinct() {
		t.Error("Expected true, got false")
	}
	b.aliveCount = 1
	b.generation = 2
	if b.IsExtinct() {
		t.Error("Expected false, got true")
	}
}

func TestBoard_IsStable(t *testing.T) {
	var b = NewBoard(3, 3)
	if b.IsStable() {
		t.Error("Expected false, got true")
	}
	b.NextGeneration()
	if !b.IsStable() {
		t.Error("Expected true, got false")
	}

	b.Randomize()
	b.NextGeneration()
	if b.IsStable() {
		t.Error("Expected false, got true")
	}

}
