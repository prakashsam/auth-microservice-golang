package route

import (
	"authservice/controllers"
	"authservice/db"
	"authservice/repository"
	"authservice/services"

	"github.com/kataras/iris/v12"
)

func SetupRoutes(app *iris.Application, jwtSecret string) {
	userRepo := &repository.UserRepository{DB: db.GetDB()}
	authService := &services.AuthService{Repo: userRepo, JWTSecret: jwtSecret}
	authController := &controllers.AuthController{Service: authService}

	app.Post("/auth/register", authController.Register)
	app.Post("/auth/login", authController.Login)
}
