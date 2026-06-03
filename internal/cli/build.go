package cli

import (
	"fmt"
	"os"
	"path"
	"strings"

	"FilimoDownloader-GholamTaksir/internal/helper"
	"FilimoDownloader-GholamTaksir/internal/stream"
)

type BuildOption struct {
	Name string
	Path string
}

func Build(input string, output string) {
	fmt.Println("\n==========================================")
	fmt.Println("  Build Mode")
	fmt.Println("==========================================")

	selectedFormat := helper.QuestionSingle("Output format:", []string{"MP4", "MKV"})
	ext := ".mp4"
	if selectedFormat == 1 {
		ext = ".mkv"
	}
	fmt.Printf("  Format: %s\n", strings.ToUpper(ext[1:]))

	builder := stream.Builder{
		Input:         input,
		Output:        output,
		FileExtension: ext,
		Compress:      true,
	}

	videoDir := stream.VideoDir(input)
	audioDir := stream.AudioDir(input)
	subtitleDir := stream.SubtitleDir(input)

	if !helper.IsFileExists(videoDir) {
		helper.ShowErrorAndExit(fmt.Sprintf("Video directory not found: %s", videoDir))
	}

	videoOptions := getOptions(videoDir, true)
	if len(videoOptions) == 0 {
		helper.ShowErrorAndExit("No video options found")
	}

	var audioOptions, subtitleOptions []BuildOption
	if helper.IsFileExists(audioDir) {
		audioOptions = getOptions(audioDir, true)
	}
	if helper.IsFileExists(subtitleDir) {
		subtitleOptions = getOptions(subtitleDir, false)
	}

	// انتخاب ویدیو - دقیقاً یکی
	sel := ask(VIDEO_STEP, videoOptions)
	if len(sel) != 1 {
		helper.ShowErrorAndExit("Select exactly one video quality")
	}
	builder.Video = getPaths(sel, videoOptions)[0]

	// انتخاب صدا - میشه چند تا انتخاب کرد
	if len(audioOptions) > 0 {
		fmt.Println("  (You can select multiple audio tracks, e.g: 1,2)")
		builder.Audio = getPaths(ask(AUDIO_STEP, audioOptions), audioOptions)
	}

	// انتخاب زیرنویس - میشه چند تا انتخاب کرد
	if len(subtitleOptions) > 0 {
		fmt.Println("  (Subtitles will be embedded in the video AND kept as .srt files)")
		fmt.Println("  (You can select multiple subtitles, e.g: 1,2)")
		builder.Subtitle = getPaths(ask(SUBTITLE_STEP, subtitleOptions), subtitleOptions)
	}

	builder.Build()

	// حذف پوشه‌های video و audio بعد از build
	fmt.Println("\n  Removing raw media directories...")
	stream.CleanupMediaDirs(input)

	fmt.Println("\n==========================================")
	fmt.Println("  Build complete!")
	fmt.Println("==========================================")
	helper.WaitForEnter()
}

func getOptions(base string, isPlaylist bool) []BuildOption {
	options := []BuildOption{}
	entries, err := os.ReadDir(base)
	if err != nil {
		helper.ShowErrorAndExit(fmt.Sprintf("Cannot read %s: %v", base, err))
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		dirPath := path.Join(base, entry.Name())
		var checkFile string
		if isPlaylist {
			checkFile = stream.PlaylistFile(dirPath)
		} else {
			checkFile = stream.SrtFile(dirPath)
		}
		if !helper.IsFileExists(checkFile) {
			continue
		}
		options = append(options, BuildOption{Name: entry.Name(), Path: dirPath})
	}
	return options
}

func getPaths(selected []int, options []BuildOption) []string {
	paths := []string{}
	for _, idx := range selected {
		paths = append(paths, options[idx].Path)
	}
	return paths
}

func ask(label string, options []BuildOption) []int {
	items := []string{}
	for _, o := range options {
		items = append(items, o.Name)
	}
	return helper.Question(label, items)
}
