package repositories

import (
	"app/models"
	"app/utils"

	"gorm.io/gorm"
)

type CommentRepository struct{ DB *gorm.DB }

func (r CommentRepository) Create(c *models.Comment) error {
	if err := r.DB.Create(c).Error; err != nil {
		return utils.MapDBError(err)
	}
	return nil
}

func (r CommentRepository) ListByPost(postID uint64) ([]models.Comment, error) {
	var data []models.Comment
	return data, r.DB.Where("post_id = ?", postID).Order("id ASC").Find(&data).Error
}
