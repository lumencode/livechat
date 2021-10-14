package controllers
	
import (
	"log"
	"fmt"
	"strings"
	"crypto/sha1"

	"github.com/lumencode/livechat/util"
	"github.com/lumencode/livechat/models"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {

	db, err := util.Pg()

	if err != nil {
		log.Print(err)
		c.JSON(503, nil)
		return
	}

	username := c.PostForm("username")
	email := c.PostForm("email")

	var user models.User

	row := db.Select("user_id").Find(
		&user,
		"username = lower(?) OR email = lower(?)",
		strings.ToLower(username),
		strings.ToLower(email),
	)

	if row.RowsAffected == 1 {
		c.JSON(400, nil)
		return
	}

	password := fmt.Sprintf("%x", sha1.Sum([]byte(c.PostForm("password"))))

	newUser := models.User{ Username: username, Password: password, Email: email, }

	result := db.Create(&newUser)

	if result.RowsAffected == 1 && result.Error == nil {
		c.JSON(201, newUser.Id)
		return
	}

	c.JSON(500, nil)
}
