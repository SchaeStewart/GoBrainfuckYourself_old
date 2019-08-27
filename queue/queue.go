package queue

import (
	"errors"

	"../common"
)

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

func (cq *CommandQueue) calculateDepth(command common.Commands) error {
	if cq.currentDepth == 0 && command == common.LoopEnd {
		return errors.New("Invalid syntax. Cannot end loop ")
	} else if command == common.LoopEnd {
		cq.currentDepth--
	} else if command == common.LoopStart {
		cq.currentDepth++
	}
	return nil
}

// AddCommand adds a command to the queue
func (cq *CommandQueue) AddCommand(command common.Commands) {
	err := cq.calculateDepth(command)
	if err != nil {
		panic(err)
	}

	qi := queueItems{command: command, depth: cq.currentDepth}
	if cq.currentIdx < len(cq.commands) {
		cq.commands[cq.currentIdx] = qi
	} else {
		cq.commands = append(cq.commands, qi)
	}
	cq.currentIdx++
}
