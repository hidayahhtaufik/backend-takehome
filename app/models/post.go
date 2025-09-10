package models

import "time"

type Post struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content" gorm:"type:TEXT"`
	AuthorID  uint64    `json:"author_id" gorm:"index"`
	Author    User      `json:"-" gorm:"foreignKey:AuthorID;references:ID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comments  []Comment `json:"comments,omitempty" gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;"`
}
