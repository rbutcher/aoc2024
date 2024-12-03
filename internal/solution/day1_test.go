package solution

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const day1TestInput string = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestDay1_Part1(t *testing.T) {
	day1Input = day1TestInput

	sut := NewDay1()
	res, err := sut.Part1()

	require.NoError(t, err)
	assert.Equal(t, "11", res)
}

func TestDay1_Part2(t *testing.T) {
	day1Input = day1TestInput

	sut := NewDay1()
	res, err := sut.Part2()

	require.NoError(t, err)
	assert.Equal(t, "31", res)
}
