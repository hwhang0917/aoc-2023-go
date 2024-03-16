package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/hwhang0917/aoc-2023-go/utils"
)

var numericStringMap = map[string]int{
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

func containsNumericString(s string) (bool, int) {
	for k, v := range numericStringMap {
		if strings.Contains(s, k) {
			return true, v
		}
	}
	return false, 0
}

func partOne(lines []string) int {
	totalCalibration := 0
	for _, line := range lines {
		var firstDigit, secondDigit int
		// Get first digit
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				firstDigit = int(line[i] - '0')
				break
			}
		}
		// Get last last digit
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				secondDigit = int(line[i] - '0')
				break
			}
		}
		totalCalibration += (firstDigit*10 + secondDigit)
	}
	return totalCalibration
}

func partTwo(lines []string) int {
	totalCalibration := 0
	for _, line := range lines {
		var firstDigit, secondDigit int
		// Get first digit
		for i := 0; i < len(line); i++ {
			if found, d := containsNumericString(line[:i]); found {
				firstDigit = d
				break
			} else if unicode.IsDigit(rune(line[i])) {
				firstDigit = int(line[i] - '0')
				break
			}
		}
		// Get last last digit
		for i := len(line) - 1; i >= 0; i-- {
			if found, d := containsNumericString(line[i:]); found {
				secondDigit = d
				break
			} else if unicode.IsDigit(rune(line[i])) {
				secondDigit = int(line[i] - '0')
				break
			}
		}
		totalCalibration += (firstDigit*10 + secondDigit)
	}
	return totalCalibration
}

func main() {
	lines := utils.ReadFileEachLine()

	fmt.Println("Part I: ", partOne(lines))
	fmt.Println("Part II: ", partTwo(lines))
}
