package res

//ListData 用于返回列表数据
type ListData struct {
	Total int         `json:"total"`
	Items interface{} `json:"items"`
}

//Results 用于批量请求的返回数据
type Results struct {
	RequestResList []Result `json:"requestResList"`
}

//Result 批量请求时，每个请求对应的请求结果
type Result struct {
	ID     string      `json:"id"`
	ResMsg string      `json:"resMsg"`
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}
