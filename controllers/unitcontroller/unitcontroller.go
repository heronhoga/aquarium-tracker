package unitcontroller

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func IndexOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
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

