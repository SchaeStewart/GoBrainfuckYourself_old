package main

import "fmt"

// Cell struct
type Cell struct {
	Value int
}

// NewCell return Cell
func NewCell() *Cell {
	return &Cell{}
}

// Increment value of cell
func (c *Cell) Increment() {
	c.Value++
}

// Decrement value of cell
func (c *Cell) Decrement() {
	c.Value--
}

// Memory struct
type Memory struct {
	Cells       []Cell
	CurrentCell int
}

// NewMemory structure
func NewMemory(s ...int) *Memory {
	size := 300000
	if len(s) > 0 {
		size = s[0]
	}
	return &Memory{
		Cells: make([]Cell, size),
	}
}

// GetCurrentCell returns a pointer to the current cell
func (m *Memory) GetCurrentCell() *Cell {
	return &m.Cells[m.CurrentCell]
}

// Right moves the pointer of the current cell to the right
func (m *Memory) Right() int {
	m.CurrentCell++
	return m.Cells[m.CurrentCell].Value
}

// Left moves the pointer of the current cell to the left
func (m *Memory) Left() int {
	m.CurrentCell--
	return m.Cells[m.CurrentCell].Value
}

func main() {
	fmt.Println("hello")
}
