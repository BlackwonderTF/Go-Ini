package main

import (
	"log"

	"github.com/BlackwonderTF/go-ini"
)

func main() {
	log.Println("Reading file")
	iniFile := ini.Load("test")
	log.Println(iniFile.Section("Owner").GetProperty("Dead").Bool())
}
