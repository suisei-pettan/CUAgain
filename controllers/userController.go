package controllers

import (
	"CUAgain/models"
	"CUAgain/services"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
)

type userSendPassword struct {
	Password string `json:"password"`
}

func UserLogin(c *gin.Context) {
	temp, err := template.ParseFiles("./static/html/userLogin.html")
	if err != nil {
		panic(err)
	}
	err = temp.Execute(c.Writer, nil)
}

func UserLoginApi(c *gin.Context) {
	var sendPassword userSendPassword
	err := c.ShouldBindJSON(&sendPassword)
	if err != nil {
		c.Writer.WriteHeader(502)
		return
	}
	services.UserLoginApi(c, sendPassword.Password)
}

func GetRSAPublicKey(c *gin.Context) {
	rsaPublicKey, _ := os.ReadFile(models.GetConfig().CUAgain.RSAPublicKeyPath)
	_, err := c.Writer.Write(rsaPublicKey)
	if err != nil {
		return
	}
}
