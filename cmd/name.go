package cmd

import (
	"strings"
)

func getName(s string) string {
	return s[strings.Index(s, "[")+1 : strings.Index(s, "]")]
}
