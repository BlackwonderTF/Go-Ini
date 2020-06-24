package feature

import (
	"log"
	"strings"
)

type Section struct {
	Name   string
	values map[string]*Property
}

func IsSection(line string) bool {
	return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
}

func GetSectionName(line string) string {
	return strings.TrimSuffix(strings.TrimPrefix(line, "["), "]")
}

func CreateSection() *Section {
	return &Section{
		values: make(map[string]*Property),
	}
}

func (s *Section) AddProperty(property *Property) {
	key := strings.ToLower(property.Key)
	s.values[key] = property
}

func (s Section) GetProperty(key string) *Property {
	property := s.values[strings.ToLower(key)]

	if property == nil {
		log.Fatalf("Property \"%s\" on section \"%s\" does not exist", key, s.Name)
	}

	return property
}
