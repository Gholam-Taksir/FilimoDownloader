package helper

import (
	"os"
	"path/filepath"
)

func MakeDirectories(p string) {
	if p == "." || IsFileExists(p) {
		return
	}
	if err := os.MkdirAll(p, 0755); err != nil {
		panic(err)
	}
}

func CreateFile(p string) *os.File {
	MakeDirectories(filepath.Dir(p))
	file, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return file
}

func IsFileExists(p string) bool {
	_, err := os.Stat(p)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	panic(err)
}

func ReadFile(p string) string {
	data, err := os.ReadFile(p)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func WriteFile(p string, content string) {
	file := CreateFile(p)
	defer file.Close()
	if _, err := file.WriteString(content); err != nil {
		panic(err)
	}
}

func DeleteFile(p string) {
	if err := os.Remove(p); err != nil {
		panic(err)
	}
}
