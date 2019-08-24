package queue

import "testing"

func TestNewCommandQueue(t *testing.T) {
	// Should create a CommandQueue and default to a size opf 1500
	cq := NewCommandQueue()
	if cq == nil || len(cq.commands) != 1500 {
		t.Error("CommandQueue should default to a size of 1500")
	}
}
func TestNewCommandQueueSize(t *testing.T) {
	// Should create a CommandQueue and have len of the customSize
	customSize := 300
	cq := NewCommandQueue(customSize)
	if cq == nil || len(cq.commands) != customSize {
		t.Error("CommandQueue should default to a size of 1500")
	}
}
