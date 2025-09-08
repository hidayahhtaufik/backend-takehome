package models

import "time"

type User struct {
	ID           uint64    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Email        string    `json:"email" gorm:"uniqueIndex"` // unique email
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Blogs        []Blog    `json:"-" gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE;"`
}
