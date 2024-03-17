package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hwhang0917/aoc-2023-go/utils"
)

// Given cube counts by the Elf
var cubeCounts = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func isPossible(rounds [][3]int) bool {
	var redCount, greenCount, blueCount int = 0, 0, 0
	for _, round := range rounds {
		for colorIdx, count := range round {
			switch colorIdx {
			case 0:
				redCount += count
				if redCount > cubeCounts["red"] {
					return false
				}
			case 1:
				greenCount += count
				if greenCount > cubeCounts["green"] {
					return false
				}
			case 2:
				blueCount += count
				if blueCount > cubeCounts["blue"] {
					return false
				}
			}

		}
	}
	return true
}

func parseGame(s string) (int, [][3]int) {
	sp := strings.Split(s, ":")
	strId := sp[0][5:]
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	roundsStr := strings.Split(sp[1], ";")
	var rounds [][3]int
	for _, roundStr := range roundsStr {
		round := [3]int{0, 0, 0}
		colorStrings := strings.Split(roundStr, ",")
		for _, colorString := range colorStrings {
			trimmedColor := strings.TrimSpace(colorString)
			splitColor := strings.Split(trimmedColor, " ")

			count, _ := strconv.Atoi(splitColor[0])
			color := splitColor[len(splitColor)-1]

			if color == "red" {
				round[0] += count
			} else if color == "green" {
				round[1] += count
			} else if color == "blue" {
				round[2] += count
			}
		}
		rounds = append(rounds, round)
	}

	return id, rounds
}

func solution(lines []string) int {
	answer := 0
	for _, line := range lines {
		id, rounds := parseGame(line)
		if isPossible(rounds) {
			answer += id
		}
	}
	return answer
}

func main() {
	lines := utils.ReadFileEachLine()
	fmt.Println(solution(lines))
}
