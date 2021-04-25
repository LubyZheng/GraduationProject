//For executor
package result

import (
	"Gproject/sandbox/constants"
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
	r.Status = constants.PASS
	r.Output = output
	r.Detail = constants.NA
	r.ExecuteTime = executeTime
	r.ExecuteMemory = executeMemory
	return r.ConJson()
}

func (r ExecuteResult) PackTimeOutResult() string {
	//status作为标志位在程序内提前赋值，用来区分kill是因为超时还是超内存
	r.Output = constants.NA
	r.Detail = constants.NA
	r.ExecuteTime = constants.NA
	r.ExecuteMemory = constants.NA
	return r.ConJson()
}

func (r ExecuteResult) PackMemoryOutErrorResult() string {
	r.Status = constants.MemoryOutError
	r.Output = constants.NA
	r.Detail = constants.NA
	r.ExecuteTime = constants.NA
	r.ExecuteMemory = constants.NA
	return r.ConJson()
}

func (r ExecuteResult) PackRunTimeErrorResult(detail string) string {
	r.Status = constants.RunTimeError
	r.Output = constants.NA
	r.Detail = detail
	r.ExecuteTime = constants.NA
	r.ExecuteMemory = constants.NA
	return r.ConJson()
}

func (r ExecuteResult) PackUnknownErrorResult(detail string) string {
	r.Status = constants.UnknownError
	r.Output = constants.NA
	r.Detail = detail
	r.ExecuteTime = constants.NA
	r.ExecuteMemory = constants.NA
	return r.ConJson()
}

func (r ExecuteResult) ConJson() string {
	b, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

