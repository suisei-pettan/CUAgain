package routes

import (
	"CUAgain/controllers"
	"CUAgain/interceptor"
	"CUAgain/models"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine) {
	println("Setting up routes...")
	//api修改分组
	modifyApi := route.Group("/api", interceptor.CheckTheSignInStatus)
	{
		modifyApi.GET("/characters", controllers.Character)
		modifyApi.GET("/news", controllers.News)
		modifyApi.GET("/characterAssets", controllers.CharacterAssets)
		modifyApi.GET("/login", controllers.FakeLogin)
		modifyApi.GET("/user", controllers.FakeLogin)
		modifyApi.GET("/storeItemList", controllers.FakeStore)
		modifyApi.GET("/user/info", controllers.FakeUserInfo)
		//modifyApi.GET("/login", controllers.FakeLogin)
	}
	//数据库上传分组
	databaseUpload := route.Group("/upload")
	{
		databaseUpload.POST("/characterInfo", controllers.UploadCharacterInfo)
		databaseUpload.DELETE("/characterInfo", controllers.DelCharacterInfo)
		databaseUpload.GET("/page", controllers.UploadPage)
	}
	//用户登录验证
	if models.GetConfig().CUAgain.LoginAuth {
		userManage := route.Group("/user")
		{
			userApi := userManage.Group("/api")
			{
				userApi.GET("/public-key", controllers.GetRSAPublicKey)
				userApi.POST("/login", controllers.UserLoginApi)
			}
			userManage.GET("/login", controllers.UserLogin)
		}
	}

	//资源文件代理
	route.Any("/asset/*filepath", interceptor.CheckTheSignInStatus, controllers.AssetProxyHandler)

	//剩余的走代理
	route.NoRoute(controllers.ProxyHandler)

}
