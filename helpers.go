package ini

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/BlackwonderTF/go-ini/config"
	"github.com/BlackwonderTF/go-ini/feature"
	"github.com/BlackwonderTF/go-ini/utils"
)

func isSection(line string) bool {
	line = strings.TrimSpace(line)
	return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
}

func isProperty(line string) bool {
	keyRegex := "([a-zA-Z]+[a-zA-Z0-9]*\\s*)"
	valueRegex := "(.+)"

	regex := regexp.MustCompile(fmt.Sprintf("^%s%s%s$", keyRegex, config.GetSeperatorRegex(), valueRegex))
	return regex.MatchString(strings.TrimSpace(line))
}

func createSection(line string, parent *feature.Section) *feature.Section {
	section := feature.Section{
		Name:   strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(line), "["), "]"),
		Prefix: getFeaturePrefix(line),
	}

	if parent != nil {
		parent.AddSection(&section)
	}

	return &section
}

func getProperty(line string) (feature.Property, error) {
	var property feature.Property
	if !isProperty(line) {
		return property, fmt.Errorf("\"%s\" is not a property", line)
	}

	// Handle the split
	split, err := utils.RegSplitFirst(line, config.GetSeperatorRegex())
	if err != nil {
		return property, err
	} else if len(split) <= 2 {
		return property, fmt.Errorf("\"%s\" is not a valid property", line)
	}

	property.SetSeperator(split[1])
	split[1] = strings.Join(split[2:], " ")
	split = split[:2]

	property.Key = strings.TrimSpace(split[0])

	trimmedValue := strings.TrimSpace(split[1])
	quotesRegex := config.GetQuotesRegex()
	commentRegex, err := utils.RegSplitFirst(trimmedValue, fmt.Sprintf("(%s?.+%s?)(%s)", quotesRegex, quotesRegex, config.GetCommentsRegex()))

	value := trimmedValue

	if commentRegex != nil {
		property.SetComment(commentRegex[1][len(commentRegex[1])-1:], strings.TrimSpace(commentRegex[2]))
		value = strings.TrimSpace(commentRegex[1][:len(commentRegex[1])-1])
	}

	property.SetValue(removeQuotes(value))

	return property, nil
}

func removeQuotes(value string) string {
	var noQuotes string

	for _, char := range config.GetQuotesChars() {
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

func getFeaturePrefix(line string) string {
	return utils.Match(line, "^[\\s]+")
}
