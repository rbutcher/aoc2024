package solution

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

const day5TestInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestDay5_Part1(t *testing.T) {
	day5Input = day5TestInput

	sut := NewDay5()
	res, err := sut.Part1()

	require.NoError(t, err)
	assert.Equal(t, "143", res)
}
func TestDay5_Part2(t *testing.T) {
	day5Input = day5TestInput

	sut := NewDay5()
	res, err := sut.Part2()

	require.NoError(t, err)
	assert.Equal(t, "123", res)
}
