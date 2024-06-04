package response

// Response 结构体表示 HTTP 响应的数据结构
type Response struct {
	Code   int         `json:"code"`   // 响应状态码
	Status string      `json:"status"` // 响应状态描述
	Data   interface{} `json:"data"`   // 响应数据
}
