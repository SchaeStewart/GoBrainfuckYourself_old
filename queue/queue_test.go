package queue

import (
	"fmt"
	"testing"

	"../common"
)

func TestNewCommandQueue(t *testing.T) {
	// Should create a CommandQueue and default to a size opf 1500
	cq := NewCommandQueue()
	if cq == nil || len(cq.Commands) != 1500 {
		t.Error("CommandQueue should default to a size of 1500")
	}
	fmt.Println(cq.Commands[0])
	fmt.Println(common.Left)
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

	if cq.Commands[0] != common.Right {
		t.Error("Could not add comamnd to queue")
	}
}

func TestAddCommandResizeSlice(t *testing.T) {
	// Should automatically resize commands slice when inital size is execeeded
	cq := NewCommandQueue(1)
	cq.AddCommand(common.Right)
	cq.AddCommand(common.Left)

	if cq.Commands[1] != common.Left && len(cq.Commands) != 2 {
		t.Error("Commands slice did not automatically resize")
	}
}

func TestLoadQueue(t *testing.T) {
	helloworldBF := []byte("+[-[<<[+[--->]-[<<<]]]>>>-]>-.---.>..>.<<<<-.<+.>>>>>.>.<<.<-.")

	cq := NewCommandQueue()
	cq.LoadQueue(helloworldBF)
	if cq.Commands[0] != common.Inc {
		t.Errorf("Error expected first item in Commands to be Inc, instead got %d", cq.Commands[0])
	}
}
