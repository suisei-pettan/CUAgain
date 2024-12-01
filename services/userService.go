package services

import (
	"CUAgain/dao"
	"github.com/gin-gonic/gin"
)

func UserLoginApi(c *gin.Context, password string) {
	userIp := dao.GetIp(c, config.CUAgain.GetIpMethod)
	passwordDecoded, err := dao.DecodePassword(password, config.CUAgain.RSAPrivateKeyPath)
	if err != nil {
		c.Writer.WriteHeader(502)
		return
	}
	if passwordDecoded == config.CUAgain.LoginPassword {
		c.Writer.WriteHeader(200)
		dao.AllowUser(userIp)
		c.JSON(200, gin.H{
			"data": "success",
		})
	} else {
		c.Writer.WriteHeader(403)
	}
}
