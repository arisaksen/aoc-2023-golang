package main

import (
	utils "aoc/golang"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//puzzle1Input := utils.ReadLines("day01/puzzle1.txt")
	//fmt.Println("Part 1: ", getCalibrationValues(puzzle1Input))
	//
	puzzle2Input := utils.ReadLines("day01/puzzle2.txt")
	fmt.Println("Part 2: ", getCalibrationValues(puzzle2Input))
}

func getCalibrationValues(lines []string) int {
	totalCalibrationValue := 0
	for _, line := range lines {
		firstDigit := getFirstDigit(line)
		lastDigit := getLastDigit(line)
		combined := strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)
		totalCalibrationValue += fromString(combined)
	}
	return totalCalibrationValue
}

func fromString(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}

func getFirstDigit(s string) int {
	for i := 0; i < len(s); i++ {
		if found, d := containsSpelledDigit(s[:i]); found {
			return d
		} else if unicode.IsDigit(rune(s[i])) {
			return int(s[i] - '0')
		}
	}
	panic("No digit found in " + s)
}

func getLastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if found, d := containsSpelledDigit(s[i:]); found {
			return d
		} else if unicode.IsDigit(rune(s[i])) {
			return int(s[i] - '0')
		}
	}
	panic("No digit found in " + s)
}

var spelledDigits = map[string]int{
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

func containsSpelledDigit(s string) (bool, int) {
	for k, v := range spelledDigits {
		if strings.Contains(s, k) {
			return true, v
		}
	}
	return false, 0
}
