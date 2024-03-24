package routes

import (
	"fmt"
	"net/http"

	"github.com/Joao-Carlos2002/Api-User/config"
	"github.com/Joao-Carlos2002/Api-User/models"
	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	db = config.Db
	user := models.User{}
	userV := models.User{}
	// Get user for method POST
	err := c.ShouldBind(&user)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, fmt.Sprint(err))
	}

	db.Where("email = ?", user.Email).Find(&userV)

	// Verify if email is already in use
	if userV.Email == user.Email {
		fmt.Println("Email is already in use")
		c.String(http.StatusBadRequest, "Email is already in use")
		return
	}
	// Create user in Database
	tx := db.Create(&user)
	if tx.RowsAffected == 0 {
		fmt.Println("Account not created")
		c.String(http.StatusBadRequest, "Account not created")
		return
	}
	c.String(http.StatusCreated, "Account created")
}
