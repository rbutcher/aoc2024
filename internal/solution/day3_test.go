package solution

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDay3_Part1(t *testing.T) {
	const day3TestInput = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	day3Input = day3TestInput

	sut := NewDay3()
	res, err := sut.Part1()

	require.NoError(t, err)
	assert.Equal(t, "161", res)
}

func TestDay3_Part2(t *testing.T) {
	const day3TestInput = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	day3Input = day3TestInput

	sut := NewDay3()
	res, err := sut.Part2()

	require.NoError(t, err)
	assert.Equal(t, "48", res)
}
