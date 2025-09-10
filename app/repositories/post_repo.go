package repositories

import (
	"errors"

	"app/models"
	"app/utils"

	"gorm.io/gorm"
)

type PostRepository struct{ DB *gorm.DB }

var ErrNotFoundOrNotOwner = errors.New("post not found or not owned by user")

func (r PostRepository) Create(p *models.Post) error {
	if err := r.DB.Create(p).Error; err != nil {
		return utils.MapDBError(err)
	}
	return nil
}

func (r PostRepository) GetByID(id uint64) (*models.Post, error) {
	var data models.Post
	if err := r.DB.
		Preload("Comments").
		First(&data, id).Error; err != nil {
		return nil, utils.MapDBError(err)
	}
	return &data, nil
}

func (r PostRepository) List() ([]models.Post, error) {
	var data []models.Post
	return data, r.DB.Preload("Comments").Order("id DESC").Find(&data).Error
}

func (r PostRepository) Update(b *models.Post) error {
	res := r.DB.Model(&models.Post{}).
		Where("id = ? AND author_id = ?", b.ID, b.AuthorID).
		Updates(map[string]any{"title": b.Title, "content": b.Content})

	if res.Error != nil {
		return utils.MapDBError(res.Error)
	}
	if res.RowsAffected == 0 {
		return ErrNotFoundOrNotOwner
	}
	return nil
}

func (r PostRepository) Delete(id, authorID uint64) error {
	res := r.DB.Where("id = ? AND author_id = ?", id, authorID).Delete(&models.Post{})
	if res.Error != nil {
		return utils.MapDBError(res.Error)
	}
	if res.RowsAffected == 0 {
		return ErrNotFoundOrNotOwner
	}
	return nil
}
