package queue

import (
	"../common"
)

// CommandQueue loads commands into a queue
type CommandQueue struct {
	Commands   []common.Commands
	currentIdx int
}

// NewCommandQueue returns a CommandQueue of a given size. Defaults to 1500
func NewCommandQueue(s ...int) *CommandQueue {
	size := 1500
	if len(s) > 0 {
		size = s[0]
	}
	c := make([]common.Commands, size) // what does array default to? 0 maybe?
	cq := CommandQueue{
		Commands: c,
	}
	return &cq
}

// AddCommand adds a command to the queue
func (cq *CommandQueue) AddCommand(command common.Commands) {
	if cq.currentIdx < len(cq.Commands) {
		cq.Commands[cq.currentIdx] = command
	} else {
		cq.Commands = append(cq.Commands, command)
	}
	cq.currentIdx++
}

// LoadQueue reads a file and returns an array of brainfuck commands
func (cq *CommandQueue) LoadQueue(content []byte) {
	// Create commandList, loop through content, if item == soemthing golang, add it to commandlist
	for _, b := range content {
		s := string(b)
		if val, ok := common.CommandMap[s]; ok {
			cq.AddCommand(val)
		}
	}
}
