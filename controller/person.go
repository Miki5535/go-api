package controller

import (
	"go-grom/dto"
	"go-grom/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func PersonController(router *gin.Engine) {
	routes := router.Group("/person")
	{
		routes.GET("/", getAllPersons)
		routes.POST("/", createuser)
	}
}

func createuser(c *gin.Context) {
	user := dto.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// c.JSON(200, user)
	person := model.Person{}
	copier.Copy(&person, &user)
	c.JSON(200, person)

}

func getAllPersons(c *gin.Context) {
	person1 := model.Person{
		ID:          1,
		FirstName:   "Potchara",
		LastName:    "P",
		Age:         30,
		Email:       "potchara@example.com",
		Password:    "password123",
		PostAddress: model.Address{HouseNo: "123", City: "Bangkok"},
	}
	person2 := model.Person{
		ID:          2,
		FirstName:   "John",
		LastName:    "Doe",
		Age:         25,
		Email:       "john.doe@example.com",
		Password:    "securepass",
		PostAddress: model.Address{HouseNo: "456", City: "New York"},
	}
	persons := []model.Person{person1, person2}
	c.JSON(200, persons)
}
