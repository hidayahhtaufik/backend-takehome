package controllers

import (
	"net/http"
	"strconv"

	"app/dto"
	"app/usecases"

	"github.com/gin-gonic/gin"
)

type CommentController struct{ UC usecases.CommentUsecase }

// CreateComment godoc
// @Summary Add comment to a blog
// @Tags comments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "post id"
// @Param payload body dto.CommentCreateRequest true "comment payload"
// @Success 201 {object} models.Comment
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /posts/{id}/comments [post]
func (cc CommentController) Create(c *gin.Context) {
	postID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req dto.CommentCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	data, err := cc.UC.Create(postID, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, data)
}

// ListComments godoc
// @Summary List comments for a post
// @Tags comments
// @Produce json
// @Param id path int true "post id"
// @Success 200 {array} models.Comment
// @Router /posts/{id}/comments [get]
func (cc CommentController) List(c *gin.Context) {
	postID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	data, err := cc.UC.ListByPost(postID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, data)
}
