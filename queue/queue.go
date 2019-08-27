package queue

import "../common"

type queueItems struct {
	command common.Commands
	depth   int
}

// CommandQueue loads commands into a queue
type CommandQueue struct {
	commands     []queueItems
	currentIdx   int
	currentDepth int
}

// NewCommandQueue returns a CommandQueue of a given size. Defaults to 1500
func NewCommandQueue(s ...int) *CommandQueue {
	size := 1500
	if len(s) > 0 {
		size = s[0]
	}
	c := make([]queueItems, size)
	cq := CommandQueue{
		commands: c,
	}
	return &cq
}

func (cq *CommandQueue) calculateDepth(command common.Commands) {
	if cq.currentDepth == 0 && command == common.LoopEnd {
		// TODO: throw error
	} else if command == common.LoopEnd {
		cq.currentDepth--
	} else if command == common.LoopStart {
		cq.currentDepth++
	}
}

// AddCommand adds a command to the queue
func (cq *CommandQueue) AddCommand(command common.Commands) {
	cq.calculateDepth(command)
	qi := queueItems{command: command, depth: cq.currentDepth}
	if cq.currentIdx < len(cq.commands) {
		cq.commands[cq.currentIdx] = qi
	} else {
		cq.commands = append(cq.commands, qi)
	}
	cq.currentIdx++
}
