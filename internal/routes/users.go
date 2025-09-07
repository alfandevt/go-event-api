package routes

import (
	"log"
	"net/http"

	"github.com/alfandevt/go-even-api/internal/models"
	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	err = user.Save()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success register user"})
}
