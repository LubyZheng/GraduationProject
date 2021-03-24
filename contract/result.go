package contract

import "encoding/json"

type Result struct {
	Status        string `json:"status"`
	Detail        string `json:"detail"`
	CompileTime   string `json:"compile_time(ms)"` //CPU时间，不包含sleep,阻塞等
	CompileMemory string `json:"compile_memory(kb)"`
	ExecuteTime   string `json:"execute_time(ms)"` //CPU时间，不包含sleep,阻塞等
	ExecuteMemory string `json:"execute_memory(kb)"`
}

func (r Result) PackPassResult(ExecuteTime, ExecuteMemory string) string {
	r.Status = "Accepted"
	r.ExecuteTime = ExecuteTime
	r.ExecuteMemory = ExecuteMemory
	return r.ConJson()
}

func (r Result) PackWrongResult(ExecuteTime, ExecuteMemory string) string {
	r.Status = "Wrong Answer"
	r.ExecuteTime = ExecuteTime
	r.ExecuteMemory = ExecuteMemory
	return r.ConJson()
}

func (r Result) PackCompileFailResult(Detail string) string {
	r.Status = "Compile Error"
	r.Detail = Detail
	return r.ConJson()
}

func (r Result) PackTimeOutErrorResult() string {
	r.Status = "Time Limit Exceeded"
	return r.ConJson()
}

func (r Result) PackMemoryOutErrorResult() string {
	r.Status = "Memory Limit Exceeded"
	return r.ConJson()
}

func (r Result) PackRunTimeErrorResult(Detail string) string {
	r.Status = "Runtime Error"
	r.Detail = Detail
	return r.ConJson()
}

func (r Result) PackUnknownErrorResult(Detail string) string {
	r.Status = "Unknown error"
	r.Detail = Detail
	return r.ConJson()
}

func (r Result) ConJson() string {
	b, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
