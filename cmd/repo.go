package cmd

import (
	"strings"
)

func getRepo(s string) string {
	r := s[strings.Index(s, staredLineMarker)+len(staredLineMarker):]
	r = r[:strings.Index(r, ")")]
	r = strings.TrimSuffix(r, "/")
	return r
}
