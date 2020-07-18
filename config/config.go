package config

import (
	"fmt"
	"strings"
)

var commentChars = []string{";", "#"}
var quotesChars = []string{"\"", "'"}
var seperatorChars = []string{"=", ":"}
var defaultSeperator = &seperatorChars[0]

func GetQuotesRegex() string {
	return fmt.Sprintf("[%s]", strings.Join(quotesChars, ""))
}

func GetCommentsRegex() string {
	return fmt.Sprintf("[%s]", strings.Join(commentChars, ""))
}

func GetSeperatorRegex() string {
	return fmt.Sprintf("[%s]", strings.Join(seperatorChars, ""))
}

func GetQuotesChars() []string {
	return quotesChars
}

func GetDefaultSeperator() string {
	return *defaultSeperator
}

func SetDefaultSeperator(seperator string) {
	for _, sep := range seperatorChars {
		if sep == seperator {
			*defaultSeperator = sep
			return
		}
	}

	seperatorChars = append(seperatorChars, seperator)
	defaultSeperator = &seperatorChars[len(seperatorChars)-1]
}
