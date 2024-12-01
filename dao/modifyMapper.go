package dao

import (
	"CUAgain/models"
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func InitDb() *sql.DB {
	db, err := sql.Open("sqlite3", "CUAgain.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ReadJsonFromFile(name string) ([]byte, error) {
	file, err := os.ReadFile("./json/" + name + ".json")
	if err != nil {
		log.Fatal(err)
	}
	return file, nil
}

func GetOfficialInfo(host string, path string, headers http.Header) []byte {
	target := host + path
	client := &http.Client{}
	req, _ := http.NewRequest("GET", target, nil)
	for header, values := range headers {
		for _, value := range values {
			req.Header.Add(header, value)
		}
	}
	req.Header.Set("X-hololy-version", models.GetConfig().Hololy.VersionBypass)
	resp, _ := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	//如果发生错误，重新请求
	if resp.Body == nil {
		GetOfficialInfo(host, path, headers)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte("official website error")
	}
	return body
}

func GetModifyJsonByCovered(covered string) models.CharacterAsset {
	//验证covered是否为数字
	if _, err := strconv.Atoi(covered); err != nil {
		return models.CharacterAsset{Message: "invalid input", CharacterAssets: nil}
	}
	var modifyJson models.CharacterAsset
	var dbJson string
	err := Db.QueryRow("SELECT json FROM character_assets WHERE avatar_id = $1", covered).Scan(&dbJson)
	if err != nil {
		return models.CharacterAsset{Message: "error", CharacterAssets: nil}
	}
	err = json.Unmarshal([]byte(dbJson), &modifyJson)
	return modifyJson
}

func HolostarMovetionsJson2Object() models.CharacterAsset {
	file, err := ReadJsonFromFile("holostarMovement")
	var movementObject models.CharacterAsset
	movementObject.Message = "success"
	err = json.Unmarshal(file, &movementObject.CharacterAssets)
	if err != nil {
		log.Println(err)
	}
	return movementObject
}
