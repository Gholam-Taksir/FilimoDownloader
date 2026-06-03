package config

import (
	"encoding/json"
	"os"
	"path"

	"FilimoDownloader-GholamTaksir/internal/helper"
)

type Settings struct {
	DefaultQuality   string `json:"default_quality"`
	DefaultFormat    string `json:"default_format"`
	DownloadPath     string `json:"download_path"`
	MaxThreads       int    `json:"max_threads"`
	AutoOpenFolder   bool   `json:"auto_open_folder"`
	ShowInfoBeforeDL bool   `json:"show_info_before_dl"`
}

var DefaultSettings = Settings{
	DefaultQuality:   "720p",
	DefaultFormat:    "mp4",
	DownloadPath:     "Downloads",
	MaxThreads:       2,
	AutoOpenFolder:   true,
	ShowInfoBeforeDL: true,
}

// DataDir returns the data subdirectory path
func DataDir(basePath string) string {
	return path.Join(basePath, "data")
}

func GetConfigPath(basePath string) string {
	return path.Join(DataDir(basePath), "config.json")
}

func Load(basePath string) *Settings {
	// Ensure data directory exists
	helper.MakeDirectories(DataDir(basePath))

	configPath := GetConfigPath(basePath)
	if !helper.IsFileExists(configPath) {
		Save(basePath, &DefaultSettings)
		return &DefaultSettings
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return &DefaultSettings
	}

	var settings Settings
	if err := json.Unmarshal(data, &settings); err != nil {
		return &DefaultSettings
	}

	return &settings
}

func Save(basePath string, settings *Settings) {
	helper.MakeDirectories(DataDir(basePath))
	configPath := GetConfigPath(basePath)
	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return
	}
	helper.WriteFile(configPath, string(data))
}
