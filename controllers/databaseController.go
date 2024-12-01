package controllers

import (
	"CUAgain/models"
	"CUAgain/services"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

func UploadCharacterInfo(c *gin.Context) {
	var character models.CharacterAssetMap
	err := c.ShouldBindJSON(&character)
	if err != nil {
		c.Writer.WriteHeader(502)
		_, err := c.Writer.Write([]byte("server error"))
		if err != nil {
			return
		}
		return
	}
	if services.UploadCharacterInfo(character) {
		//上传成功
		_, err := c.Writer.Write([]byte("upload success"))
		if err != nil {
			return
		}
	} else {
		//401
		c.Writer.WriteHeader(401)
		_, err := c.Writer.Write([]byte("upload failed"))
		if err != nil {
			return
		}
	}
}

func DelCharacterInfo(c *gin.Context) {
	var delObject models.DelObject
	err := c.ShouldBindJSON(&delObject)
	if err != nil {
		c.Writer.WriteHeader(502)
		_, err := c.Writer.Write([]byte("server error"))
		if err != nil {
			return
		}
		return
	}
	services.DelCharacterInfo(delObject, c)
}

func UploadPage(c *gin.Context) {
	tmpl, err := template.ParseFiles("./static/html/addCharacter.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(c.Writer, models.GetConfig())
	if err != nil {
		return
	}
}
