package repository

import (
	"github.com/x14n/goExperimental/gin-crud-api/v1/model"
)

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(tagsId int)
	FindById(tagsId int) (tags model.Tags, err error)
	FindAll() []model.Tags
}
