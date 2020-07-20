package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BlackwonderTF/go-ini/enums/subsection"
	"github.com/BlackwonderTF/go-ini/utils"
)

type Config struct {
	commentChars        []string
	quotesChars         []string
	seperatorChars      []string
	defaultSeperator    *string
	subsectionType      subsection.SubSectionType
	subsectionSeperator *string
}

func InitDefault() Config {
	config := Config{
		commentChars:   []string{";", "#"},
		quotesChars:    []string{"\"", "'"},
		seperatorChars: []string{"=", ":"},
		subsectionType: subsection.Indented,
	}

	config.defaultSeperator = &config.seperatorChars[0]

	return config
}

// GetQuotesRegex returns the regex format of the quotes characters
func (c Config) GetQuotesRegex() string {
	return fmt.Sprintf("[%s]", strings.Join(c.quotesChars, ""))
}

// GetCommentsRegex returns the regex format of the comment charachters
func (c Config) GetCommentsRegex() string {
	return fmt.Sprintf("[%s]", strings.Join(c.commentChars, ""))
}

// GetSeperatorRegex returns the regex format of the seperator character
func (c Config) GetSeperatorRegex() string {
	return fmt.Sprintf("[%s]", strings.Join(c.seperatorChars, ""))
}

// GetQuotesChars returns the quotes characters
func (c Config) GetQuotesChars() []string {
	return c.quotesChars
}

// GetDefaultSeperator returns the default property seperator character
func (c Config) GetDefaultSeperator() string {
	return *c.defaultSeperator
}

// SetDefaultSeperator sets the default seperator that is to be used
func (c Config) SetDefaultSeperator(seperator string) {
	for _, sep := range c.seperatorChars {
		if sep == seperator {
			*c.defaultSeperator = sep
			return
		}
	}

	c.seperatorChars = append(c.seperatorChars, seperator)
	c.defaultSeperator = &c.seperatorChars[len(c.seperatorChars)-1]
}

func (c Config) GetSubSectionType() subsection.SubSectionType {
	return c.subsectionType
}

// SetSubSectionType sets the way a subsection is defined.
// If the type requires no seperator (like indented), a nil value can be given.
func (c *Config) SetSubSectionType(t subsection.SubSectionType, seperator string) error {
	if err := t.Validate(); err != nil {
		return err
	}

	c.subsectionType = t

	if t.RequiresSeperator() && seperator == "" {
		return errors.New("No seperator specified")
	}

	c.subsectionSeperator = utils.CreateStringPointer(seperator)
	return nil
}
