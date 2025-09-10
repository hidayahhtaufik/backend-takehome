package usecases

import (
	"time"

	"app/dto"
	"app/models"
	"app/repositories"
)

type PostUsecase struct{ Posts repositories.PostRepository }

func (u PostUsecase) Create(authorID uint64, req dto.PostCreateRequest) (*models.Post, error) {
	b := &models.Post{Title: req.Title, Content: req.Content, AuthorID: authorID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := u.Posts.Create(b); err != nil {
		return nil, err
	}
	return b, nil
}

func (u PostUsecase) Update(authorID, id uint64, req dto.PostUpdateRequest) error {
	b := &models.Post{ID: id, Title: req.Title, Content: req.Content, AuthorID: authorID}
	return u.Posts.Update(b)
}

func (u PostUsecase) GetByID(id uint64) (*models.Post, error) {
	return u.Posts.GetByID(id)
}

func (u PostUsecase) List() ([]models.Post, error) {
	return u.Posts.List()
}

func (u PostUsecase) Delete(id, authorID uint64) error {
	return u.Posts.Delete(id, authorID)
}
