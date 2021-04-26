//For judger
package judger

import (
	"Gproject/sandbox/constants"
	"bytes"
	"encoding/json"
)

type Result struct {
	Status        string `json:"status"`
	Detail        string `json:"detail"`
	CompileTime   string `json:"compile_time(ms)"` //CPU时间，不包含sleep,阻塞等
	CompileMemory string `json:"compile_memory(kb)"`
	ExecuteTime   string `json:"execute_time(ms)"` //CPU时间，不包含sleep,阻塞等
	ExecuteMemory string `json:"execute_memory(kb)"`
}

func (r Result) PackPassResult(ExecuteTime, ExecuteMemory string) string {
	r.Status = constants.AC
	r.Detail = constants.NA
	r.ExecuteTime = ExecuteTime
	r.ExecuteMemory = ExecuteMemory
	return r.ConJson()
}

func (r Result) PackWrongResult(ExecuteTime, ExecuteMemory string) string {
	r.Status = constants.WA
	r.Detail = constants.NA
	r.ExecuteTime = ExecuteTime
	r.ExecuteMemory = ExecuteMemory
	return r.ConJson()
}

func (r Result) PackCompileFailResult(Detail string) string {
	r.Status = constants.CE
	r.Detail = Detail
	return r.ConJson()
}

func (r Result) PackTimeOutErrorResult() string {
	r.Status = constants.TLE
	r.Detail = constants.NA
	return r.ConJson()
}

func (r Result) PackMemoryOutErrorResult() string {
	r.Status = constants.MLE
	r.Detail = constants.NA
	return r.ConJson()
}

func (r Result) PackRunTimeErrorResult(Detail string) string {
	r.Status = constants.RE
	r.Detail = Detail
	return r.ConJson()
}

func (r Result) PackUnknownErrorResult(Detail string) string {
	r.Status = constants.UE
	r.Detail = Detail
	return r.ConJson()
}

func (r Result) ConJson() string {
	b, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(out.Bytes())
}
