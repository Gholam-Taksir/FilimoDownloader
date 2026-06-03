package stream

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"FilimoDownloader-GholamTaksir/internal/api"
	"FilimoDownloader-GholamTaksir/internal/helper"
)

func DownloadDir(name string) string {
	absDir, _ := filepath.Abs("Downloads")
	absPath, err := filepath.Abs(name)
	if err != nil {
		panic(err)
	}
	rel, _ := filepath.Rel(absDir, absPath)
	if !strings.HasPrefix(rel, "..") {
		return absPath
	}
	return path.Join(absDir, name)
}

func SubtitleDir(base string) string { return path.Join(base, "subtitle") }
func VideoDir(base string) string    { return path.Join(base, "video") }
func AudioDir(base string) string    { return path.Join(base, "audio") }
func SrtFile(dir string) string      { return path.Join(dir, "subtitle.srt") }
func PlaylistFile(dir string) string { return path.Join(dir, "playlist.m3u8") }

// DownloadSubtitle - دانلود همه زیرنویس‌های انتخاب‌شده
func DownloadSubtitle(client helper.HttpClient, subtitle api.WatchSubtitle, base string) {
	fmt.Printf("  Downloading subtitle [%s]...\n", subtitle.Language)

	dir := path.Join(SubtitleDir(base), subtitle.Language)
	srtPath := SrtFile(dir)

	content, err := client.Get(subtitle.Link)
	if err != nil {
		fmt.Printf("  Warning: subtitle [%s] failed: %v\n", subtitle.Language, err)
		return
	}

	if strings.TrimSpace(content) == "" {
		fmt.Printf("  Warning: subtitle [%s] is empty\n", subtitle.Language)
		return
	}

	srt := convertWebVTTToSRT(content)
	helper.WriteFile(srtPath, srt)
	fmt.Printf("  Subtitle saved: subtitle/%s/subtitle.srt\n", subtitle.Language)
}

// convertWebVTTToSRT تبدیل درست WEBVTT به SRT
func convertWebVTTToSRT(webvtt string) string {
	webvtt = regexp.MustCompile(`(?m)^WEBVTT.*$`).ReplaceAllString(webvtt, "")
	webvtt = regexp.MustCompile(`(?ms)^NOTE\b.*?(\n\n|$)`).ReplaceAllString(webvtt, "")

	lines := strings.Split(strings.TrimSpace(webvtt), "\n")
	var out []string
	counter := 1
	i := 0

	tsPattern := regexp.MustCompile(`(\d{2}:\d{2}:\d{2})\.(\d{3})\s*-->\s*(\d{2}:\d{2}:\d{2})\.(\d{3})`)

	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			i++
			continue
		}
		if tsPattern.MatchString(line) {
			ts := tsPattern.ReplaceAllString(line, "$1,$2 --> $3,$4")
			out = append(out, fmt.Sprintf("%d", counter), ts)
			counter++
			i++
			for i < len(lines) {
				text := strings.TrimSpace(lines[i])
				if text == "" {
					break
				}
				out = append(out, text)
				i++
			}
			out = append(out, "")
		} else {
			i++
		}
	}
	return strings.Join(out, "\n")
}

func DownloadVideo(client helper.HttpClient, variant HlsVideoVariant, base string) {
	dir := path.Join(VideoDir(base), variant.Quality)
	downloadSegments(client, variant.Link, dir)
}

func DownloadAudio(client helper.HttpClient, track HlsAudioTrack, base string) {
	dir := path.Join(AudioDir(base), track.Language)
	downloadSegments(client, track.Link, dir)
}

// downloadSegments - نمایش شماره chunk درجا + skip اگه قبلاً دانلود شده
func downloadSegments(client helper.HttpClient, link string, dir string) {
	playlist := GetPlaylist(client, link)
	helper.WriteFile(PlaylistFile(dir), playlist.Content)

	total := len(playlist.Urls)
	fmt.Printf("  Downloading %d segments...\n", total)

	for idx, chunkUrl := range playlist.Urls {
		chunkPath := path.Join(dir, path.Base(chunkUrl.Path))
		if helper.IsFileExists(chunkPath) {
			fmt.Printf("\r  [%d/%d] Already exists, skipping...     ", idx+1, total)
			continue
		}
		fmt.Printf("\r  [%d/%d] Downloading segment...            ", idx+1, total)
		err := client.DownloadFile(chunkUrl.String(), chunkPath)
		if err != nil {
			fmt.Printf("\n  Segment %d failed, retrying...\n", idx+1)
			err2 := client.DownloadFile(chunkUrl.String(), chunkPath)
			if err2 != nil {
				helper.ShowErrorAndExit(fmt.Sprintf("Segment %d/%d failed: %v", idx+1, total, err2))
			}
		}
	}
	fmt.Printf("\r  [%d/%d] All segments downloaded.          \n", total, total)
}

// CleanupMediaDirs حذف پوشه‌های video و audio بعد از build
func CleanupMediaDirs(base string) {
	videoDir := VideoDir(base)
	audioDir := AudioDir(base)

	if helper.IsFileExists(videoDir) {
		if err := os.RemoveAll(videoDir); err != nil {
			fmt.Printf("  Warning: could not remove video dir: %v\n", err)
		} else {
			fmt.Println("  Removed: video/")
		}
	}

	if helper.IsFileExists(audioDir) {
		if err := os.RemoveAll(audioDir); err != nil {
			fmt.Printf("  Warning: could not remove audio dir: %v\n", err)
		} else {
			fmt.Println("  Removed: audio/")
		}
	}
}
