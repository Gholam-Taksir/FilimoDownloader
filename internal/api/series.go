package api

import (
	"encoding/json"
	"fmt"
	"strings"

	"FilimoDownloader-GholamTaksir/internal/helper"
)

// GetEpisodeWatch دریافت اطلاعات یه قسمت - با error به جای panic
func GetEpisodeWatch(client helper.HttpClient, episodeID string) (Watch, error) {
	url := fmt.Sprintf("https://api.filimo.com/api/fa/v1/movie/watch/watch/uid/%s", episodeID)
	response, err := client.Get(url)
	if err != nil {
		return Watch{}, fmt.Errorf("failed to get episode %s: %w", episodeID, err)
	}
	var watch Watch
	if err := json.Unmarshal([]byte(response), &watch); err != nil {
		return Watch{}, fmt.Errorf("failed to parse episode %s: %w", episodeID, err)
	}
	return watch, nil
}

// IsSeries چک میکنه آیا ID یه قسمت سریاله
func IsSeries(client helper.HttpClient, id string) bool {
	url := fmt.Sprintf("https://api.filimo.com/api/fa/v1/movie/watch/watch/uid/%s", id)
	response, err := client.Get(url)
	if err != nil {
		return false
	}

	var result struct {
		Data struct {
			Attributes struct {
				MovieTitle    string `json:"movie_title"`
				SeasonNumber  int    `json:"seasonNumber"`
				EpisodeNumber int    `json:"episodeNumber"`
			} `json:"attributes"`
		} `json:"data"`
	}

	if err := json.Unmarshal([]byte(response), &result); err != nil {
		return false
	}

	if result.Data.Attributes.SeasonNumber > 0 || result.Data.Attributes.EpisodeNumber > 0 {
		return true
	}

	title := result.Data.Attributes.MovieTitle
	return strings.Contains(title, "فصل") && strings.Contains(title, "قسمت")
}
