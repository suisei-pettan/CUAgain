package controllers

import (
	"CUAgain/dao"
	"CUAgain/models"
	"CUAgain/services"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func AssetProxyHandler(c *gin.Context) {
	// 检查路径是否包含 "Provisioning"
	if strings.Contains(c.Request.URL.Path, "Provisioning") {
		ModifyProvision(c)
		return
	}

	// 修改请求头
	c.Request.Header.Set("Host", "production-assetbundle-hololy.hololive.tv")

	config := models.GetConfig()
	// 检查资源反向代理是否开启
	if !config.CUAgain.AssetsProxy {
		c.JSON(403, gin.H{"error": "Assets proxy is disabled"})
		return
	}

	// 检查资源反向代理缓存是否开启
	if config.CUAgain.AssetsCache {
		if services.AssetProxyReadCache(c) {
			return
		} else {
			asset := services.GetAssetWithCacheFromOfficial(c)
			if asset != nil {
				c.Writer.WriteHeader(200)
				c.Writer.Header().Set("Content-Length", strconv.Itoa(len(asset)))
				// 如果资产不为 png
				if c.Request.URL.Path[len(c.Request.URL.Path)-3:] != "png" {
					c.Writer.Header().Set("Content-Type", "binary/octet-stream")
				}
				_, err := c.Writer.Write(asset)
				if err != nil {
					return
				}
				return
			}
		}
	}

	asset := dao.GetOfficialInfo("https://production-assetbundle-hololy.hololive.tv", c.Request.URL.Path, c.Request.Header)
	_, err := c.Writer.Write(asset)
	if err != nil {
		return
	}
}
