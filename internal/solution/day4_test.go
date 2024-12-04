package solution

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const day4TestInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestDay4_Part1(t *testing.T) {
	day4Input = day4TestInput

	sut := NewDay4()
	res, err := sut.Part1()

	require.NoError(t, err)
	assert.Equal(t, "18", res)
}

func TestDay4_Part2(t *testing.T) {
	day4Input = day4TestInput

	assert.Fail(t, "not implemented")
}
