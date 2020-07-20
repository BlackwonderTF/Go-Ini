package main

import (
	"log"

	"github.com/BlackwonderTF/go-ini"
	"github.com/BlackwonderTF/go-ini/enums/subsection"
)

func main() {
	log.Println("Reading file")
	iniFile := ini.CreateFile()
	indented := ini.Load("indented", iniFile)

	log.Println("======== Indented File ========")
	log.Println(indented.GetSection("Section").GetProperty("key").Int64())
	log.Println(indented.GetSection("Section").GetSection("SubSection").GetProperty("key").String())
	log.Println(indented.GetSection("Section").GetSection("SubSection2").GetProperty("key").String())
	log.Println(indented.GetSection("Section").GetSection("SubSection2").GetSection("SubSubSection").GetProperty("key").String())
	log.Println(indented.GetSection("Section").GetSection("SubSection2").GetSection("SubSubSection").GetProperty("key").String())
	log.Println(indented.GetSection("Section").GetSection("SubSection2").GetSection("SubSubSection").GetSection("SubSubSubSection").GetProperty("key").String())
	log.Println(indented.GetSection("Section2").GetProperty("Key").String())

	iniFile = ini.CreateFile()
	iniFile.Config.SetSubSectionType(subsection.Seperated, ".")
	seperated := ini.Load("seperated", iniFile)
	log.Println("======== Seperated File ========")
	log.Println(seperated.GetSection("Section").GetProperty("key").Int64())
	log.Println(seperated.GetSection("Section").GetSection("SubSection").GetProperty("key").String())
	log.Println(seperated.GetSection("Section").GetSection("SubSection2").GetProperty("key").String())
	log.Println(seperated.GetSection("Section").GetSection("SubSection2").GetSection("SubSubSection").GetProperty("key").String())
	log.Println(seperated.GetSection("Section").GetSection("SubSection2").GetSection("SubSubSection").GetProperty("key").String())
	log.Println(seperated.GetSection("Section").GetSection("SubSection2").GetSection("SubSubSection").GetSection("SubSubSubSection").GetProperty("key").String())
	log.Println(seperated.GetSection("Section2").GetProperty("Key").String())
}
