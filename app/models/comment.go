package models

import "time"

type Comment struct {
	ID         uint64    `json:"id" gorm:"primaryKey"`
	PostID     uint64    `json:"post_id" gorm:"index"`
	Post       Post      `json:"-" gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE;"`
	AuthorName string    `json:"author_name"`
	Content    string    `json:"content" gorm:"type:TEXT"`
	CreatedAt  time.Time `json:"created_at"`
}
