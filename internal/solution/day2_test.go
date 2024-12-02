package solution

import (
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const day2TestInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestDay2_Part1(t *testing.T) {
	day2Input = day2TestInput

	sut := NewDay2()
	res, err := sut.Part1()

	require.NoError(t, err)
	assert.Equal(t, res, "2")
}

func TestDay2_Part2(t *testing.T) {
	day2Input = day2TestInput

	sut := NewDay2()
	res, err := sut.Part2()

	require.NoError(t, err)
	assert.Equal(t, res, "4")
}
