package solution

import (
	_ "embed"
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
					if d.checkStringDirection(ix, iy, direction(dir), "MAS") {
						count++
					}
				}
			}
		}
	}

	return strconv.Itoa(count), nil
}

func (d *day4) Part2() (string, error) {
	count := 0
	for x := 0; x < len(d.data[0]); x++ {
		for y := 0; y < len(d.data); y++ {
			current := d.data[y][x]
			if current == 'A' && d.checkForXMas(x, y) {
				count++
			}
		}
	}

	return strconv.Itoa(count), nil
}

func (d *day4) checkStringDirection(x, y int, dir direction, check string) bool {
	l := log.With().
		Str("context", "solution.checkStringDirection").
		Int("x", x).
		Int("y", y).
		Int("direction", int(dir)).
		Logger()

	if len(d.data) == 0 {
		l.Debug().Msg("passed data has no values")
		return false
	}

	if y < 0 || y > len(d.data) {
		l.Debug().Msg("passed y is out of bounds")
		return false
	}

	if x < 0 || x > len(d.data[0]) {
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

	x += dx
	y += dy
	for _, r := range check {
		if y < 0 || y >= len(d.data) {
			return false
		}

		if x < 0 || x >= len(d.data[y]) {
			return false
		}

		c := rune(d.data[y][x])
		if r != c {
			return false
		}

		x += dx
		y += dy
	}

	return true
}

func (d *day4) checkForXMas(x, y int) bool {
	l := log.With().Str("context", "solution.checkForXMas").Logger()
	if y-1 < 0 || y+1 >= len(d.data) {
		l.Debug().Msg("y is out of range")
		return false
	}

	if x-1 < 0 || x+1 >= len(d.data[y]) {
		l.Debug().Msg("x is out of range")
		return false
	}

	isM := func(r uint8) bool { return r == 'M' }
	isS := func(r uint8) bool { return r == 'S' }

	// check up
	if isM(d.data[y-1][x-1]) && isM(d.data[y-1][x+1]) && isS(d.data[y+1][x-1]) && isS(d.data[y+1][x+1]) {
		return true
	}

	// check right
	if isM(d.data[y-1][x+1]) && isM(d.data[y+1][x+1]) && isS(d.data[y-1][x-1]) && isS(d.data[y+1][x-1]) {
		return true
	}

	// check down
	if isM(d.data[y+1][x-1]) && isM(d.data[y+1][x+1]) && isS(d.data[y-1][x-1]) && isS(d.data[y-1][x+1]) {
		return true
	}

	// check left
	if isM(d.data[y-1][x-1]) && isM(d.data[y+1][x-1]) && isS(d.data[y-1][x+1]) && isS(d.data[y+1][x+1]) {
		return true
	}

	return false
}
