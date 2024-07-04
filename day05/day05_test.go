package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 35
		actual := part1("example.txt")

		assert.Equal(t, expected, actual)
	})

}

func TestSeeds(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := map[string][]int{"seeds": {1, 2, 3, 4}}
		actual := getSeeds("seeds: 1 2 3 4")

		assert.Equal(t, expected, actual)
	})

}

//func TestPart2(t *testing.T) {
//
//	t.Run("Part 2", func(t *testing.T) {
//		expected := 71503
//		actual := part2("test-input.txt")
//		assert.Equal(t, expected, actual)
//	})
//
//}
