package app_auth

import (
	"fmt"

	app_users "github.com/DenysNykoriak/go-api/app/users"
	"github.com/DenysNykoriak/go-api/core"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func InitializeRoutes(router *gin.Engine) {
	router.POST("/auth/sign-up", signUp)
}

func signUp(c *gin.Context) {
	body := SignUpBody{}

	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	valid, errs := body.Validate()

	if !valid {
		c.JSON(400, gin.H{"error": errs})
		return
	}

	var hashedPassword []byte

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser := app_users.User{
		FirstName:      body.FirstName,
		LastName:       body.LastName,
		Email:          body.Email,
		HashedPassword: string(hashedPassword),
	}

	result := core.PostgresDB.Create(&newUser)

	if result.Error != nil {
		var errorMessage string

		fmt.Println(result.Error)

		switch result.Error.(type) {
		default:
			errorMessage = "Failed to create user"
		}

		c.JSON(500, gin.H{"error": errorMessage})
		return
	}

	userDTO := newUser.ToDTO()

	c.JSON(200, gin.H{
		"status": "success",
		"user":   userDTO,
	})
}
