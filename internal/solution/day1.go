package solution

import (
	_ "embed"
	"github.com/rs/zerolog/log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input/day1.txt
var day1Input string

type day1 struct {
	left  []int
	right []int
}

func NewDay1() Solution {
	var left []int
	var right []int

	lines := strings.Split(day1Input, "\n")
	for i, line := range lines {
		split := strings.Split(line, "   ")
		if len(split) != 2 {
			log.Fatal().Msgf("failed to parse day1 input, too many numbers at line %d", i)
		}

		l, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal().Msgf("failed to parse day 1 input, left string is not number %s, line %d", split[0], i)
		}

		r, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal().Msgf("failed to parse day 1 input, right string is not number %s, line %d", split[0], i)
		}

		left = append(left, l)
		right = append(right, r)
	}

	return &day1{
		left:  left,
		right: right,
	}
}

func (d *day1) Part1() (string, error) {
	slices.Sort(d.left)
	slices.Sort(d.right)

	totalDistance := 0
	for i := range d.left {
		dist := d.left[i] - d.right[i]
		if dist < 0 {
			dist *= -1
		}
		totalDistance += dist
		log.Debug().
			Int("left", d.left[i]).
			Int("right", d.right[i]).
			Int("dist", dist).
			Int("total", totalDistance).
			Send()
	}

	return strconv.Itoa(totalDistance), nil
}

func (d *day1) Part2() (string, error) {
	similarity := 0
	for _, l := range d.left {
		count := 0
		for _, r := range d.right {
			if l == r {
				count += 1
			}
		}

		similarity += count * l
		log.Debug().
			Int("count", count).
			Int("left", l).
			Int("similarity", similarity).
			Send()
	}

	return strconv.Itoa(similarity), nil
}
