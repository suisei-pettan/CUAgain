package models

type Characters struct {
	Message    string `json:"message"`
	Characters []struct {
		Id        int    `json:"id"`
		Office    int    `json:"office"`
		Priority  int    `json:"priority"`
		Name      string `json:"name"`
		Thumbnail string `json:"thumbnail"`
		Vrms      []struct {
			Id          int    `json:"id"`
			Category    int    `json:"category"`
			DisplayName string `json:"display_name"`
			AssetNo     int    `json:"asset_no"`
			AssetName   string `json:"asset_name"`
			Thumbnail   string `json:"thumbnail"`
			Parameter   string `json:"parameter"`
			Priority    int    `json:"priority"`
			Hash        string `json:"hash"`
		} `json:"vrms"`
	} `json:"characters"`
}
