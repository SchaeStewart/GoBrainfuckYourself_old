package queue

import (
	"testing"

	"../common"
)

func TestNewCommandQueue(t *testing.T) {
	// Should create a CommandQueue and default to a size opf 1500
	cq := NewCommandQueue()
	if cq == nil || len(cq.Commands) != 1500 {
		t.Error("CommandQueue should default to a size of 1500")
	}
}
func TestNewCommandQueueSize(t *testing.T) {
	// Should create a CommandQueue and have len of the customSize
	customSize := 300
	cq := NewCommandQueue(customSize)
	if cq == nil || len(cq.Commands) != customSize {
		t.Error("CommandQueue should default to a size of 1500")
	}
}

func TestAddCommand(t *testing.T) {
	// Should add a commands Enum to the queue
	cq := NewCommandQueue()
	cq.AddCommand(common.Right)

	if cq.Commands[0].Command != common.Right {
		t.Error("Could not add comamnd to queue")
	}
}

func TestAddCommandResizeSlice(t *testing.T) {
	// Should automatically resize commands slice when inital size is execeeded
	cq := NewCommandQueue(1)
	cq.AddCommand(common.Right)
	cq.AddCommand(common.Left)

	if cq.Commands[1].Command != common.Left && len(cq.Commands) != 2 {
		t.Error("Commands slice did not automatically resize")
	}
}

func TestCalcuateDepth_LoopStart(t *testing.T) {
	cq := NewCommandQueue()
	depthBefore := cq.currentDepth
	cq.AddCommand(common.LoopStart)
	depthAfter := cq.currentDepth

	if depthBefore != 0 || depthAfter != 1 {
		t.Errorf("Expected depth to be %d, instead got %d", 1, depthAfter)
	}
}

func TestCalcuateDepth_LoopEnd(t *testing.T) {
	// Setup
	cq := NewCommandQueue()
	cq.AddCommand(common.LoopStart)

	// Test
	depthBefore := cq.currentDepth
	cq.AddCommand(common.LoopEnd)
	depthAfter := cq.currentDepth

	if depthBefore != 1 || depthAfter != 0 {
		t.Errorf("Expected depth to be %d, instead got %d", 1, depthAfter)
	}

}

// test that added commands have a proper depth assinged

func TestCalcuateDepthError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	cq := NewCommandQueue()
	cq.AddCommand(common.LoopEnd)
}

func TestLoadQueue(t *testing.T) {
	helloworldBF := []byte("+[-[<<[+[--->]-[<<<]]]>>>-]>-.---.>..>.<<<<-.<+.>>>>>.>.<<.<-.")

	cq := NewCommandQueue()
	cq.LoadQueue(helloworldBF)
	if cq.Commands[0].Command != common.Inc {
		t.Errorf("Error expected first item in Commands to be Inc, instead got %d", cq.Commands[0].Command)
	}
}
