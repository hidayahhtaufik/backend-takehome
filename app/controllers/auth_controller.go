package controllers

import (
	"net/http"
	"os"

	"app/dto"
	"app/usecases"
	"app/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct{ UC usecases.AuthUsecase }

// Register godoc
// @Summary Register new user
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body dto.RegisterRequest true "register payload"
// @Success 201 {object} map[string]any
// @Failure 400 {object} map[string]string
// @Router /register [post]
func (a AuthController) Register(c *gin.Context) {
	var in dto.RegisterRequest
	if err := c.ShouldBindJSON(&in); err != nil {
		c.Error(err)
		return
	}
	u, err := a.UC.Register(in)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": u.ID, "name": u.Name, "email": u.Email})
}

// Login godoc
// @Summary Login and get JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body dto.LoginRequest true "login payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func (a AuthController) Login(c *gin.Context) {
	var in dto.LoginRequest
	if err := c.ShouldBindJSON(&in); err != nil {
		c.Error(err)
		return
	}
	u, err := a.UC.Verify(in)
	if err != nil {
		c.Error(err)
		return
	}
	token, err := utils.SignJWT(os.Getenv("JWT_SECRET"), u.ID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
