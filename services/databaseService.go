package services

import (
	"CUAgain/dao"
	"CUAgain/models"
	"github.com/gin-gonic/gin"
)

var config = models.GetConfig()

func UploadCharacterInfo(characterAssetMap models.CharacterAssetMap) bool {
	if characterAssetMap.Password != config.CUAgain.Password {
		return false
	}
	if dao.Repeat(characterAssetMap) {
		dao.Update(characterAssetMap)
	} else {
		dao.Insert(characterAssetMap)
	}
	return true
}

func DelCharacterInfo(delObject models.DelObject, c *gin.Context) {
	if delObject.Password != config.CUAgain.Password {
		c.Writer.WriteHeader(401)
		return
	}
	if dao.Del(delObject.Id) != nil {
		c.Writer.WriteHeader(502)
		return
	}
	return
}
