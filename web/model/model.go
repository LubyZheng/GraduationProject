package model

type Request struct {
	QuestionID string `json:"questionid"`
	Language   string `json:"language"`
	Code       string `json:"code"`
}

type Response struct {
	Status        string `json:"status"`
	Detail        string `json:"detail"`
	CompileTime   string `json:"compile_time(ms)"` //CPU时间，不包含sleep,阻塞等
	CompileMemory string `json:"compile_memory(kb)"`
	ExecuteTime   string `json:"execute_time(ms)"` //CPU时间，不包含sleep,阻塞等
	ExecuteMemory string `json:"execute_memory(kb)"`
}
