package evaluator

import (
	"fmt"

	"../common"
	"../memory"
	"../queue"
)

// Evaluate takes a command queue and a memory and runs the commands in the cq against the mem
func Evaluate(qi []queue.Item, mem *memory.Memory) {
	loopDepthMap := make(map[int]int)
	for i := 0; i < len(qi); i++ {
		item := qi[i]
		if item.Command == common.LoopStart {
			loopDepthMap[item.Depth] = i + 1
		} else if item.Command == common.LoopEnd {
			if mem.GetCurrentCell().Value != 0 {
				i = loopDepthMap[item.Depth+1] - 1
			}
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
