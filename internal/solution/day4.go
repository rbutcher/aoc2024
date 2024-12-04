package solution

import (
	_ "embed"
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
)

type direction int

const (
	directionUp direction = iota
	directionUpRight
	directionRight
	directionDownRight
	directionDown
	directionDownLeft
	directionLeft
	directionUpLeft
)

//go:embed input/day4.txt
var day4Input string

type day4 struct {
	data []string
}

func NewDay4() Solution {
	split := strings.Split(day4Input, "\n")
	return &day4{data: split}
}

func (d *day4) Part1() (string, error) {
	count := 0
	for ix := 0; ix < len(d.data[0]); ix++ {
		for iy := 0; iy < len(d.data); iy++ {
			current := d.data[iy][ix]
			if current == 'X' {
				for dir := 0; dir <= int(directionUpLeft); dir++ {
					if d.checkXmasDirection(d.data, ix, iy, direction(dir)) {
						count++
					}
				}
			}
		}
	}

	return strconv.Itoa(count), nil
}

func (d *day4) Part2() (string, error) {
	return "", errors.New("not implemented")
}

func (d *day4) checkXmasDirection(data []string, x, y int, dir direction) bool {
	l := log.With().
		Str("context", "solution.checkXmasDirection").
		Int("x", x).
		Int("y", y).
		Int("direction", int(dir)).
		Logger()

	if len(data) == 0 {
		l.Debug().Msg("passed data has no values")
		return false
	}

	if y < 0 || y > len(data) {
		l.Debug().Msg("passed y is out of bounds")
		return false
	}

	if x < 0 || x > len(data[0]) {
		l.Debug().Msg("passed x is out of bounds")
		return false
	}

	var dx, dy int
	switch dir {
	case directionUp:
		dx = 0
		dy = -1

	case directionUpRight:
		dx = 1
		dy = -1

	case directionRight:
		dx = 1
		dy = 0

	case directionDownRight:
		dx = 1
		dy = 1

	case directionDown:
		dx = 0
		dy = 1

	case directionDownLeft:
		dx = -1
		dy = 1

	case directionLeft:
		dx = -1
		dy = 0

	case directionUpLeft:
		dx = -1
		dy = -1

	default:
		return false
	}

	searchLetters := []rune{'M', 'A', 'S'}
	x += dx
	y += dy
	for _, r := range searchLetters {
		if y < 0 || y >= len(data) {
			return false
		}

		if x < 0 || x >= len(data[y]) {
			return false
		}

		c := rune(data[y][x])
		if r != c {
			return false
		}

		x += dx
		y += dy
	}

	return true
}
