package controller

// import (
// 	"bonoo_clin_patho_api/api/responses"
// 	"cloud-upload/imageUpload"
// 	"cloud-upload/models"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"path"
// 	"path/filepath"

// 	"github.com/gin-gonic/gin"
// )

// //Adding a person
// func CreatePerson(c *gin.Context) {
// 	// db := c.MustGet("db").(*gorm.DB)
// 	var person models.Person

// 	// ParseMultipartForm parses a request body as multipart/form-data
// 	//f := c.Request.ParseMultipartForm(32 << 20)

// 	f, err := c.MultipartForm()
// 	if err != nil {
// 		responses.ERROR(c, http.StatusInternalServerError, err)
// 		return
// 	}

// 	person = models.Person{}

// 	errForm := c.Request.ParseForm()
// 	log.Println(errForm)

// 	photo := f.File["photo"]

// 	photoName := path.Join(filepath.Base(photo.Filename))

// 	file, _, err := c.Request.FormFile("photo")

// 	switch err {
// 	case nil:
// 		defer file.Close()

// 		photo = imageUpload.ImageUpload(photoName)
// 		if photo == "" {
// 			fmt.Println("error getting path of image")
// 			return
// 		}
// 	default:
// 		photo = ""
// 		log.Println("Error getting image", err)
// 		return
// 	}

// 	c.BindJSON(&person)
// 	err = models.AddPerson(&person)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, person)
// 	}
// }
