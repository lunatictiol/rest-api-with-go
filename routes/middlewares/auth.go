package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.comlunatictiol/rest-api-with-go/utils"
)

func Aunthenticate(contex *gin.Context) {
	token := contex.Request.Header.Get("Authorization")
	if token == "" {
		contex.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "empty token"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		contex.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorised token"})
		return
	}
	contex.Set("userId", userId)
	contex.Next()

}
