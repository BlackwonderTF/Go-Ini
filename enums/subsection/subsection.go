package subsection

import (
	"errors"
	"log"
	"strings"
)

type SubSectionType string

const (
	Indented  SubSectionType = "Indented"
	Seperated SubSectionType = "Seperated"
)

func (t SubSectionType) RequiresSeperator() bool {
	switch t {
	case Indented:
		return false
	case Seperated:
		return true
	default:
		log.Fatal("Incorrect SubSectionType found")
	}

	return true
}

func (t SubSectionType) CheckLine(line string) bool {
	line = strings.TrimSpace(line)
	switch t {
	case Indented:
		return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
	case Seperated:
		return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
	}

	return false
}

func (t SubSectionType) Validate() error {
	switch t {
	case Indented:
		return nil
	case Seperated:
		return nil
	}

	return errors.New("Invalid subsection type")
}
