package evaluator

import (
	"fmt"

	"../common"
	"../memory"
)

func convertByteToCommand(b byte) common.Commands {
	str := string(b)
	if item, ok := common.CommandMap[str]; ok {
		return item
	}
	return 0
}

func findNextLoopEndIdx(qi []byte) int {
	for i := 0; i < len(qi); i++ {
		item := convertByteToCommand(qi[i])
		if item == common.LoopEnd {
			return i
		}
	}
	return 0 // This will probably break something
}

func findPreviousLoopStartIdx(qi []byte) int {
	for i := 0; i < len(qi); i-- {
		item := convertByteToCommand(qi[i])
		if item == common.LoopEnd {
			return i
		}
	}
	return 0 // This will probably break something
}

// Evaluate takes a command queue and a memory and runs the commands in the cq against the mem
func Evaluate(commands []byte, mem *memory.Memory) {
	for i := 0; i < len(commands); i++ {
		item := convertByteToCommand(commands[i])
		if item == 0 {
			continue
		} else if item == common.LoopStart {
			if mem.Cells[mem.CurrentCell].Value == 0 {
				i = findNextLoopEndIdx(commands[i:]) + i + 1
			}
		} else if item == common.LoopEnd {
			if mem.Cells[mem.CurrentCell].Value != 0 {
				i = (findPreviousLoopStartIdx(commands[i:]) + i) - 1
			}
		} else if item == common.Inc {
			mem.Inc()
		} else if item == common.Dec {
			mem.Dec()
		} else if item == common.Left {
			mem.Left()
		} else if item == common.Right {
			mem.Right()
		} else if item == common.Get { // TODO: mem.Get()
			fmt.Println("YOOOOO get isn't implemented")
		} else if item == common.Put { // TODO: mem.Put()
			fmt.Print(string(mem.GetCurrentCell().Value))
		}
	}
}
