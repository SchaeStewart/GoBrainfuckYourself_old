package queue

import (
	"testing"

	"../common"
)

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

func TestAddCommand(t *testing.T) {
	// Should add a commands Enum to the queue
	cq := NewCommandQueue()
	cq.AddCommand(common.Right)

	if cq.commands[0].command != common.Right {
		t.Error("Could not add comamnd to queue")
	}
}

func TestAddCommandResizeSlice(t *testing.T) {
	// Should automatically resize commands slice when inital size is execeeded
	cq := NewCommandQueue(1)
	cq.AddCommand(common.Right)
	cq.AddCommand(common.Left)

	if cq.commands[1].command != common.Left && len(cq.commands) != 2 {
		t.Error("Commands slice did not automatically resize")
	}
}

func TestCalcuateDepth(t *testing.T) {
	cq := NewCommandQueue()
	depthBefore := cq.currentDepth
	cq.AddCommand(common.LoopStart)
	depthAfter := cq.currentDepth

	if depthBefore != 0 || depthAfter != 1 {
		t.Errorf("Expected depth to be %d, instead got %d", 1, depthAfter)
	}
}

// test that added commands have a proper depth assinged

func TestCalcuateDepthError(t *testing.T) {
	cq := NewCommandQueue()
	cq.AddCommand(common.LoopEnd)
	// Should panic/throw error
}
