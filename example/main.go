package main

import (
	"log"

	"github.com/BlackwonderTF/go-ini"
)

func main() {
	log.Println("Reading file")
	iniFile := ini.Load("test")
	// log.Println(iniFile.GetSection("Owner").GetProperty("dead").Bool())
	log.Println(iniFile.GetSection("Section").GetProperty("key").Int64())
	log.Println(iniFile.GetSection("Section").GetSection("SubSection").GetProperty("key").String())
	log.Println(iniFile.GetSection("Section").GetSection("SubSection2").GetProperty("key").String())
	log.Println(iniFile.GetSection("Section").GetSection("SubSection2").GetSection("SubSubSection").GetProperty("key").String())
	log.Println(iniFile.GetSection("Section").GetSection("SubSection2").GetSection("SubSubSection").GetProperty("key").String())
	log.Println(iniFile.GetSection("Section").GetSection("SubSection2").GetSection("SubSubSection").GetSection("SubSubSubSection").GetProperty("key").String())
	log.Println(iniFile.GetSection("Section2").GetProperty("Key").String())
}
