package queue

import (
	"errors"

	"../common"
)

// Item is an item in the command queue
type Item struct {
	Command common.Commands
	Depth   int
}

// CommandQueue loads commands into a queue
type CommandQueue struct {
	Commands     []Item
	currentIdx   int
	currentDepth int
}

// NewCommandQueue returns a CommandQueue of a given size. Defaults to 1500
func NewCommandQueue(s ...int) *CommandQueue {
	size := 1500
	if len(s) > 0 {
		size = s[0]
	}
	c := make([]Item, size)
	cq := CommandQueue{
		Commands: c,
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

	qi := Item{Command: command, Depth: cq.currentDepth}
	if cq.currentIdx < len(cq.Commands) {
		cq.Commands[cq.currentIdx] = qi
	} else {
		cq.Commands = append(cq.Commands, qi)
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
