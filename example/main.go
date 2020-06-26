package main

import (
	"log"

	"github.com/BlackwonderTF/go-ini"
)

func main() {
	log.Println("Reading file")
	iniFile := ini.Load("test")
	log.Println(iniFile.Section("Database").GetProperty("Port").Int64())
	log.Println(iniFile.Section("settings").GetProperty("comment").String())
}
