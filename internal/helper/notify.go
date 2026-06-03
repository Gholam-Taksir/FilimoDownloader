package helper

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
)

func Beep() {
	fmt.Print("\a")
}

// OpenFolder پوشه دانلود رو باز میکنه - با مسیر absolute
func OpenFolder(folderPath string) {
	absPath, err := filepath.Abs(folderPath)
	if err != nil {
		absPath = folderPath
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		// explorer مستقیم پوشه رو باز میکنه
		cmd = exec.Command("explorer", absPath)
	case "darwin":
		cmd = exec.Command("open", absPath)
	case "linux":
		cmd = exec.Command("xdg-open", absPath)
	default:
		return
	}
	cmd.Start() // نه Run - بذار background بمونه
}
