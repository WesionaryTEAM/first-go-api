package controller

import (
	"cloud-upload/imageUpload"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ImageUploadHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := header.Filename
	fmt.Println("FILENAME -> ", filename)
	defer file.Close()

	tempFile, err := os.Create("/tmp/" + filename)
	if err != nil {
		fmt.Println("File Create Error -->", err.Error())
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		fmt.Println("Read Error -->", err.Error())
	}

	imageUrl := imageUpload.ImageUpload("/tmp/" + filename)

	c.JSON(http.StatusOK, gin.H{"krakenURL": imageUrl})
}
