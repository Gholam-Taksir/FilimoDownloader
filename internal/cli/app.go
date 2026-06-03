package cli

import (
	"fmt"

	"FilimoDownloader-GholamTaksir/internal/config"
	"FilimoDownloader-GholamTaksir/internal/helper"
	"FilimoDownloader-GholamTaksir/internal/history"
)

type App struct {
	Name     string
	Author   string
	Version  string
	BasePath string
	History  *history.History
	Config   *config.Settings
}

func (app App) PrintVersion() {
	fmt.Println(app.Name, app.Version)
}

func (app App) PrintAuthor() {
	fmt.Println("Author:", app.Author)
}

func NewApp(isProduction bool) App {
	var basePath string
	if isProduction {
		basePath = helper.ProductionBasePath()
	} else {
		basePath = helper.DebugBasePath()
	}

	// Ensure data directory exists on startup
	helper.MakeDirectories(helper.DataDirPath(basePath))

	hist := history.NewHistory(basePath)
	hist.Load()
	cfg := config.Load(basePath)

	return App{
		Name:     "FilimoDownloader",
		Author:   "GholamTaksir",
		Version:  "v2.0.0",
		BasePath: basePath,
		History:  hist,
		Config:   cfg,
	}
}
