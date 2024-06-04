package service

import (
	"github.com/x14n/goExperimental/gin-crud-api/v1/data/request"
	"github.com/x14n/goExperimental/gin-crud-api/v1/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
