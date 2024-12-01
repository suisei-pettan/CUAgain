package dao

import (
	"CUAgain/models"
	"database/sql"
)

var Db *sql.DB

func Insert(assetMap models.CharacterAssetMap) {
	_, err := Db.Exec("INSERT INTO character_assets (id, json, avatar_id) VALUES ($1,$2,$3)", assetMap.Id, assetMap.Json, assetMap.Covered)
	if err != nil {
		return
	}
}

func Repeat(assetMap models.CharacterAssetMap) bool {
	if Db.QueryRow("SELECT avatar_id FROM character_assets WHERE avatar_id = ?", assetMap.Covered).Scan(&assetMap.Covered) != nil {
		return false
	}
	return true
}

func Update(assetMap models.CharacterAssetMap) {
	_, err := Db.Exec("UPDATE character_assets SET id = ?, json = ?  WHERE avatar_id = ?", assetMap.Id, assetMap.Json, assetMap.Covered)
	if err != nil {
		return
	}
}

func Del(id string) error {
	_, err := Db.Exec("DELETE FROM character_assets WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
