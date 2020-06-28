package utils

import (
	"errors"
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

func RegSplitFirst(text string, delimeter string) ([]string, error) {
	reg := regexp.MustCompile(delimeter)
	index := reg.FindStringIndex(text)

	if index == nil {
		return nil, errors.New("No entries found")
	}

	result := make([]string, 3)

	result[0] = text[:index[0]]
	result[1] = text[index[0]:index[1]]
	result[2] = text[index[1]:]

	return result, nil
}

func Match(text string, delimeter string) string {
	reg := regexp.MustCompile(delimeter)
	return reg.FindString(text)
}
