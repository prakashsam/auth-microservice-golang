package controllers

import (
	"authservice/models"
	"authservice/services"

	"github.com/kataras/iris/v12"
)

type AuthController struct {
	Service *services.AuthService
}

func (c *AuthController) Register(ctx iris.Context) {
	var user models.User
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}
	if err := c.Service.Register(user); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Registration failed"})
		return
	}
	ctx.JSON(iris.Map{"message": "User registered successfully"})
}

func (c *AuthController) Login(ctx iris.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ReadJSON(&input); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}
	token, err := c.Service.Login(input.Email, input.Password)
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Invalid credentials"})
		return
	}
	ctx.JSON(iris.Map{"token": token})
}
