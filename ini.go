package ini

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/BlackwonderTF/go-ini/feature"
)

type File struct {
	Sections []feature.Section
}

func Parse(filePath string) {
	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(fmt.Sprintf("%s\\%s.ini", currentDir, filePath))

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if feature.IsComment(line) {
			log.Printf("Comment: %s\n", line)
			continue
		}

		if feature.IsSection(line) {
			// Handle section stuff
			// Check if within other section
			log.Printf("Section: %s\n", line)
			continue
		}

		if feature.IsProperty(line) {
			// Handle property stuff
			// Check if global
			property, err := feature.GetProperty(line)

			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Property: %s:%s\n", property.Key, property.Value)
			continue
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
