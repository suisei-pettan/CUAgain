package main

import (
	"CUAgain/dao"
	"CUAgain/models"
	"CUAgain/routes"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	dao.Db = dao.InitDb()
	route := gin.Default()
	routes.SetupRoutes(route)
	err := route.Run(":" + strconv.Itoa(models.GetConfig().CUAgain.Port))
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
