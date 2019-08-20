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

func main() {
	fmt.Println("hello")
}
