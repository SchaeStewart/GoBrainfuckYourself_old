package evaluator

import (
	"fmt"
	"testing"

	"../memory"
	"../queue"
)

func TestExecute(t *testing.T) {
	cq := queue.NewCommandQueue()
	mem := memory.NewMemory(100)
	helloworldBF := []byte("++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<++.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-]<+.")
	cq.LoadQueue(helloworldBF)

	// TODO: actually test something here.
	Evaluate(cq.Commands, mem)
	fmt.Println("")
}

// TODO: add more tests
