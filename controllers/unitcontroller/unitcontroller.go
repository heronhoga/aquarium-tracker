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
	var unit models.Unit
	if err := c.ShouldBindJSON(&unit); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&unit)
	c.JSON(http.StatusOK, gin.H{"data": unit})
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}

