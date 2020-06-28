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
	Key         string
	value       string
	seperator   string
	comment     *string
	commentChar *string
}

func IsProperty(line string) bool {
	keyRegex := "([a-zA-Z]+[a-zA-Z0-9]*\\s*)"
	valueRegex := "(.+)"

	regex := regexp.MustCompile(fmt.Sprintf("^%s%s%s$", keyRegex, GetSeperatorRegex(), valueRegex))
	return regex.MatchString(strings.TrimSpace(line))
}

func GetProperty(line string) (Property, error) {
	var property Property
	if !IsProperty(line) {
		return property, fmt.Errorf("\"%s\" is not a property", line)
	}

	// Handle the split
	split, err := utils.RegSplitFirst(line, GetSeperatorRegex())
	if err != nil {
		return property, err
	} else if len(split) <= 2 {
		return property, fmt.Errorf("\"%s\" is not a valid property", line)
	}

	property.seperator = split[1]
	split[1] = strings.Join(split[2:], " ")
	split = split[:2]

	property.Key = strings.TrimSpace(split[0])

	trimmedValue := strings.TrimSpace(split[1])
	quotesRegex := GetQuotesRegex()
	commentRegex, err := utils.RegSplitFirst(trimmedValue, fmt.Sprintf("(%s?.+%s?)(%s)", quotesRegex, quotesRegex, GetCommentsRegex()))

	value := trimmedValue

	if commentRegex != nil {
		property.comment = utils.CreateStringPointer(strings.TrimSpace(commentRegex[2]))
		property.commentChar = utils.CreateStringPointer(commentRegex[1][len(commentRegex[1])-1:])
		value = strings.TrimSpace(commentRegex[1][:len(commentRegex[1])-1])
	}

	property.value = removeQuotes(value)

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
