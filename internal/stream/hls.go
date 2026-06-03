package stream

import (
	"fmt"
	"regexp"
	"strings"

	"FilimoDownloader-GholamTaksir/internal/api"
	"FilimoDownloader-GholamTaksir/internal/helper"
)

type Hls struct {
	Variants []HlsVideoVariant
	Tracks   []HlsAudioTrack
}

type HlsVideoVariant struct {
	Quality    string
	Resolution string
	Link       string
}

type HlsAudioTrack struct {
	Language string
	Link     string
}

func GetHls(client helper.HttpClient, watch api.Watch) Hls {
	l := extractHlsLink(watch)
	if l == "" {
		helper.ShowErrorAndExit("HLS link not found in watch data")
	}
	hlsContent, err := client.Get(l)
	if err != nil {
		helper.ShowErrorAndExit(fmt.Sprintf("Failed to get HLS playlist: %v", err))
	}
	return Hls{
		Variants: parseVariants(hlsContent),
		Tracks:   parseTracks(hlsContent),
	}
}

func extractHlsLink(watch api.Watch) string {
	for _, list := range watch.Data.Attributes.Sources {
		for _, source := range list {
			if source.Type == "application/vnd.apple.mpegurl" {
				return source.Link
			}
		}
	}
	return ""
}

// parseVariants - FIX: more robust regex that handles various M3U8 formats
func parseVariants(hls string) []HlsVideoVariant {
	list := []HlsVideoVariant{}

	lines := strings.Split(hls, "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "#EXT-X-STREAM-INF:") {
			continue
		}

		// Extract BANDWIDTH or NAME as quality label
		quality := ""
		bwPattern := regexp.MustCompile(`BANDWIDTH=(\d+)`)
		if m := bwPattern.FindStringSubmatch(line); len(m) > 1 {
			bw := m[1]
			// Convert bandwidth to human-readable
			switch {
			case len(bw) >= 7:
				quality = "1080p"
			case len(bw) >= 6:
				quality = "720p"
			default:
				quality = "360p"
			}
		}

		// Extract RESOLUTION
		resolution := ""
		resPattern := regexp.MustCompile(`RESOLUTION=([0-9x]+)`)
		if m := resPattern.FindStringSubmatch(line); len(m) > 1 {
			resolution = m[1]
			// Derive quality from resolution height
			parts := strings.Split(resolution, "x")
			if len(parts) == 2 {
				quality = parts[1] + "p"
			}
		}

		// Next non-empty line is the URL
		link := ""
		for j := i + 1; j < len(lines); j++ {
			candidate := strings.TrimSpace(lines[j])
			if candidate != "" && !strings.HasPrefix(candidate, "#") {
				link = candidate
				break
			}
		}

		if link == "" {
			continue
		}

		list = append(list, HlsVideoVariant{
			Quality:    quality,
			Resolution: resolution,
			Link:       link,
		})
	}

	if len(list) == 0 {
		helper.ShowErrorAndExit("No video variants found in HLS playlist")
	}
	return list
}

// parseTracks - FIX: more robust regex for audio tracks
func parseTracks(hls string) []HlsAudioTrack {
	list := []HlsAudioTrack{}
	seen := make(map[string]bool)

	pattern := regexp.MustCompile(`#EXT-X-MEDIA:TYPE=AUDIO[^\n]*LANGUAGE="([^"]*)"[^\n]*URI="([^"]*)"`)
	for _, m := range pattern.FindAllStringSubmatch(hls, -1) {
		if len(m) < 3 {
			continue
		}
		lang := m[1]
		link := m[2]
		if link == "" || seen[lang] {
			continue
		}
		seen[lang] = true
		list = append(list, HlsAudioTrack{
			Language: lang,
			Link:     link,
		})
	}
	return list
}
