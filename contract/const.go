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
	DefaultExecuteTime   = 5000
	DefaultExecuteMemory = 10000
)
