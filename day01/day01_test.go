package main

import (
	utils "aoc/golang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalibrations(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		exampleLines := utils.ReadLines("example1.txt")
		expected := []int{12, 38, 15, 77}
		actual := getCalibrations(exampleLines)
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 1", func(t *testing.T) {
		exampleLines := utils.ReadLines("example1.txt")
		expected := 142
		actual := sumUpCalibrations(getCalibrations(exampleLines))
		assert.Equal(t, expected, actual)
	})

}
