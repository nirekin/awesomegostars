package cmd

import (
	"strings"
)

func getCategory(c string) string {
	return strings.Replace(strings.ToLower(strings.Trim(c, " ")), " ", "-", -1)
}
