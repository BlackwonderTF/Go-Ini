package ini

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/shellucas/go-ini/config"
	"github.com/shellucas/go-ini/enums/subsection"
	"github.com/shellucas/go-ini/feature"
	"github.com/shellucas/go-ini/utils"
)

var Files map[string]*file = make(map[string]*file)

type file struct {
	feature.Section
	Config config.Config
}

func CreateFile() file {
	f := file{
		Config: config.InitDefault(),
	}
	return f
}

func Load(filePath string, file file) *file {
	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	fileName := strings.TrimSuffix(filePath, ".ini")

	content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.ini", currentDir, fileName))
	if err != nil {
		log.Fatal(err)
	}

	text := utils.RegSplit(string(content), "[\\n\\r]+")

	iniFile := readFile(text, file)

	iniFile.Name = fileName
	Files[fileName] = iniFile

	return iniFile
}

func readFile(content []string, iniFile file) *file {
	switch iniFile.Config.GetSubSectionType() {
	case subsection.Indented:
		readIndentedSection(&iniFile.Section, content, 0, iniFile.Config)
		break
	case subsection.Seperated:
		readSeperatedSection(&iniFile.Section, content, iniFile.Config)
		break
	}

	return &iniFile
}

func readProperty(line string, config config.Config) feature.Property {
	property, err := getProperty(line, config)

	if err != nil {
		log.Fatal(err)
	}

	return property
}

func readIndentedSection(currentSection *feature.Section, file []string, index int, config config.Config) int {
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
			if checkSectionLineDepth(&line, currentSection, config.GetSubSectionType()) > 0 {
				section = createSection(line, currentSection, config.GetSubSectionType())
				index = readIndentedSection(section, file, index+1, config) - 1
			} else if checkSectionLineDepth(&line, currentSection, config.GetSubSectionType()) < 0 {
				return index
			} else {
				section = createSection(line, parent, config.GetSubSectionType())
				currentSection = section
			}
		} else if isProperty(line, config) {
			prefix := getFeaturePrefix(line, config.GetSubSectionType())
			if currentSection.Prefix != prefix {
				return index
			}

			property := readProperty(line, config)
			currentSection.AddProperty(&property)
		}
		index++
	}
	return index
}

func readSeperatedSection(parent *feature.Section, file []string, config config.Config) {
	var currentSection *feature.Section
	currentSection = parent

	for i := 0; i < len(file); i++ {
		line := file[i]
		if isSection(line) {
			parents := utils.RegSplit(line, `\.`)
			currentSection = parent
			for _, x := range parents[:len(parents)-1] {
				currentSection = currentSection.GetSection(strings.TrimPrefix(x, "["))
			}

			currentSection = createSection(line, currentSection, config.GetSubSectionType())
		} else if isProperty(line, config) {
			property := readProperty(line, config)
			currentSection.AddProperty(&property)
		}
	}
}
