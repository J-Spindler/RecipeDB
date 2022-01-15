package gateway

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	db *gorm.DB
}

func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	userController := UserController{
		db: db,
	}

	r.GET("/users", userController.getUser)
	r.POST("/users", userController.createUser)
	r.DELETE("/users", userController.deleteUser)
}

func (controller UserController) getUser(c *gin.Context) {
	var users []User
	controller.db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func (controller UserController) createUser(c *gin.Context) {
	var input User

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result := controller.db.Omit("ID").Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusOK, input)
}

func (controller UserController) deleteUser(c *gin.Context) {
	id, _ := c.GetQuery("id")
	result := controller.db.Delete(&User{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
		return
	}
	c.JSON(http.StatusNoContent, "")
}
