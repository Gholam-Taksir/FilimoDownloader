package stream

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"FilimoDownloader-GholamTaksir/internal/helper"
)

type Playlist struct {
	Content string
	Urls    []*url.URL
}

func GetPlaylist(client helper.HttpClient, link string) Playlist {
	playlist, err := client.Get(link)
	if err != nil {
		helper.ShowErrorAndExit(fmt.Sprintf("Failed to get playlist: %v", err))
	}
	return Playlist{
		Content: cleanContent(playlist),
		Urls:    extractUrls(playlist, link),
	}
}

func cleanContent(playlist string) string {
	pattern := `\?([^"\n]*)`
	return regexp.MustCompile(pattern).ReplaceAllString(playlist, "")
}

func extractUrls(playlist string, link string) []*url.URL {
	urls := []*url.URL{}

	// FIX: check key existence before using it
	keyPattern := regexp.MustCompile(`#EXT-X-KEY[^\n]*URI="([^"]*)"`)
	keyMatches := keyPattern.FindStringSubmatch(playlist)

	if len(keyMatches) < 2 || keyMatches[1] == "" {
		helper.ShowErrorAndExit("Playlist encryption key not found")
	}
	key := keyMatches[1]
	urls = append(urls, helper.AbsoluteUrl(link, key))

	// Extract chunk URLs (non-comment lines)
	chunksPattern := regexp.MustCompile(`(?m)^([^#\s].*)`)
	for _, chunk := range chunksPattern.FindAllString(playlist, -1) {
		chunk = strings.TrimSpace(chunk)
		if chunk == "" {
			continue
		}
		urls = append(urls, helper.AbsoluteUrl(link, chunk))
	}

	if len(urls) <= 1 {
		helper.ShowErrorAndExit("Playlist has no media chunks")
	}

	return urls
}
