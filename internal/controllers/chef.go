package controllers

import (
	"net/http"

	"github.com/altschool/go-app/internal/models"
	"github.com/altschool/go-app/internal/repo"
	"github.com/gin-gonic/gin"
)

func NewChef(c *gin.Context) {
	var chef models.Chef

	if err := c.ShouldBindJSON(&chef); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo.NewChef(chef)

	c.JSON(http.StatusOK, chef)
}

func ListChefs(c *gin.Context) {
	chefs := repo.ListChefs()
	c.JSON(http.StatusOK, chefs)
}

func UpdateChef(c *gin.Context) {
	id := c.Param("chef-id")

	var chef models.Chef

	if err := c.ShouldBindJSON(&chef); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ok, r := repo.UpdateChef(chef, id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "chef not found",
		})
		return
	}

	c.JSON(http.StatusOK, &r)
}

func DeleteChef(c *gin.Context) {
	id := c.Param("chef-id")
	if !repo.DeleteChef(id) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "chef not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "chef deleted",
	})
}
