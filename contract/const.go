package contract

const NA = "N/A"

const (
	PASS = iota
	WRONG
	CompileFailError
	TimeOutError
	MemoryOutError
	RunTimeError
	UnknownError
)

const (
	CodeDir              = "./code"
	InputDir             = "./input"
	OutputDir            = "./output"
	DefaultExecuteTime   = 5000
	DefaultExecuteMemory = 10000
)
