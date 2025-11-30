package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/markHiarley/projetinho/internal/controller"
	"github.com/markHiarley/projetinho/internal/middleware"
	"github.com/markHiarley/projetinho/internal/services"
	"github.com/markHiarley/projetinho/internal/usecase"
	"github.com/markHiarley/projetinho/pkg/postgres"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	userService := services.NewUserService(db)
	loginService := services.NewLoginService(db)

	userUseCase := usecase.NewUserUseCase(*userService)
	loginUseCase := usecase.NewLoginUseCase(*loginService)

	userController := controller.NewUserController(userUseCase)
	loginController := controller.NewLoginController(loginUseCase)

	router := gin.Default()

	api := router.Group("/api")
	{

		api.POST("/users", userController.CreateUser)
		api.POST("/login", loginController.AuthenticateUser)

	}

	adminRoutes := router.Group("/api")
	adminRoutes.Use(middleware.ValidateTokenJwt(), middleware.AdminMiddleware())
	{
		adminRoutes.DELETE("/users/:id", userController.DeleteUser)
	}

	protectedChat := router.Group("/chatonline/ws")
	protectedChat.Use(middleware.ValidateTokenJwt())

	{
		handleConnections := services.HandleConnections
		handleMessages := services.HandleMessages

		protectedChat.GET("", func(c *gin.Context) {
			handleConnections(c.Writer, c.Request)
		})
		go handleMessages()
	}

	log.Printf(" Server starting on :19090")
	if err := router.Run(":19090"); err != nil {
		log.Fatalf(" Server failed (deu ruim): %v", err)
	}
}
