package ini

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BlackwonderTF/go-ini/feature"
)

type File struct {
	items    []feature.Feature
	sections map[string]*feature.Section
	globals  map[string]*feature.Property
}

func createFile() File {
	return File{
		items:    make([]feature.Feature, 0),
		sections: make(map[string]*feature.Section),
		globals:  make(map[string]*feature.Property),
	}
}

func (f File) Section(section string) feature.Section {
	return *f.sections[strings.ToLower(section)]
}

func Load(filePath string) *File {
	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(fmt.Sprintf("%s/%s.ini", currentDir, filePath))

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	iniFile := readFile(scanner, createFile())

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return iniFile
}

func readFile(scanner *bufio.Scanner, iniFile File) *File {
	var currentSection *feature.Section
	for scanner.Scan() {
		line := scanner.Text()

		if feature.IsSection(line) {
			currentSection = feature.CreateSection()
			currentSection.Name = feature.GetSectionName(line)

			iniFile.items = append(iniFile.items, currentSection)
			iniFile.sections[strings.ToLower(currentSection.Name)] = currentSection
		} else if feature.IsProperty(line) {
			property, err := feature.GetProperty(line)

			if err != nil {
				log.Fatal(err)
			}

			iniFile.items = append(iniFile.items, property)
			if currentSection == nil {
				iniFile.globals[strings.ToLower(property.Key)] = property
			} else {
				currentSection.AddProperty(property)
			}
		}
	}

	return &iniFile
}
