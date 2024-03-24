package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Joao-Carlos2002/Api-User/config"
	"github.com/Joao-Carlos2002/Api-User/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetAccount(c *gin.Context) {
	db = config.Db
	id := c.Param("id")
	idF, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	var user models.User
	db.Find(&user, idF)

	if user.ID == 0 {
		c.String(http.StatusNotFound, "Account not found")
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}
