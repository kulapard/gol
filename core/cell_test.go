package core

import "testing"

func TestCell_String(t *testing.T) {
	c1 := Cell{isAlive: true}
	c2 := Cell{isAlive: false}
	if c1.String() == c2.String() {
		t.Errorf("Expected different strings, got %s and %s", c1.String(), c2.String())
	}
}

func TestCell_Kill_Revive(t *testing.T) {
	c := Cell{isAlive: true}
	if !c.IsAlive() {
		t.Error("Expected true, got ", c.IsAlive())
	}

	c.Kill()
	if c.IsAlive() {
		t.Error("Expected false, got ", c.IsAlive())
	}

	c.Revive()
	if !c.IsAlive() {
		t.Error("Expected true, got ", c.IsAlive())
	}
}
