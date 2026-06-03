package helper

import "strings"

func StringToList(value string) []string {
	list := []string{}
	value = strings.TrimSpace(value)
	for _, item := range strings.Split(value, ",") {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		list = append(list, item)
	}
	return list
}

// SanitizeFilename replaces characters invalid in filenames
func SanitizeFilename(name string) string {
	replacer := strings.NewReplacer(
		" ", "_",
		"/", "-",
		"\\", "-",
		":", "-",
		"*", "-",
		"?", "-",
		"\"", "-",
		"<", "-",
		">", "-",
		"|", "-",
	)
	return replacer.Replace(name)
}
