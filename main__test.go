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
