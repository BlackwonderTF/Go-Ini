package feature

import "strings"

var commentChars = []string{";", "#"}

func IsComment(line string) bool {
	for _, commentChar := range commentChars {
		if strings.HasPrefix(line, commentChar) {
			return true
		}
	}
	return false
}
