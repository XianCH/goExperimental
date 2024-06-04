package repository

import (
	"database/sql"
	"errors"

	"github.com/x14n/goExperimental/gin-crud-api/v1/helper"
	"github.com/x14n/goExperimental/gin-crud-api/v1/model"
)

type TagsRepositoryImpl struct {
	Db *sql.DB
}

func NewTagsRepositoryImpl(Db *sql.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

func (t *TagsRepositoryImpl) Save(tags model.Tags) {
	_, err := t.Db.Exec("INSERT INTO tags (name) VALUES (?)", tags.Name)
	helper.ErrorPanic(err)
}

func (t *TagsRepositoryImpl) Update(tags model.Tags) {
	_, err := t.Db.Exec("UPDATE tags SET name = ? WHERE id = ?", tags.Name, tags.ID)
	helper.ErrorPanic(err)
}

func (t *TagsRepositoryImpl) Delete(tagsId int) {
	_, err := t.Db.Exec("DELETE FROM tags WHERE id = ?", tagsId)
	helper.ErrorPanic(err)
}

func (t *TagsRepositoryImpl) FindById(tagsId int) (model.Tags, error) {
	var tag model.Tags
	err := t.Db.QueryRow("SELECT id, name FROM tags WHERE id = ?", tagsId).Scan(&tag.ID, &tag.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tag, errors.New("tag is not found")
		}
		return tag, err
	}
	return tag, nil
}

func (t *TagsRepositoryImpl) FindAll() []model.Tags {
	rows, err := t.Db.Query("SELECT id, name FROM tags")
	helper.ErrorPanic(err)
	defer rows.Close()

	var tags []model.Tags
	for rows.Next() {
		var tag model.Tags
		err := rows.Scan(&tag.ID, &tag.Name)
		helper.ErrorPanic(err)
		tags = append(tags, tag)
	}
	err = rows.Err()
	helper.ErrorPanic(err)

	return tags
}
