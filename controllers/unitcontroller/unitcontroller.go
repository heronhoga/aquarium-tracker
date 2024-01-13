package unitcontroller

import (
	"encoding/json"
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
	var unit models.Unit
	id := c.Param("id")
	if err := c.ShouldBindJSON(&unit); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if models.DB.Model(&unit).Where("id = ?", id).Updates(&unit).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating unit"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Unit updated successfully","data": unit})
}

func Delete(c *gin.Context) {
	var unit models.Unit
	
	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	
	id, _ := input.Id.Int64()
if err := models.DB.Delete(&unit, id).Error; err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting unit"})
	return
}
c.JSON(http.StatusOK, gin.H{"message": "Unit deleted successfully"})
}
