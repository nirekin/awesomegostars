package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitle(t *testing.T) {

	s := "# Title"
	ti := getTitle(s)
	assert.Equal(t, "title", ti)

	s = "# Title Part1 Part2"
	ti = getTitle(s)
	assert.Equal(t, "title-part1-part2", ti)

	s = "############################## Title Part1 Part2"
	ti = getTitle(s)
	assert.Equal(t, "title-part1-part2", ti)

	s = "############################## Title Part1 Part2          "
	ti = getTitle(s)
	assert.Equal(t, "title-part1-part2", ti)

}
