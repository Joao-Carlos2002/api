package routes

import (
	"fmt"
	"net/http"

	"github.com/Joao-Carlos2002/Api-User/config"
	"github.com/Joao-Carlos2002/Api-User/models"
	"github.com/gin-gonic/gin"
)

func UpdateAccount(c *gin.Context) {
	db = config.Db
	user := models.User{}
	userU := models.User{}
	err := c.ShouldBind(&userU)
	if err != nil {
		fmt.Sprintln(err)
		c.String(http.StatusBadRequest, "%v", err)
		return
	}
	tx := db.Find(&user, userU.ID)
	if tx.RowsAffected == 0 {
		fmt.Println("Account not found")
		c.String(http.StatusBadRequest, "Account not found")
		return
	}
	db.UpdateColumns(userU)
	c.String(http.StatusOK, fmt.Sprintf("%v, %v", user, userU))
}
