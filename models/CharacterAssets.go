package models

type CharacterAsset struct {
	Message         string `json:"message"`
	CharacterAssets []struct {
		ID          int    `json:"id"`
		Category    int    `json:"category"`
		AssetNo     int    `json:"asset_no"`
		DisplayName string `json:"display_name"`
		AssetName   string `json:"asset_name"`
		Thumbnail   string `json:"thumbnail"`
		Parameter   string `json:"parameter"`
		Priority    int    `json:"priority"`
		Hash        string `json:"hash"`
	} `json:"characterAssets"`
}
