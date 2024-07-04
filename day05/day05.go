package main

import (
	utils "aoc/golang"
	"fmt"
)

func main() {
	fmt.Println("Part 1: ", part1("day05/puzzle.txt"))

}

func part1(filename string) int {
	puzzle2Input := utils.ReadLines(filename)

	seeds := getSeeds(puzzle2Input[0])
	fmt.Println(seeds)

	for _, line := range puzzle2Input {
		fmt.Println(line)
	}

	return 1 * 2 * 3
}

func getSeeds(seeds string) string {
	var x = ""

	return x
}
