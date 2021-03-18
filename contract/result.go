package contract

import "encoding/json"

type Result struct {
	Outcome       string `json:"outcome"`
	Detail        string `json:"detail"`
	CompileTime   string `json:"compile_time(ms)"`
	CompileMemory string `json:"compile_memory(kb)"`
	ExecuteTime   string `json:"execute_time(ms)"`
	ExecuteMemory string `json:"execute_memory(kb)"`
}

func (r Result) PackCompileFailResult() string {
	r.Outcome = "false"
	r.Detail = "Compile Failed!"
	r.ExecuteTime = NA
	r.ExecuteMemory = NA
	return r.ConJson()
}

func (r Result) ConJson() string {
	b, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
