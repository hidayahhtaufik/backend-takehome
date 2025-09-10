// @title Simple Blog API
// @version 1.0
// @description REST API for users, blogs (posts) and comments
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"app/config"
	"app/controllers"
	"app/repositories"
	"app/routes"
	"app/usecases"
)

func main() {
	db, err := config.OpenGormWithRetry(context.Background(), 90*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Repos
	userRepo := repositories.UserRepository{DB: db}
	postRepo := repositories.PostRepository{DB: db}
	cmRepo := repositories.CommentRepository{DB: db}

	// Usecases
	authUC := usecases.AuthUsecase{Users: userRepo}
	postUC := usecases.PostUsecase{Posts: postRepo}
	cmUC := usecases.CommentUsecase{Comments: cmRepo}

	// Controllers
	auth := controllers.AuthController{UC: authUC}
	post := controllers.PostController{UC: postUC}
	cm := controllers.CommentController{UC: cmUC}

	h := routes.New(auth, post, cm)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", h))
}
