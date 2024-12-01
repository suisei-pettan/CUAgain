package controllers

import (
	"CUAgain/models"
	"CUAgain/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http/httputil"
	"net/url"
)

func ProxyHandler(c *gin.Context) {
	target, err := url.Parse("https://production-hololy.hololive.tv")
	if err != nil || c.Writer.Size() == 0 {
		c.JSON(500, gin.H{"error": "Error parsing URL"})
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	c.Request.Header.Set("X-hololy-version", models.GetConfig().Hololy.VersionBypass)
	proxy.ServeHTTP(c.Writer, c.Request)
}

func Character(c *gin.Context) {
	config := models.GetConfig()

	//Read JSON based on the file name
	character := services.Character("characters")
	if len(character) == 0 {
		ProxyHandler(c)
		return
	}

	//Remove angle limit
	if config.CUAgain.RemoveAngleLimit {
		var characterObj models.Characters
		err := json.Unmarshal(character, &characterObj)
		if err != nil {
			log.Print(err)
			ProxyHandler(c)
			return
		}
		for _, character := range characterObj.Characters {
			for i := range character.Vrms {
				character.Vrms[i].Parameter = "{\"angleLimitHipDot\": 1, \"angleLimitLegDot\": 1, \"angleLimitLegCenterDot\": 1}"
			}
		}

		character, err = json.Marshal(characterObj)
	}

	_, err := c.Writer.Write(character)
	if err != nil || c.Writer.Size() == 0 {
		return
	}
}

func News(c *gin.Context) {
	news := services.Character("news")
	if len(news) == 0 {
		ProxyHandler(c)
		return
	}
	_, err := c.Writer.Write(news)
	if err != nil || c.Writer.Size() == 0 {
		return
	}
}

func FakeLogin(c *gin.Context) {
	news := services.Character("login")
	if len(news) == 0 {
		ProxyHandler(c)
		return
	}
	_, err := c.Writer.Write(news)
	if err != nil || c.Writer.Size() == 0 {
		return
	}
}

func FakeStore(c *gin.Context) {
	storeItemList := services.Character("storeItemList")
	if len(storeItemList) == 0 {
		ProxyHandler(c)
		return
	}
	_, err := c.Writer.Write(storeItemList)
	if err != nil || c.Writer.Size() == 0 {
		return
	}
}

func FakeUserInfo(c *gin.Context) {
	storeItemList := services.Character("userInfo")
	if len(storeItemList) == 0 {
		ProxyHandler(c)
		return
	}
	_, err := c.Writer.Write(storeItemList)
	if err != nil || c.Writer.Size() == 0 {
		return
	}
}

func ModifyProvision(c *gin.Context) {
	provision := services.Character("provision")
	if len(provision) == 0 {
		ProxyHandler(c)
		return
	}
	_, err := c.Writer.Write(provision)
	if err != nil || c.Writer.Size() == 0 {
		return
	}
}

func CharacterAssets(c *gin.Context) {
	resp := services.CharacterAssets(c)
	c.Header("Content-Type", "application/json")
	_, err := c.Writer.Write(resp)
	if err != nil {
		return
	}
}
