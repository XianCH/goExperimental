package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/x14n/goExperimental/gin-crud-api/v1/data/request"
	"github.com/x14n/goExperimental/gin-crud-api/v1/data/response"
	"github.com/x14n/goExperimental/gin-crud-api/v1/helper"
	"github.com/x14n/goExperimental/gin-crud-api/v1/model"
	"github.com/x14n/goExperimental/gin-crud-api/v1/repository"
)

type TagsServiceImpl struct {
	TagRepository repository.TagsRepository
	Validate      *validator.Validate
}

func NewTagServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagRepository: tagRepository,
		Validate:      validate,
	}
}

func (t TagsServiceImpl) Create(tag request.CreateTagsRequest) {
	err := t.Validate.Struct(tag)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tag.Name,
	}
	t.TagRepository.Save(tagModel)
}

func (t TagsServiceImpl) Update(tag request.UpdateTagsRequest) {
	tagData, err := t.TagRepository.FindById(tag.Id)
	helper.ErrorPanic(err)
	tagData.Name = tag.Name
	t.TagRepository.Update(tagData)
}

func (t TagsServiceImpl) Delete(tagId int) {
	t.TagRepository.Delete(tagId)
}

func (t TagsServiceImpl) FindById(tagId int) response.TagsResponse {
	tagData, err := t.TagRepository.FindById(tagId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		ID:   tagData.ID,
		Name: tagData.Name,
	}
	return tagResponse
}

func (t TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			ID:   value.ID,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}
	return tags
}
