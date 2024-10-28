package routes

import (
	"RestApi/models"
	"RestApi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not create user. Try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"Message": "User created!!"})

}

func login(context *gin.Context) {
	var user models.User

	// user ID is not populated
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request"})
		return
	}

	// user ID is populated here
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Loggin successful!!", "token": token})
}
