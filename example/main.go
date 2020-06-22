package main

import (
	"log"

	"github.com/BlackwonderTF/go-ini"
)

func main() {
	log.Println("Reading file")
	ini.Parse("example/test")
}
