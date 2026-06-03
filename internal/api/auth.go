package api

import (
	"encoding/json"
	"fmt"

	"FilimoDownloader-GholamTaksir/internal/helper"
)

type Auth struct {
	Data AuthData `json:"data"`
}

type AuthData struct {
	User AuthUser `json:"user"`
}

type AuthUser struct {
	Profile AuthProfile `json:"selectedProfile"`
}

type AuthProfile struct {
	Name string `json:"name"`
}

func GetUserName(client helper.HttpClient) string {
	var auth Auth
	response, err := client.Get("https://api.filimo.com/api/fa/v1/web/config/uxEvent")
	if err != nil {
		helper.ShowErrorAndExit(fmt.Sprintf("Failed to get user info: %v", err))
	}
	err = json.Unmarshal([]byte(response), &auth)
	if err != nil {
		helper.ShowErrorAndExit(fmt.Sprintf("Failed to parse user info: %v", err))
	}
	return auth.Data.User.Profile.Name
}
