package main

import (
	"log"

	"cq-fth.com/go-utilities/internal"
)

func init() {
	log.SetFlags(log.Llongfile | log.LUTC | log.LstdFlags)
}

func main() {
	_, err := internal.CreateCSV()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created!")
}
