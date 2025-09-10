package repositories

import (
	"app/models"
	"app/utils"

	"gorm.io/gorm"
)

type UserRepository struct{ DB *gorm.DB }

func (r UserRepository) Create(u *models.User) error {
	if err := r.DB.Create(u).Error; err != nil {
		return utils.MapDBError(err)
	}
	return nil
}

func (r UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, utils.MapDBError(err)
	}
	return &user, nil
}
