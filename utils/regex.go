package utils

import (
	"regexp"
)

func RegSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:]
	return result
}

func RegSplitFirst(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	index := reg.FindStringIndex(text)
	result := make([]string, 3)

	result[0] = text[:index[0]]
	result[1] = text[index[0]:index[1]]
	result[2] = text[index[1]:]

	return result
}
