package controllers

import (
	"net/http"
	"strconv"

	"app/dto"
	"app/middlewares"
	"app/usecases"

	"github.com/gin-gonic/gin"
)

type PostController struct{ UC usecases.PostUsecase }

// CreatePost godoc
// @Summary Create post
// @Tags posts
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param payload body dto.PostCreateRequest true "post payload"
// @Success 201 {object} models.Post
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /posts [post]
func (b PostController) Create(c *gin.Context) {
	uid := c.MustGet(middlewares.CtxUID).(uint64)
	var req dto.PostCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	post, err := b.UC.Create(uid, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, post) // 201
}

// GetPost godoc
// @Summary Get post by ID
// @Tags posts
// @Produce json
// @Param id path int true "post id"
// @Success 200 {object} models.Post
// @Failure 404 {object} map[string]string
// @Router /posts/{id} [get]
func (b PostController) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	data, err := b.UC.GetByID(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// ListPosts godoc
// @Summary List all posts
// @Tags posts
// @Produce json
// @Success 200 {array} models.Post
// @Router /posts [get]
func (b PostController) List(c *gin.Context) {
	data, err := b.UC.List()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// UpdatePost godoc
// @Summary Update post
// @Tags posts
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "post id"
// @Param payload body dto.PostUpdateRequest true "update payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /posts/{id} [put]
func (b PostController) Update(c *gin.Context) {
	uid := c.MustGet(middlewares.CtxUID).(uint64)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req dto.PostUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	if err := b.UC.Update(uid, id, req); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success update blog post"})
}

// DeletePost godoc
// @Summary Delete post
// @Tags posts
// @Security BearerAuth
// @Param id path int true "post id"
// @Success 204 "no content"
// @Failure 401 {object} map[string]string
// @Router /posts/{id} [delete]
func (b PostController) Delete(c *gin.Context) {
	uid := c.MustGet(middlewares.CtxUID).(uint64)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := b.UC.Delete(id, uid); err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
