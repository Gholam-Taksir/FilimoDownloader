package stream

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"FilimoDownloader-GholamTaksir/internal/helper"
)

type Builder struct {
	directory     string
	temporary     []string
	Input         string
	Output        string
	Video         string
	Audio         []string
	Subtitle      []string
	FileExtension string
	Compress      bool
}

func (b *Builder) outputFile(dir string) string {
	fileName := path.Base(dir)
	if fileName == "." {
		fileName = "output"
	}
	if !strings.HasSuffix(fileName, b.FileExtension) {
		fileName += b.FileExtension
	}
	return path.Join(b.directory, fileName)
}

func (b *Builder) buildPlaylist(dir string) {
	playlistFile := PlaylistFile(dir)
	tempOutput := b.outputFile(dir)
	b.temporary = append(b.temporary, tempOutput)

	args := []string{
		"-allowed_extensions", "ALL",
		"-protocol_whitelist", "file,crypto",
		"-i", playlistFile,
	}
	if b.Compress {
		args = append(args,
			"-c:v", "libx264", "-preset", "fast", "-crf", "23",
			"-c:a", "aac", "-b:a", "128k",
			"-movflags", "+faststart",
		)
	} else {
		args = append(args, "-c", "copy")
	}
	args = append(args, "-y", tempOutput)

	fmt.Printf("  Converting: %s\n", path.Base(tempOutput))
	runWithSpinner(args)
}

func (b *Builder) make() {
	inputIndex := 0
	inputs := []string{"-i", b.outputFile(b.Video)}
	mapping := []string{"-map", fmt.Sprintf("%d:v", inputIndex)}
	actions := []string{"-c:v", "copy", "-c:a", "copy"}
	meta := []string{}

	var outputFile string
	if b.Output == "" {
		outputFile = b.outputFile(b.Input)
	} else {
		outputFile = b.outputFile(b.Output)
	}
	if !strings.HasSuffix(outputFile, ".mp4") {
		outputFile = strings.TrimSuffix(outputFile, path.Ext(outputFile)) + ".mp4"
	}

	// اگه audio جداگانه نداره از ویدیو audio بگیر
	if len(b.Audio) == 0 {
		mapping = append(mapping, "-map", fmt.Sprintf("%d:a?", inputIndex))
	}

	// audio tracks - همه track ها با هم
	for idx, audio := range b.Audio {
		inputIndex++
		inputs = append(inputs, "-i", b.outputFile(audio))
		mapping = append(mapping, "-map", fmt.Sprintf("%d:a", inputIndex))
		meta = append(meta, buildMeta("a", idx, audio)...)
	}

	// subtitle tracks - همه زیرنویس‌ها embed بشن
	subIndex := 0
	for _, subtitle := range b.Subtitle {
		srtFile := SrtFile(subtitle)
		if !helper.IsFileExists(srtFile) {
			fmt.Printf("  Warning: subtitle file not found: %s\n", srtFile)
			continue
		}
		content := helper.ReadFile(srtFile)
		if len(strings.TrimSpace(content)) < 10 {
			fmt.Printf("  Warning: subtitle %s is empty, skipping\n", path.Base(subtitle))
			continue
		}
		inputIndex++
		inputs = append(inputs, "-i", srtFile)
		mapping = append(mapping, "-map", fmt.Sprintf("%d:s", inputIndex))
		actions = append(actions, "-c:s", "mov_text")
		meta = append(meta, buildMeta("s", subIndex, subtitle)...)
		subIndex++
		fmt.Printf("  Embedding subtitle: %s\n", path.Base(subtitle))
	}

	args := []string{}
	args = append(args, inputs...)
	args = append(args, mapping...)
	args = append(args, actions...)
	args = append(args, meta...)
	args = append(args, "-y", outputFile, "-progress", "pipe:1", "-nostats")

	fmt.Println("  Merging all tracks...")
	runWithSpinner(args)

	if !helper.IsFileExists(outputFile) {
		fmt.Println("  Merge failed, trying simple copy fallback...")
		runSimple([]string{"-i", b.outputFile(b.Video), "-c", "copy", "-y", outputFile})
	} else {
		fmt.Printf("  Output: %s\n", path.Base(outputFile))
	}
}

func (b *Builder) cleanup() {
	for _, tmp := range b.temporary {
		if helper.IsFileExists(tmp) {
			helper.DeleteFile(tmp)
		}
	}
}

func (b *Builder) Build() {
	if b.Output != "" {
		b.directory = path.Dir(b.Output)
		helper.MakeDirectories(b.directory)
	}
	if b.directory == "" {
		b.directory = b.Input
	}

	fmt.Println("  [1/3] Converting video...")
	b.buildPlaylist(b.Video)

	if len(b.Audio) > 0 {
		fmt.Printf("  [2/3] Converting %d audio track(s)...\n", len(b.Audio))
		for _, audio := range b.Audio {
			b.buildPlaylist(audio)
		}
	} else {
		fmt.Println("  [2/3] No separate audio track.")
	}

	fmt.Println("  [3/3] Merging...")
	b.make()

	fmt.Println("  Cleaning up temp files...")
	b.cleanup()
}

func runWithSpinner(args []string) {
	cmd := exec.Command(ffmpegPath(), args...)
	cmd.Stdout = nil
	cmd.Stderr = nil

	done := make(chan bool)
	go func() {
		sp := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
		i := 0
		for {
			select {
			case <-done:
				fmt.Printf("\r  Done!                        \n")
				return
			default:
				fmt.Printf("\r  %s Processing...              ", sp[i%len(sp)])
				i++
				time.Sleep(120 * time.Millisecond)
			}
		}
	}()

	cmd.Run()
	close(done)
}

func runSimple(args []string) {
	cmd := exec.Command(ffmpegPath(), args...)
	cmd.Run()
}

func buildMeta(streamType string, idx int, dir string) []string {
	lang := path.Base(dir)
	return []string{
		fmt.Sprintf("-metadata:s:%s:%d", streamType, idx),
		fmt.Sprintf("language=%s", helper.ConvertISO6391ToISO6392(lang)),
	}
}

func ffmpegPath() string {
	app, err := exec.LookPath("ffmpeg")
	if err != nil {
		fmt.Println("ERROR: FFmpeg not found! Please install FFmpeg and add it to PATH.")
		os.Exit(1)
	}
	return app
}
