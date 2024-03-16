package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hwhang0917/aoc-2023-go/utils"
)

func partOneSolution(lines []string) string {
	numregex := regexp.MustCompile(`\d`)
	var calibrations []int
	for _, line := range lines {
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
	return strconv.Itoa(solution)
}

func partTwoSolution(lines []string) string {
	numericStringMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	var keys []string
	for k := range numericStringMap {
		keys = append(keys, k)
	}

	regexString := strings.Join([]string{"(", `\d|`, strings.Join(keys, "|"), ")"}, "")

	numregex := regexp.MustCompile(regexString)

	var calibrations []int
	for _, line := range lines {
		matches := numregex.FindAllString(line, -1)
		var numbers []int
		for _, match := range matches {
			if num, ok := numericStringMap[match]; ok {
				numbers = append(numbers, num)
				continue
			} else {
				num, _ := strconv.Atoi(match)
				numbers = append(numbers, num)
			}
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
	return strconv.Itoa(solution)
}

func main() {
	lines := utils.ReadFileEachLine()

	fmt.Println("Part I: ", partOneSolution(lines))
	fmt.Println("Part II: ", partTwoSolution(lines))
}
