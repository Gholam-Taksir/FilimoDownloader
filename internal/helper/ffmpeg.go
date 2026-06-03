package helper

import (
	"fmt"
	"runtime"
)

func GetFFmpegThreadLimit() []string {
	cpuCores := runtime.NumCPU()
	threads := cpuCores / 2
	if threads < 1 {
		threads = 1
	}
	if threads > 4 {
		threads = 4
	}
	return []string{"-threads", fmt.Sprintf("%d", threads)}
}
