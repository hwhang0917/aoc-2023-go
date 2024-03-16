package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFileEachLine() []string {
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
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
