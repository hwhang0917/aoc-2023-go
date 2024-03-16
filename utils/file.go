package utils

import (
	"log"
	"os"
)

func ReadFile() *os.File {
	// Read file from command line argument or use default filename
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = "input.txt"
	}
	log.Println("Reading file: ", filename)

	// open file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
