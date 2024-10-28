package middlewares

import (
	"RestApi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	// use token for authentication
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Not authorized, token is empty"})
		return
	}

	userID, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Not authorized, verify token"})
		return
	}

	// available for later requests
	context.Set("userId", userID)
	context.Next()
}
