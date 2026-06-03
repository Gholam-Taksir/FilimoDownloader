package cli

import (
	"fmt"
	"path"
	"strings"

	"FilimoDownloader-GholamTaksir/internal/api"
	"FilimoDownloader-GholamTaksir/internal/helper"
	"FilimoDownloader-GholamTaksir/internal/stream"
)

func Download(app App, args Args) {
	if args.Version {
		app.PrintVersion()
		return
	}
	if args.Author {
		app.PrintAuthor()
		return
	}
	if args.Help {
		args.PrintHelp()
		return
	}

	printBanner(app)

	// ─── Auth ───
	authToken := helper.AuthToken{
		Path:      helper.TokenPath(app.BasePath),
		FlagValue: args.Token,
	}
	token := authToken.Get()
	if token == "" {
		fmt.Println("No auth token found.")
		token = helper.Input("Enter your Filimo auth token:")
		if token == "" {
			helper.ShowErrorAndExit("No token provided!")
		}
		authToken.Set(token)
	}

	httpClient := helper.HttpClient{
		Token:     token,
		UserAgent: helper.GetUserAgent(),
	}

	username := api.GetUserName(httpClient)
	if username == "" {
		authToken.Delete()
		helper.ShowErrorAndExit("Invalid auth token! Token deleted.")
	}
	fmt.Printf("Logged in as: %s\n\n", username)

	// ─── اول نوع محتوا ───
	fmt.Println("==========================================")
	fmt.Println("  What do you want to download?")
	fmt.Println("==========================================")
	contentType := helper.QuestionSingle("Content type:", []string{
		"Movie",
		"Series - Single episode",
	})

	// ─── بعد ID ───
	id := args.Id
	if id == "" {
		switch contentType {
		case 0:
			id = helper.Input("\nEnter Movie ID or URL:")
		case 1:
			id = helper.Input("\nEnter Episode ID or URL:")
		}
		if id == "" {
			helper.ShowErrorAndExit("No ID provided!")
		}
	}
	if strings.HasPrefix(id, "https://") {
		u := helper.NewUrl(id)
		id = path.Base(u.Path)
	}

	switch contentType {
	case 0:
		downloadMovie(app, httpClient, id, args)
	case 1:
		downloadSingleEpisode(app, httpClient, id, args)
	}
}

// ─────────────────────────────────────────────
//  MOVIE
// ─────────────────────────────────────────────

func downloadMovie(app App, client helper.HttpClient, id string, args Args) {
	fmt.Println("\n==========================================")
	fmt.Println("  Movie Mode")
	fmt.Println("==========================================")

	if app.Config.ShowInfoBeforeDL {
		_, info, err := api.GetMovieInfo(client, id)
		if err == nil {
			fmt.Println(info)
		}
	}

	watch := api.GetWatch(client, id)
	contentName := helper.SanitizeFilename(watch.Data.Attributes.Name)
	fmt.Printf("Title: %s\n", watch.Data.Attributes.Name)

	downloadDir := resolveDir(args.Output, contentName)
	downloadAndBuild(app, client, watch, contentName, downloadDir)
}

// ─────────────────────────────────────────────
//  SERIES - SINGLE EPISODE
// ─────────────────────────────────────────────

func downloadSingleEpisode(app App, client helper.HttpClient, episodeId string, args Args) {
	fmt.Println("\n==========================================")
	fmt.Println("  Series - Single Episode")
	fmt.Println("==========================================")

	watch, err := api.GetEpisodeWatch(client, episodeId)
	if err != nil {
		helper.ShowErrorAndExit(fmt.Sprintf("Failed to get episode: %v", err))
	}

	contentName := helper.SanitizeFilename(watch.Data.Attributes.Name)
	fmt.Printf("Episode: %s\n", watch.Data.Attributes.Name)

	downloadDir := resolveDir(args.Output, contentName)
	downloadAndBuild(app, client, watch, contentName, downloadDir)
}

// ─────────────────────────────────────────────
//  DOWNLOAD + BUILD
// ─────────────────────────────────────────────

func downloadAndBuild(app App, client helper.HttpClient, watch api.Watch, contentName string, downloadDir string) {
	hls := stream.GetHls(client, watch)

	fmt.Println("\n==========================================")
	fmt.Printf("  Content: %s\n", contentName)
	fmt.Println("==========================================")

	// انتخاب ویدیو
	var selVariants []int
	if len(hls.Variants) > 0 {
		opts := []string{}
		for _, v := range hls.Variants {
			opts = append(opts, fmt.Sprintf("%s  [%s]", v.Quality, v.Resolution))
		}
		selVariants = helper.Question(VIDEO_STEP, opts)
	}

	// انتخاب صدا - چند تا همزمان
	var selTracks []int
	if len(hls.Tracks) > 0 {
		opts := []string{}
		for _, t := range hls.Tracks {
			opts = append(opts, t.Language)
		}
		fmt.Println("  (Select multiple audio tracks with comma, e.g: 1,2)")
		selTracks = helper.Question(AUDIO_STEP, opts)
	}

	// انتخاب زیرنویس - چند تا همزمان
	var selSubs []int
	if len(watch.Data.Attributes.Subtitles) > 0 {
		opts := []string{}
		for _, s := range watch.Data.Attributes.Subtitles {
			opts = append(opts, s.Language)
		}
		fmt.Println("  (Select multiple subtitles with comma, e.g: 1,2)")
		selSubs = helper.Question(SUBTITLE_STEP, opts)
	}

	fmt.Println("\n==========================================")
	fmt.Println("  Downloading...")
	fmt.Println("==========================================")

	// دانلود همه زیرنویس‌های انتخاب شده
	for _, item := range selSubs {
		stream.DownloadSubtitle(client, watch.Data.Attributes.Subtitles[item], downloadDir)
	}

	// دانلود همه کیفیت‌های ویدیو انتخاب شده
	for _, item := range selVariants {
		fmt.Printf("\nDownloading video [%s]...\n", hls.Variants[item].Quality)
		stream.DownloadVideo(client, hls.Variants[item], downloadDir)
	}

	// دانلود همه track های صوتی انتخاب شده
	for _, item := range selTracks {
		fmt.Printf("\nDownloading audio [%s]...\n", hls.Tracks[item].Language)
		stream.DownloadAudio(client, hls.Tracks[item], downloadDir)
	}

	fmt.Println("\n==========================================")
	fmt.Println("  Download complete! Building final file...")
	fmt.Println("==========================================")
	helper.Beep()

	Build(downloadDir, "")

	// باز کردن پوشه دانلود
	if app.Config.AutoOpenFolder {
		fmt.Printf("\nOpening folder: %s\n", downloadDir)
		helper.OpenFolder(downloadDir)
	}
}

// ─────────────────────────────────────────────
//  HELPERS
// ─────────────────────────────────────────────

func resolveDir(outputArg string, contentName string) string {
	if outputArg != "" {
		return stream.DownloadDir(outputArg)
	}
	return stream.DownloadDir(contentName)
}

func printBanner(app App) {
	fmt.Println("\n==========================================")
	fmt.Printf("  %s %s\n", app.Name, app.Version)
	fmt.Printf("  by %s\n", app.Author)
	fmt.Println("==========================================\n")
}
