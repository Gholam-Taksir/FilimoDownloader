package history

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"FilimoDownloader-GholamTaksir/internal/helper"
)

type DownloadRecord struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Quality      string    `json:"quality"`
	Format       string    `json:"format"`
	SizeMB       float64   `json:"size_mb"`
	DownloadedAt time.Time `json:"downloaded_at"`
	FilePath     string    `json:"file_path"`
}

type History struct {
	Downloads []DownloadRecord `json:"downloads"`
	FilePath  string
}

func NewHistory(basePath string) *History {
	// Store history inside data/ subdirectory
	historyFile := path.Join(basePath, "data", "download_history.json")
	return &History{
		Downloads: []DownloadRecord{},
		FilePath:  historyFile,
	}
}

func (h *History) Load() {
	if !helper.IsFileExists(h.FilePath) {
		return
	}
	data, err := os.ReadFile(h.FilePath)
	if err != nil {
		fmt.Printf("Warning: Could not load history: %v\n", err)
		return
	}
	if err := json.Unmarshal(data, &h.Downloads); err != nil {
		fmt.Printf("Warning: Could not parse history: %v\n", err)
	}
}

func (h *History) Save() {
	// Ensure directory exists
	helper.MakeDirectories(path.Dir(h.FilePath))
	data, err := json.MarshalIndent(h.Downloads, "", "  ")
	if err != nil {
		fmt.Printf("Warning: Could not save history: %v\n", err)
		return
	}
	if err := os.WriteFile(h.FilePath, data, 0644); err != nil {
		fmt.Printf("Warning: Could not write history: %v\n", err)
	}
}

func (h *History) Add(record DownloadRecord) {
	// Prepend new record
	h.Downloads = append([]DownloadRecord{record}, h.Downloads...)
	// Keep last 100
	if len(h.Downloads) > 100 {
		h.Downloads = h.Downloads[:100]
	}
	h.Save()
}

func (h *History) Print() {
	if len(h.Downloads) == 0 {
		fmt.Println("No download history found.")
		return
	}
	fmt.Println("\n==========================================")
	fmt.Println("          Download History")
	fmt.Println("==========================================")
	for i, record := range h.Downloads {
		fmt.Printf("%d. %s\n", i+1, record.Title)
		fmt.Printf("   ID: %s | Quality: %s | Format: %s | Size: %.2f MB\n",
			record.ID, record.Quality, record.Format, record.SizeMB)
		fmt.Printf("   Date: %s\n", record.DownloadedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("   File: %s\n", record.FilePath)
		fmt.Println("---")
	}
}
