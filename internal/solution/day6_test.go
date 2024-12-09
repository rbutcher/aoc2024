package solution

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var day6TestInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestDay6_Part1(t *testing.T) {
	day6Input = day6TestInput

	var sut = NewDay6()
	res, err := sut.Part1()

	require.NoError(t, err)
	assert.Equal(t, "41", res)
}
