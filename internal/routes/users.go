package routes

import (
	"net/http"

	"github.com/alfandevt/go-even-api/internal/models"
	"github.com/alfandevt/go-even-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	var signUpRequest models.SignUpDTO
	var user models.User

	err := ctx.ShouldBindBodyWithJSON(&signUpRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	user.Email = signUpRequest.Email
	user.Username = signUpRequest.Username
	user.Password = signUpRequest.Password

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success register user"})
}

func signIn(ctx *gin.Context) {
	var signInRequest models.SignInDTO
	var user models.User

	err := ctx.ShouldBindBodyWithJSON(&signInRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	user.Email = signInRequest.Email
	user.Password = signInRequest.Password
	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	jwtToken, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "signed in success", "token": jwtToken})
}
