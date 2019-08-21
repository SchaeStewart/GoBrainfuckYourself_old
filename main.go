package main

import "fmt"

type Cell struct {
	Value int
}

func NewCell() *Cell {
	return &Cell{}
}

func (c *Cell) Increment() {
	c.Value++
}

func (c *Cell) Decrement() {
	c.Value--
}

type Memory struct {
	Cells       []Cell
	CurrentCell int
}

// TODO: should be able to pass in default length
func NewMemory(size ...int) *Memory {
	return &Memory{
		Cells: make([]Cell, 300000),
	}
}

func (m *Memory) GetCurrentCell() *Cell {
	return &m.Cells[m.CurrentCell]
}

func (m *Memory) Right() int {
	m.CurrentCell++
	return m.Cells[m.CurrentCell].Value
}

func (m *Memory) Left() int {
	m.CurrentCell--
	return m.Cells[m.CurrentCell].Value
}

func main() {
	fmt.Println("hello")
}
