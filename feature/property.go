package feature

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/BlackwonderTF/go-ini/utils"
)

type Property struct {
	Key       string
	value     string
	seperator string
}

func IsProperty(line string) bool {
	// regex := regexp.MustCompile("^([a-zA-Z]+[a-zA-Z0-9]*\\s*)=(\\s*[^=]*([^=]*\\=[^=]*)*)$")

	keyRegex := "([a-zA-Z]+[a-zA-Z0-9]*\\s*)"
	separatorRegex := "([=:])"
	valueRegex := "(.+)"

	regex := regexp.MustCompile(fmt.Sprintf("^%s%s%s$", keyRegex, separatorRegex, valueRegex))
	return regex.MatchString(line)
}

func GetProperty(line string) (*Property, error) {
	if !IsProperty(line) {
		return nil, fmt.Errorf("\"%s\" is not a property", line)
	}

	property := new(Property)

	// Handle the split
	split := utils.RegSplitFirst(line, "[=:]")
	if len(split) <= 2 {
		return nil, fmt.Errorf("\"%s\" is not a valid property", line)
	}

	property.seperator = split[1]
	split[1] = strings.Join(split[2:], " ")
	split = split[:2]

	property.Key = strings.TrimSpace(split[0])

	trimmedValue := strings.TrimSpace(split[1])
	property.value = strings.TrimSuffix(strings.TrimPrefix(trimmedValue, "\""), "\"")

	return property, nil
}

func (p Property) String() string {
	return p.value
}

func (p Property) Bool() bool {
	value, err := strconv.ParseBool(p.value)
	if err != nil {
		log.Printf("Key: \"%s\" with Value: \"%s\" can not be read as a boolean\n", p.Key, p.value)
	}
	return value
}

func (p Property) Int() int {
	value, err := strconv.ParseInt(p.value, 10, 0)
	if err != nil {
		log.Printf("Key: \"%s\" with Value: \"%s\" can not be read as an integer\n", p.Key, p.value)
	}
	return int(value)
}

func (p Property) Int8() int8 {
	value, err := strconv.ParseInt(p.value, 10, 8)
	if err != nil {
		log.Printf("Key: \"%s\" with Value: \"%s\" can not be read as a 8 bit integer\n", p.Key, p.value)
	}
	return int8(value)
}

func (p Property) Int16() int16 {
	value, err := strconv.ParseInt(p.value, 10, 16)
	if err != nil {
		log.Printf("Key: \"%s\" with Value: \"%s\" can not be read as a 16 bit integer\n", p.Key, p.value)
	}
	return int16(value)
}

func (p Property) Int32() int32 {
	value, err := strconv.ParseInt(p.value, 10, 32)
	if err != nil {
		log.Printf("Key: \"%s\" with Value: \"%s\" can not be read as a 32 bit integer\n", p.Key, p.value)
	}
	return int32(value)
}

func (p Property) Int64() int64 {
	value, err := strconv.ParseInt(p.value, 10, 64)
	if err != nil {
		log.Printf("Key: \"%s\" with Value: \"%s\" can not be read as a 64 bit integer\n", p.Key, p.value)
	}
	return int64(value)
}

func (p Property) UInt() uint {
	value, err := strconv.ParseUint(p.value, 10, 0)
	if err != nil {
		log.Printf("Key: \"%s\" with Value: \"%s\" can not be read as an unsigned integer\n", p.Key, p.value)
	}
	return uint(value)
}

func (p Property) UInt8() uint8 {
	value, err := strconv.ParseUint(p.value, 10, 8)
	if err != nil {
		log.Printf("Key: \"%s\" with Value: \"%s\" can not be read as a 8 bit unsigned integer\n", p.Key, p.value)
	}
	return uint8(value)
}

func (p Property) UInt16() uint16 {
	value, err := strconv.ParseUint(p.value, 10, 16)
	if err != nil {
		log.Printf("Key: \"%s\" with Value: \"%s\" can not be read as a 16 bit unsigned integer\n", p.Key, p.value)
	}
	return uint16(value)
}

func (p Property) UInt32() uint32 {
	value, err := strconv.ParseUint(p.value, 10, 32)
	if err != nil {
		log.Printf("Key: \"%s\" with Value: \"%s\" can not be read as a 32 bit unsigned integer\n", p.Key, p.value)
	}
	return uint32(value)
}

func (p Property) UInt64() uint64 {
	value, err := strconv.ParseUint(p.value, 10, 64)
	if err != nil {
		log.Printf("Key: \"%s\" with Value: \"%s\" can not be read as a 64 bit unsigned integer\n", p.Key, p.value)
	}
	return uint64(value)
}
