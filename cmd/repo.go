package cmd

import (
	"strings"
)

func getRepo(s string) string {
	r := s[strings.Index(s, stared_line_marker)+len(stared_line_marker):]
	r = r[:strings.Index(r, ")")]
	r = strings.TrimSuffix(r, "/")
	return r
}
