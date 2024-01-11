package unitcontroller

import (
	"go-restapi-gin-aquarium/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var units []models.Unit
	result := models.DB.Find(&units)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading units"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": units})
}

func IndexOne(c *gin.Context) {
	var unit models.Unit
	result := models.DB.First(&unit, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading unit"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": unit})
}

func Create(c *gin.Context) {

}

func Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

