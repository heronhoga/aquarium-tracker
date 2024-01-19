package unitcontroller

import (
	"encoding/json"
	"go-restapi-gin-aquarium/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GET ALL DATA
func Index(c *gin.Context) {
	var units []models.Unit
	result := models.DB.Raw("SELECT id, name, temperature, location, time FROM units").Scan(&units)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading units"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": units})
}

//GET SINGLE DATA
func IndexOne(c *gin.Context) {
	var unit models.Unit
	id := c.Param("id")

	result := models.DB.Raw("SELECT id, name, temperature, location, time FROM units WHERE id = ?", id).Scan(&unit)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading unit"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No unit found with given id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": unit})
}

//CREATE NEW DATA
func Create(c *gin.Context) {
	var unit models.Unit
	if err := c.ShouldBindJSON(&unit); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if unit.Name == "" || unit.Temperature == 0 || unit.Location == "" || unit.Time.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body must contain name, temperature, location, and time"})
		return
	}

	query := "INSERT INTO units (name, temperature, location, time) VALUES (?, ?, ?, ?)"
	models.DB.Exec(query, unit.Name, unit.Temperature, unit.Location, unit.Time)

	c.JSON(http.StatusOK, gin.H{"data": unit})
}

//UPDATE DATA
func Update(c *gin.Context) {
	var unit models.Unit
	id := c.Param("id")
	if err := c.ShouldBindJSON(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateQuery := "UPDATE units SET name = ?, temperature = ?, location = ?, time = ? WHERE id = ?"
	rowsAffected := models.DB.Exec(updateQuery, unit.Name, unit.Temperature, unit.Location, unit.Time, id).RowsAffected

	if rowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error updating unit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unit updated successfully","data": unit})
}

//DELETE DATA
func Delete(c *gin.Context) {
	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	id, _ := input.Id.Int64()

	deleteQuery := "DELETE FROM units WHERE id = ?"
	result := models.DB.Exec(deleteQuery, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting unit"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No unit found with the given id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unit deleted successfully"})
}
