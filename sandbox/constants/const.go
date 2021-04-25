package constants

const NA = "N/A"

const (
	PASS = iota
	TimeOutError
	MemoryOutError
	RunTimeError
	UnknownError
)

const (
	AC  = "Accepted"
	WA  = "Wrong Answer"
	CE  = "Compile Error"
	TLE = "Time Limit Exceeded"
	MLE = "Memory Limit Exceeded"
	RE  = "Runtime Error"
	UE  = "Unknown Error"
)

const (
	CodeDir              = "./code"
	InputDir             = "./input"
	OutputDir            = "./output"
	DefaultExecuteTime   = 10000 //10s
	DefaultExecuteMemory = 65536 //65536kb
)
