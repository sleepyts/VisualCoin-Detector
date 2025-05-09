package model

// 外层的响应结构
type ApiResponse struct {
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data []TickerData `json:"data"` // 内层的数组
}
