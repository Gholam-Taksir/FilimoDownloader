package api

import (
	"encoding/json"
	"fmt"

	"FilimoDownloader-GholamTaksir/internal/helper"
)

type Watch struct {
	Data WatchData `json:"data"`
}

type WatchData struct {
	Attributes WatchAttributes `json:"attributes"`
}

type WatchAttributes struct {
	Name      string          `json:"movie_title"`
	Sources   [][]WatchSource `json:"multiSRC"`
	Subtitles []WatchSubtitle `json:"tracks"`
}

type WatchSource struct {
	Type string `json:"type"`
	Link string `json:"src"`
}

type WatchSubtitle struct {
	Language string `json:"srclang"`
	Link     string `json:"src"`
}

func GetWatch(client helper.HttpClient, id string) Watch {
	var watch Watch
	response, err := client.Get(fmt.Sprintf("https://api.filimo.com/api/fa/v1/movie/watch/watch/uid/%s", id))
	if err != nil {
		helper.ShowErrorAndExit(fmt.Sprintf("Failed to get watch info for ID %s: %v", id, err))
	}
	err = json.Unmarshal([]byte(response), &watch)
	if err != nil {
		helper.ShowErrorAndExit(fmt.Sprintf("Failed to parse watch info: %v", err))
	}
	return watch
}

func GetMovieInfo(client helper.HttpClient, id string) (string, string, error) {
	url := fmt.Sprintf("https://api.filimo.com/api/fa/v1/movie/%s", id)
	response, err := client.Get(url)
	if err != nil {
		return "", "", err
	}

	var result struct {
		Data struct {
			Attributes struct {
				Title       string  `json:"title"`
				Year        int     `json:"year"`
				Rating      float64 `json:"rating"`
				Description string  `json:"description"`
				Genres      []struct {
					Name string `json:"name"`
				} `json:"genres"`
			} `json:"attributes"`
		} `json:"data"`
	}

	if err := json.Unmarshal([]byte(response), &result); err != nil {
		return "", "", err
	}

	info := fmt.Sprintf("\n📽  Title: %s\n📅 Year: %d\n⭐ Rating: %.1f\n🎭 Genres: ",
		result.Data.Attributes.Title,
		result.Data.Attributes.Year,
		result.Data.Attributes.Rating,
	)

	for i, g := range result.Data.Attributes.Genres {
		if i > 0 {
			info += ", "
		}
		info += g.Name
	}

	if result.Data.Attributes.Description != "" {
		desc := result.Data.Attributes.Description
		if len(desc) > 200 {
			desc = desc[:200] + "..."
		}
		info += fmt.Sprintf("\n📝 %s", desc)
	}

	return result.Data.Attributes.Title, info, nil
}
