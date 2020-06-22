package feature

import "strings"

type Section struct {
	values      []string
	subsections []Section
}

func IsSection(line string) bool {
	return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
}
