package memory

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

func TestMemoryCustomSize(t *testing.T) {
	m := NewMemory(100)
	if len(m.Cells) != 100 {
		t.Errorf("Memory should accept a custom size for cells")
	}
}

func TestMemoryCurrentCellIsMiddle(t *testing.T) {
	m := NewMemory(100)
	if m.CurrentCell != len(m.Cells)/2 {
		t.Errorf("CurrentCell should default to the middlest cell")
	}
}

func TestMemoryRight(t *testing.T) {
	m := NewMemory()
	previousCell := m.GetCurrentCell()
	m.Right()
	if m.GetCurrentCell() == previousCell {
		t.Errorf("Current cell did not change")
	}
}

func TestMemoryLeft(t *testing.T) {
	m := NewMemory()

	previousCell := m.GetCurrentCell()
	m.Left()
	if m.GetCurrentCell() == previousCell {
		t.Errorf("Current cell did not change")
	}
}

func TestMemoryInc(t *testing.T) {
	m := NewMemory()

	m.Inc()
	if m.GetCurrentCell().Value != 1 {
		t.Errorf("Expected current cell to be 1 instead got %d", m.GetCurrentCell().Value)
	}
}

func TestMemoryDec(t *testing.T) {
	m := NewMemory()

	m.Dec()
	if m.GetCurrentCell().Value != -1 {
		t.Errorf("Expected current cell to be -1 instead got %d", m.GetCurrentCell().Value)
	}
}
