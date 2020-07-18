package ini

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BlackwonderTF/go-ini/feature"
	"github.com/BlackwonderTF/go-ini/utils"
)

type File struct {
	feature.Section
}

func CreateFile() File {
	f := File{}
	return f
}

func Load(filePath string) *File {
	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.ini", currentDir, filePath))
	if err != nil {
		log.Fatal(err)
	}

	text := utils.RegSplit(string(content), "[\\n\\r]+")

	iniFile := readFile(text)

	return iniFile
}

func readFile(content []string) *File {
	iniFile := CreateFile()

	readSection(&iniFile.Section, content, 0)

	return &iniFile
}

func readSection(currentSection *feature.Section, file []string, index int) int {
	var parent *feature.Section
	if currentSection.Parent == nil {
		parent = currentSection
	} else {
		parent = currentSection.Parent
	}

	for len(file) > index {
		line := file[index]
		if isSection(line) {
			var section *feature.Section
			prefix := len(getFeaturePrefix(line))
			if prefix > len(currentSection.Prefix) {
				section = createSection(line, currentSection)
				index = readSection(section, file, index+1) - 1
			} else if prefix < len(currentSection.Prefix) {
				return index
			} else {
				section = createSection(line, parent)
				currentSection = section
			}
		} else if isProperty(line) {
			prefix := getFeaturePrefix(line)
			if currentSection.Prefix != prefix {
				return index
			}

			property := readProperty(line)
			currentSection.AddProperty(&property)
		}
		index++
	}
	return index
}

func readProperty(line string) feature.Property {
	property, err := getProperty(line)

	if err != nil {
		log.Fatal(err)
	}

	return property
}
