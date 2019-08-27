package common

// Commands is an enum of Branfuck commands
type Commands int

const (
	// Left moves the operator left
	Left Commands = iota + 1
	// Right moves the operator right
	Right
	// Inc increases the current cell by 1
	Inc
	// Dec decreases the current cell by 1
	Dec
	// Get gets input from IO
	Get
	// Put puts current cell to IO
	Put
	// LoopStart denotes the start of a loop
	LoopStart
	// LoopEnd denotes the end of a loop
	LoopEnd
)
