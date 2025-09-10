package routes

import (
	"app/controllers"
	"app/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"

	// Swagger
	docs "app/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupSwagger(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerfiles.Handler,
			ginSwagger.URL("/swagger/doc.json"),
		),
	)
}

func New(auth controllers.AuthController, post controllers.PostController, cm controllers.CommentController) http.Handler {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middlewares.HTTPError())

	// Auth
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)

	// Public posts/comments
	r.GET("/posts", post.List)
	r.GET("/posts/:id", post.GetByID)
	r.GET("/posts/:id/comments", cm.List)

	// Protected
	api := r.Group("/")
	api.Use(middlewares.AuthRequired())
	{
		api.POST("/posts", post.Create)
		api.PUT("/posts/:id", post.Update)
		api.DELETE("/posts/:id", post.Delete)
		api.POST("/posts/:id/comments", cm.Create)
	}

	// Swagger UI
	setupSwagger(r)

	return r
}
