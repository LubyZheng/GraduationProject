package contract

import "encoding/json"

type ExecuteResult struct {
	Outcome       int    `json:"outcome"`
	Detail        string `json:"detail"`
	ExecuteTime   string `json:"execute_time"`
	ExecuteMemory string `json:"execute_memory"`
}

func (r ExecuteResult) PackPassResult(executeTime, executeMemory string) string {
	r.Outcome = PASS
	r.Detail = "Your code pass!"
	r.ExecuteTime = executeTime
	r.ExecuteMemory = executeMemory
	return r.ConJson()
}

func (r ExecuteResult) PackTimeOutResult() string {
	r.Outcome = TIMEOUT
	r.Detail = "Your code can't finish executing in limited time!"
	r.ExecuteTime = NA
	r.ExecuteMemory = NA
	return r.ConJson()
}

func (r ExecuteResult) PackRunTimeErrorResult(detail string) string {
	r.Outcome = RunTimeError
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
