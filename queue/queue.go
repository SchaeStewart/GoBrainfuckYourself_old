package queue

import "../common"

type queueItems struct {
	command common.Commands
	depth   int
}

// CommandQueue loads commands into a queue
type CommandQueue struct {
	commands []queueItems
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
