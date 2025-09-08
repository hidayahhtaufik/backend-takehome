package models

import "time"

type Comment struct {
	ID         uint64    `json:"id"`
	PostID     uint64    `json:"post_id" gorm:"foreignKey:PostID"`
	AuthorName string    `json:"author_name"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}
