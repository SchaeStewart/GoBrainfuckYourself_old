package memory

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
	c := make([]Cell, size)
	m := Memory{
		Cells:       c,
		CurrentCell: len(c) / 2, // Defaults to middle so there is left space
	}
	return &m
}

// GetCurrentCell returns a pointer to the current cell
func (m *Memory) GetCurrentCell() *Cell {
	return &m.Cells[m.CurrentCell]
}

// Right moves the pointer of the current cell to the right, and returns the value of the current cell
func (m *Memory) Right() int {
	m.CurrentCell++
	return m.GetCurrentCell().Value
}

// Left moves the pointer of the current cell to the left, and returns the value of the current cell
func (m *Memory) Left() int {
	m.CurrentCell--
	return m.GetCurrentCell().Value
}
