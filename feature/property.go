package feature

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Property struct {
	Key   string
	Value string
}

func IsProperty(line string) bool {
	regex := regexp.MustCompile("^([a-zA-Z]+[a-zA-Z0-9]*\\s*)=(\\s*[^=]*([^=]*\\=[^=]*)*)$")
	return regex.MatchString(line)
}

func GetProperty(line string) (*Property, error) {
	if !IsProperty(line) {
		return nil, fmt.Errorf("\"%s\" is not a property", line)
	}

	split := strings.Split(line, "=")
	if len(split) <= 1 {
		// TODO improve this
		return nil, errors.New("Property is not valid error")
	} else if len(split) > 2 {
		for i, arg := range split {
			if strings.HasSuffix(arg, "\\") {
				split[i] = arg[:len(arg)-1]
			}
		}
		split[1] = strings.Join(split[1:], " ")
	}

	property := new(Property)

	property.Key = strings.TrimSpace(split[0])

	trimmedValue := strings.TrimSpace(split[1])
	property.Value = strings.TrimSuffix(strings.TrimPrefix(trimmedValue, "\""), "\"")

	return property, nil
}
