package request

// CreateTagsRequest 结构体表示创建标签的请求数据结构
type CreateTagsRequest struct {
	Name string `json:"name" validate:"required"`
}

// UpdateTagsRequest 结构体表示更新标签的请求数据结构
type UpdateTagsRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
