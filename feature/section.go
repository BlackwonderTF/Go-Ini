package feature

import (
	"strings"
)

type Section struct {
	Name        string
	values      map[string]*Property
	subsections map[string]Section
}

func IsSection(line string) bool {
	return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
}

func GetSectionName(line string) string {
	return strings.TrimSuffix(strings.TrimPrefix(line, "["), "]")
}

func CreateSection() *Section {
	return &Section{
		values:      make(map[string]*Property),
		subsections: make(map[string]Section),
	}
}

func (s *Section) AddProperty(property *Property) {
	key := strings.ToLower(property.Key)
	s.values[key] = property
}

func (s Section) GetProperty(key string) *Property {
	return s.values[strings.ToLower(key)]
}
