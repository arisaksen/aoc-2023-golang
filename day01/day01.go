package main

import (
	utils "aoc/golang"
	"fmt"
	"unicode"
)

func main() {
	//puzzle1Input := utils.ReadLines("day01/puzzle1.txt")
	//fmt.Println("Part 1: ", getSeeds(puzzle1Input))
	//
	puzzle1Input := utils.ReadLines("day01/puzzle1.txt")
	calibrations := getCalibrations(puzzle1Input)
	sumCalibrations := sumUpCalibrations(calibrations)
	fmt.Println("Part 1: ", sumCalibrations)
}

func getCalibrations(lines []string) []int {
	var calibrations []int

	for _, line := range lines {
		var firstDigit, lastDigit rune
		var firstDigitFound bool
		var lastDigitIndex int

		fmt.Println(line)
		for y, char := range line {
			if unicode.IsDigit(char) && !firstDigitFound {
				firstDigit = char
				firstDigitFound = true
				fmt.Printf("firstDigit found at %d: %c \n", y, char)
			}
			if unicode.IsDigit(char) {
				lastDigit = char
				lastDigitIndex = y
			}
		}
		fmt.Printf("lastDigit found at %d: %c \n", lastDigitIndex, lastDigit)
		firstDigitInt := int(firstDigit - '0')
		lastDigitInt := int(lastDigit - '0')
		combineDigits := firstDigitInt*10 + lastDigitInt

		calibrations = append(calibrations, combineDigits)
	}

	return calibrations
}

func sumUpCalibrations(calibrations []int) int {
	var sum = 0
	for _, number := range calibrations {
		sum += number
	}

	return sum
}
