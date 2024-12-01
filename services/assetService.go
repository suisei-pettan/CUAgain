package services

import (
	"CUAgain/dao"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func AssetProxyReadCache(c *gin.Context) bool {
	cache := dao.TryGetAssetCache(c.Request.URL.Path)
	if cache != nil {
		c.Writer.WriteHeader(200)
		c.Writer.Header().Set("Content-Length", strconv.Itoa(len(cache)))
		//如果资产不为png
		if c.Request.URL.Path[len(c.Request.URL.Path)-3:] != "png" {
			c.Writer.Header().Set("Content-Type", "binary/octet-stream")
		}
		_, err := c.Writer.Write(cache)
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	} else {
		return false
	}
}

func GetAssetWithCacheFromOfficial(c *gin.Context) []byte {
	asset := dao.GetOfficialInfo("https://production-assetbundle-hololy.hololive.tv", c.Request.URL.Path, c.Request.Header)
	if asset != nil {
		err := dao.WriteAssetCache(c.Request.URL.Path, asset)
		if err != nil {
			return nil
		}
		log.Println(c.Request.URL.Path + " record in cache")
	}
	return asset
}
