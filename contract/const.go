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
	DefaultExecuteTime   = 10000 //10s
	DefaultExecuteMemory = 65536 //65536kb
)
