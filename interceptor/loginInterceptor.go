package interceptor

import (
	"CUAgain/dao"
	"CUAgain/models"
	"github.com/gin-gonic/gin"
)

func CheckTheSignInStatus(c *gin.Context) {
	config := models.GetConfig()
	if !config.CUAgain.LoginAuth {
		c.Next()
		return
	}
	ip := dao.GetIp(c, config.CUAgain.GetIpMethod)
	if dao.CheckTheSignInStatus(ip) {
		c.Next()
	} else {
		c.Redirect(302, "/user/login")
		c.Abort()
		return
	}
}
