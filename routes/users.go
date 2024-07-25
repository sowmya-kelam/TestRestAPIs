package routes

import (
	"restapi/models"
	"restapi/utils"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context, database *sql.DB) {

	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = user.Save(database)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created "})

}

func Login(context *gin.Context, database *sql.DB) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = user.Validate(database)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return

	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successfull", "token": token})

}
