package evaluator

import (
	"fmt"

	"../common"
	"../memory"
	"../queue"
)

func helper(qi []queue.Item) int {
	currentDepth := qi[0].Depth
	for i, item := range qi {
		if item.Command == common.LoopEnd && item.Depth+1 == currentDepth { // TODO: fix the LoopEnd depth
			return i
		}
	}
	return 0
}

// Evaluate takes a command queue and a memory and runs the commands in the cq against the mem
func Evaluate(qi []queue.Item, mem *memory.Memory) {
	for i := 0; i < len(qi); i++ {
		item := qi[i]
		if item.Command == common.LoopStart {
			endOfLoopIdx := i + helper(qi[i:])
			Evaluate(qi[i+1:endOfLoopIdx+1], mem)
			i = endOfLoopIdx
		} else if item.Command == common.LoopEnd {
			if mem.GetCurrentCell().Value != 0 {
				Evaluate(qi, mem)
			}
			continue
		} else if item.Command == common.Inc {
			mem.Inc()
		} else if item.Command == common.Dec {
			mem.Dec()
		} else if item.Command == common.Left {
			mem.Left()
		} else if item.Command == common.Right {
			mem.Right()
		} else if item.Command == common.Get { // TODO: mem.Get()
		} else if item.Command == common.Put { // TODO: mem.Put()
			fmt.Print(string(mem.GetCurrentCell().Value))
		}
	}
}
