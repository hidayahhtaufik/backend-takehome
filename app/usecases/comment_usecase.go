package usecases

import (
	"time"

	"app/dto"
	"app/models"
	"app/repositories"
)

type CommentUsecase struct {
	Comments repositories.CommentRepository
}

func (u CommentUsecase) Create(postID uint64, in dto.CommentCreateRequest) (*models.Comment, error) {
	c := &models.Comment{PostID: postID, AuthorName: in.AuthorName, Content: in.Content, CreatedAt: time.Now()}
	if err := u.Comments.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (u CommentUsecase) ListByPost(postID uint64) ([]models.Comment, error) {
	return u.Comments.ListByPost(postID)
}
