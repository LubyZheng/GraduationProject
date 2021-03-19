package contract

const NA = "N/A"

const (
	PASS = iota
	WRONG
	CompileFail
	TIMEOUT
	MemoryOut
	RunTimeError
)

const (
	CodeDir              = "./code"
	InputDir             = "./input"
	OutputDir            = "./output"
	DefaultExecuteTime   = 5000
	DefaultExecuteMemory = 10000
)
