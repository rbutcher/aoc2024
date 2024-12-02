package solution

import (
	_ "embed"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
)

//go:embed input/day2.txt
var day2Input string

type day2 struct {
	data [][]int
}

func NewDay2() Solution {
	var result [][]int
	rowSplit := strings.Split(day2Input, "\n")
	for line, row := range rowSplit {
		numberSplit := strings.Split(row, " ")
		var numbers []int
		for col, rawNumber := range numberSplit {
			n, err := strconv.Atoi(rawNumber)
			if err != nil {
				log.Fatal().Msgf("failed to parse input \"%s\" is not a number line=%d col=%d", rawNumber, line, col)
			}
			numbers = append(numbers, n)
		}

		result = append(result, numbers)
	}

	return &day2{
		data: result,
	}
}

func (d *day2) String() string {
	b := strings.Builder{}
	for _, row := range d.data {
		for _, n := range row {
			b.WriteString(strconv.Itoa(n))
			b.WriteString(" ")
		}
		b.WriteString("\n")
	}

	return b.String()
}

// Increasing or decreasing AND between 1 and 3 apart
// Count safe reports

func (d *day2) Part1() (string, error) {
	count := 0
	for _, row := range d.data {
		if isSafe(row) {
			count += 1
		}
	}

	return strconv.Itoa(count), nil
}

func (d *day2) Part2() (string, error) {
	count := 0
	for _, row := range d.data {
		if isSafe(row) {
			count += 1
		}
	}

	return strconv.Itoa(count), nil
}

func isSafe(numbers []int) bool {
	l := log.With().Ints("row", numbers).Logger()

	fst := numbers[0]
	snd := numbers[1]
	asc := snd > fst
	diff := abs(snd - fst)
	if diff > 3 || diff == 0 {
		l.Debug().
			Bool("safe", false).
			Int("left", 0).
			Int("right", 1).
			Msg("diff is too large")
		return false
	}

	last := snd
	for _, curr := range numbers[2:] {
		casc := curr > last
		if (!casc && asc) || (casc && !asc) {
			l.Debug().
				Bool("safe", false).
				Int("last", last).
				Int("current", curr).
				Msg("set is not all ascending or descending")
			return false
		}

		diff = abs(curr - last)
		if diff > 3 || diff == 0 {
			l.Debug().
				Bool("safe", false).
				Int("last", last).
				Int("current", curr).
				Msg("diff is too large or zero")
			return false
		}
		last = curr
	}

	log.Debug().
		Ints("row", numbers).
		Bool("safe", true).
		Send()

	return true
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}

	return v
}
