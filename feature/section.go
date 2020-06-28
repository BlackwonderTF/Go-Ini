package feature

import (
	"log"
	"strings"
)

type Section struct {
	Name       string
	Prefix     string
	Parent     *Section
	properties map[string]*Property
	sections   map[string]*Section
}

func IsSection(line string) bool {
	line = strings.TrimSpace(line)
	return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
}

func CreateSection(line string, parent *Section) *Section {
	section := Section{
		Name:   strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(line), "["), "]"),
		Prefix: GetFeaturePrefix(line),
	}

	if parent != nil {
		parent.AddSection(&section)
	}

	return &section
}

func (s *Section) AddProperty(property *Property) {
	if s.properties == nil {
		s.properties = make(map[string]*Property)
	}

	key := strings.ToLower(property.Key)

	if s.properties[key] != nil {
		log.Fatalf("Section \"%s\" can not have 2 properties with the same key \"%s\"", s.Name, property.Key)
	}

	s.properties[key] = property
}

func (s Section) GetProperty(key string) *Property {
	if s.properties == nil {
		log.Fatalf("Section \"%s\" does not have any properties!", s.Name)
	}

	property := s.properties[strings.ToLower(key)]

	if property == nil {
		log.Fatalf("Property \"%s\" on section \"%s\" does not exist", key, s.Name)
	}

	return property
}

func (s Section) GetSection(name string) *Section {
	if s.sections == nil {
		log.Fatalf("Section \"%s\" does not have any sections!", s.Name)
	}

	section := s.sections[strings.ToLower(name)]
	if section == nil {
		log.Fatalf("Section with name \"%s\" does not exist!", name)
	}
	return s.sections[strings.ToLower(name)]
}

func (s *Section) AddSection(section *Section) {
	if s.sections == nil {
		s.sections = make(map[string]*Section)
	}

	key := strings.ToLower(section.Name)
	s.sections[key] = section
	section.Parent = s
}
