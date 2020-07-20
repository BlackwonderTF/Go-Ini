package ini

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/shellucas/go-ini/config"
	"github.com/shellucas/go-ini/enums/subsection"
	"github.com/shellucas/go-ini/feature"
	"github.com/shellucas/go-ini/utils"
)

func isSection(line string) bool {
	line = strings.TrimSpace(line)
	return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
}

func isProperty(line string, config config.Config) bool {
	keyRegex := "([a-zA-Z]+[a-zA-Z0-9]*\\s*)"
	valueRegex := "(.+)"

	regex := regexp.MustCompile(fmt.Sprintf("^%s%s%s$", keyRegex, config.GetSeperatorRegex(), valueRegex))
	return regex.MatchString(strings.TrimSpace(line))
}

func createSection(line string, parent *feature.Section, subType subsection.SubSectionType) *feature.Section {
	prefix := getFeaturePrefix(line, subType)
	section := feature.Section{
		Name:   strings.TrimSuffix(strings.TrimPrefix(line[len(prefix)+1:], "["), "]"),
		Prefix: prefix,
	}

	if parent != nil {
		parent.AddSection(&section)
	}

	return &section
}

func getProperty(line string, config config.Config) (feature.Property, error) {
	var property feature.Property
	if !isProperty(line, config) {
		return property, fmt.Errorf("\"%s\" is not a property", line)
	}

	// Handle the split
	split, err := utils.RegSplitFirst(line, config.GetSeperatorRegex())
	if err != nil {
		return property, err
	} else if len(split) <= 2 {
		return property, fmt.Errorf("\"%s\" is not a valid property", line)
	}

	seperator := split[1]
	split[1] = strings.Join(split[2:], " ")
	split = split[:2]

	trimmedValue := strings.TrimSpace(split[1])
	quotesRegex := config.GetQuotesRegex()
	commentRegex, err := utils.RegSplitFirst(trimmedValue, fmt.Sprintf("(%s?.+%s?)(%s)", quotesRegex, quotesRegex, config.GetCommentsRegex()))

	key := strings.TrimSpace(split[0])
	value := trimmedValue

	if commentRegex != nil {
		property.SetComment(commentRegex[1][len(commentRegex[1])-1:], strings.TrimSpace(commentRegex[2]))
		value = strings.TrimSpace(commentRegex[1][:len(commentRegex[1])-1])
	}

	value = removeQuotes(value, config)

	property = feature.CreateProperty(key, value, seperator, config)

	return property, nil
}

func removeQuotes(value string, config config.Config) string {
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

func getFeaturePrefix(line string, subType subsection.SubSectionType) string {
	switch subType {
	case subsection.Indented:
		return utils.Match(line, `^[\s]+`)
	case subsection.Seperated:
		return utils.Match(line, `[^\[].*\.`)
	}

	log.Fatal("No subsection config specified")
	return ""
}

func checkSectionLineDepth(line *string, currentSection *feature.Section, subType subsection.SubSectionType) int {
	prefixLen := getPrefixLen(getFeaturePrefix(*line, subType), subType)
	currentSectionPrefixLen := getPrefixLen(getFeaturePrefix(currentSection.Prefix, subType), subType)
	if prefixLen > currentSectionPrefixLen {
		return 1
	}
	if prefixLen == currentSectionPrefixLen {
		return 0
	}
	if prefixLen < currentSectionPrefixLen {
		return -1
	}

	log.Fatal("No subsection config specified")
	return 0
}

func getPrefixLen(prefix string, subType subsection.SubSectionType) int {
	switch subType {
	case subsection.Indented:
		return len(prefix)
	case subsection.Seperated:
		regex := regexp.MustCompile(`\.`)
		return len(regex.FindAllStringIndex(prefix, -1))
	}

	return 0
}
