package main

import (
	"testing"
)

func TestCellIncrement(t *testing.T) {
	cell := NewCell()
	cell.Increment()
	if cell.Value != 1 {
		t.Errorf("Expected cell to equal 1 instead got %d", cell.Value)
	}
}

func TestCellDecrement(t *testing.T) {
	cell := NewCell()
	cell.Decrement()
	if cell.Value != -1 {
		t.Errorf("Expected cell to equal -1 instead got %d", cell.Value)
	}
}

func TestCellsRight(t *testing.T) {
	m := NewMemory()
	previousCell := m.GetCurrentCell()
	m.Right()
	if m.GetCurrentCell() == previousCell {
		t.Errorf("Current cell did not change")
	}
}

func TestCellsLeft(t *testing.T) {
	m := NewMemory()
	// Setup
	m.Right()

	previousCell := m.GetCurrentCell()
	m.Left()
	if m.GetCurrentCell() == previousCell {
		t.Errorf("Current cell did not change")
	}
}
