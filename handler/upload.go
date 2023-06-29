package handler

import (
	"damage/logic"
	"github.com/gin-gonic/gin"
)

func UploadAdventureExcel(c *gin.Context) {

	common, err := logic.UploadAdventureExcel(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, common)

}
