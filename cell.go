package main

type Cell struct {
	isAlive bool
}

const (
	YellowBlock string = "\033[48;5;226m  \033[0m"
	GreyBlock   string = "\033[48;5;240m  \033[0m"
)

// Draw returns the string representation of the cell
func (c *Cell) String() string {
	if c.isAlive {
		return YellowBlock
	}
	return GreyBlock
}

// IsAlive returns true if the cell is alive, false otherwise
func (c *Cell) IsAlive() bool {
	return c.isAlive
}

// IsDead returns true if the cell is dead, false otherwise
func (c *Cell) IsDead() bool {
	return !c.isAlive
}

// Kill kills the cell
func (c *Cell) Kill() {
	c.isAlive = false
}

// Revive revives the cell
func (c *Cell) Revive() {
	c.isAlive = true
}
