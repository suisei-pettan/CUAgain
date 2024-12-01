package models

type CharacterAssetMap struct {
	Id       string `json:"id"`
	Json     string `json:"json"`
	Covered  string `json:"covered"`
	Password string `json:"password"`
}
type DelObject struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}
