package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.comlunatictiol/rest-api-with-go/model"
	"github.comlunatictiol/rest-api-with-go/utils"
)

func signup(contex *gin.Context) {
	var user model.User
	err := contex.ShouldBindJSON(&user)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"message": "invalid data", "err": err})
		return
	}

	err = user.Save()
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"message": "cannot save user", "err": err})
		return
	}
	contex.JSON(http.StatusCreated, gin.H{"message": "signup successful"})

}

func login(contex *gin.Context) {
	var user model.User
	err := contex.ShouldBindJSON(&user)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"message": "invalid data", "err": err})
		return
	}

	err = user.Validate()

	if err != nil {
		contex.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials", "err": err})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"message": "cannot generate token"})
		return
	}

	contex.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})

}
