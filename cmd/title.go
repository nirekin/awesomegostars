package cmd

import (
	"strings"
)

func getTitle(s string) string {
	return strings.ToLower(strings.Replace(strings.TrimSpace(strings.Replace(s, title_marker, "", -1)), " ", "-", -1))
}
