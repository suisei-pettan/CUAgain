package services

import (
	"CUAgain/dao"
	"CUAgain/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func Character(name string) []byte {
	info, err := dao.ReadJsonFromFile(name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return info
}

func CharacterAssets(c *gin.Context) []byte {
	var avatarId string
	if c.Request.Header.Get("X-hololy-avatar-id") == "" {
		avatarId = "Default"
	} else {
		avatarId = c.Request.Header.Get("X-hololy-avatar-id")
	}

	// First try to get modified JSON
	modifyObject := dao.GetModifyJsonByCovered(avatarId)
	if modifyObject.CharacterAssets != nil {
		var modifyBody []byte
		var err error
		if config.CUAgain.EnableGlobalHolostarMovement {
			modifyObject.CharacterAssets = append(modifyObject.CharacterAssets, dao.HolostarMovetionsJson2Object().CharacterAssets...)
			modifyBody, err = json.Marshal(modifyObject)
			if err != nil {
				log.Fatal(err)
			}
		}
		return modifyBody
	}

	// If modified JSON not found, fetch from official API
	officialBody := dao.GetOfficialInfo("https://production-hololy.hololive.tv", c.Request.URL.Path, c.Request.Header)
	officialBodyOri := officialBody

	var officialJson models.CharacterAsset
	if err := json.Unmarshal(officialBody, &officialJson); err != nil {
		log.Fatal(err)
	}

	if config.CUAgain.EnableGlobalHolostarMovement {
		officialJson.CharacterAssets = append(officialJson.CharacterAssets, dao.HolostarMovetionsJson2Object().CharacterAssets...)
		var err error
		officialBody, err = json.Marshal(officialJson)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Store the official data if there are assets
	if len(officialJson.CharacterAssets) > 0 {
		characterAssetMap := models.CharacterAssetMap{
			Id:       officialJson.CharacterAssets[0].AssetName,
			Json:     string(officialBodyOri),
			Covered:  avatarId,
			Password: config.CUAgain.Password,
		}
		UploadCharacterInfo(characterAssetMap)
	}

	return officialBody
}
