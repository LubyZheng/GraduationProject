package contract

import (
	"encoding/json"
)

type ExecuteResult struct {
	Status        int    `json:"status"`
	Output        string `json:"output"`
	Detail        string `json:"detail"`
	ExecuteTime   string `json:"execute_time"`
	ExecuteMemory string `json:"execute_memory"`
}

func (r ExecuteResult) PackPassResult(output, executeTime, executeMemory string) string {
	r.Status = PASS
	r.Output = output
	r.Detail = "Your code pass!"
	r.ExecuteTime = executeTime
	r.ExecuteMemory = executeMemory
	return r.ConJson()
}

func (r ExecuteResult) PackTimeOutResult() string {
	r.Status = TIMEOUT
	r.Output = ""
	r.Detail = ""
	r.ExecuteTime = NA
	r.ExecuteMemory = NA
	return r.ConJson()
}

func (r ExecuteResult) PackRunTimeErrorResult(detail string) string {
	r.Status = RunTimeError
	r.Output = ""
	r.Detail = detail
	r.ExecuteTime = NA
	r.ExecuteMemory = NA
	return r.ConJson()
}

func (r ExecuteResult) ConJson() string {
	b, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
