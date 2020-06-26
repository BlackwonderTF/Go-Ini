package feature

import (
	"fmt"
	"strings"
)

var commentChars = []string{";", "#"}
var quotesChars = []string{"\"", "'"}
var seperatorChars = []string{":", "="}

func removeQuotes(value string) string {
	var noQuotes string

	for _, char := range quotesChars {
		if strings.HasPrefix(value, char) && strings.HasSuffix(value, char) {
			noQuotes = strings.TrimSuffix(strings.TrimPrefix(value, char), char)
			break
		}
	}

	if noQuotes == "" {
		noQuotes = value
	}

	return noQuotes
}

func GetQuotesRegex() string {
	return fmt.Sprintf("[%s]", strings.Join(quotesChars, ""))
}

func GetCommentsRegex() string {
	return fmt.Sprintf("[%s]", strings.Join(commentChars, ""))
}

func GetSeperatorRegex() string {
	return fmt.Sprintf("[%s]", strings.Join(seperatorChars, ""))
}
