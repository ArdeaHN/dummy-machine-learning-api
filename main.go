package main

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/detection", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		dir, err := os.Getwd()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get directory"})
		}

		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no such file"})
			return
		}

		fileDir := filepath.Join(dir, "images", file.Filename)

		err = c.SaveUploadedFile(file, fileDir)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "file failed to upload"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"is_success": true,
			"accuracy":   9.0,
			"result":     "red disease",
		})
	})
	r.Run(":8000")
}
