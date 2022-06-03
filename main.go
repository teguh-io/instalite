package main

import (
	"instalite/configs"
	"instalite/controllers"
	"instalite/database"
	_ "instalite/docs"
	"instalite/repositories"
	"instalite/routers"
	"instalite/services"

	"github.com/gin-gonic/gin"
)

// @title           Instalite
// @version         1.0
// @description     Simple Photo social media API

// @contact.name   Teguh Ainul Darajat
// @contact.email  teguh@email.com

// @host      localhost:8080
// @BasePath  /
func main() {
	configs.GetConfig()
	server := gin.Default()
	db := database.GetDB()

	// User DI
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	userRouter := routers.NewUserRouter(server, userController)
	userRouter.Setup()

	// Photo DI
	photoRepo := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepo)
	photoController := controllers.NewPhotoController(photoService)
	photoRouter := routers.NewPhotoRouter(server, photoController)
	photoRouter.Setup()

	// Comment DI
	commentRepo := repositories.NewCommentRepository(db)
	commentService := services.NewCommentRepository(commentRepo)
	commentController := controllers.NewCommentController(commentService)
	commentRouter := routers.NewCommentRouter(server, commentController)
	commentRouter.Setup()

	// Social Media DI
	socialMediaRepo := repositories.NewSocialMediaRepository(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepo)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)
	socialMediaRouter := routers.NewSocialMediaRouter(server, socialMediaController)
	socialMediaRouter.Setup()

	// Swagger DI
	swaggerRouter := routers.NewSwaggerRouter(server)
	swaggerRouter.Start()

	server.Run(configs.App.AppPort)
}
