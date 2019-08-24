package common

// Commands is an enum of Branfuck commands
type Commands int

const (
	// Left moves the operator left
	Left Commands = 0
	// Right moves the operator right
	Right Commands = 1
	// Inc increases the current cell by 1
	Inc Commands = 2
	// Dec decreases the current cell by 1
	Dec Commands = 3
	// Get gets input from IO
	Get Commands = 4
	// Put puts current cell to IO
	Put Commands = 5
	// LoopStart denotes the start of a loop
	LoopStart Commands = 6
	// LoopEnd denotes the end of a loop
	LoopEnd Commands = 7
)
