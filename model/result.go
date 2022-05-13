package model

// 返回结果结构
type Result struct {
	Code    int         `json:"code"`    // 返回状态码
	Message string      `json:"message"` // 返回信息
	Result  interface{} `json:"result"`  // 返回内容
}
