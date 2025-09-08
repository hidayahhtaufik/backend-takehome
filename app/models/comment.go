package models

import "time"

type Comment struct {
	ID         uint64    `json:"id" gorm:"primaryKey"`
	BlogID     uint64    `json:"blog_id" gorm:"index"`
	Blog       Blog      `json:"-" gorm:"foreignKey:BlogID;references:ID;constraint:OnDelete:CASCADE;"`
	AuthorName string    `json:"author_name"`
	Content    string    `json:"content" gorm:"type:TEXT"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
