package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = "input.txt"
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numregex := regexp.MustCompile(`\d`)

	var calibrations []int

	for scanner.Scan() {
		line := scanner.Text()
		matches := numregex.FindAllString(line, -1)
		var numbers []int
		for _, match := range matches {
			num, _ := strconv.Atoi(match)
			numbers = append(numbers, num)
		}

		if len(numbers) == 1 {
			calibrations = append(calibrations, numbers[0]*11)
		} else {
			calibrations = append(calibrations, numbers[0]*10+numbers[len(numbers)-1])
		}
	}

	solution := 0
	for _, calibration := range calibrations {
		solution += calibration
	}

	fmt.Println(solution)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
