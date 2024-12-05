package solution

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/rbutcher/aoc2024/internal/helpers"
	"github.com/rs/zerolog/log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input/day5.txt
var day5Input string

type day5 struct {
	pageOrders  []helpers.Point[int]
	updatePages [][]int
}

func NewDay5() Solution {
	rangeSplit := strings.Split(day5Input, "\n\n")
	if len(rangeSplit) != 2 {
		log.Fatal().Msgf("invalid input: could not split orders from updates")
	}

	orderSplit := strings.Split(rangeSplit[0], "\n")
	var orders []helpers.Point[int]
	for i, o := range orderSplit {
		pointSplit := strings.Split(o, "|")
		if len(pointSplit) != 2 {
			log.Fatal().Msgf("invalid input: line split incorrectly line=%d value=%s", i, o)
		}

		l, err := strconv.Atoi(pointSplit[0])
		if err != nil {
			log.Fatal().Msgf("invalid input: left is not a number line=%d value=%s", i, pointSplit[0])
		}

		r, err := strconv.Atoi(pointSplit[1])
		if err != nil {
			log.Fatal().Msgf("invalid input: right is not a number line=%d value=%s", i, pointSplit[1])
		}
		orders = append(orders, helpers.Point[int]{X: l, Y: r})
	}

	updateSplit := strings.Split(rangeSplit[1], "\n")
	var updates [][]int
	for i, u := range updateSplit {
		pageSplit := strings.Split(u, ",")

		var update []int
		for _, p := range pageSplit {
			n, err := strconv.Atoi(p)
			if err != nil {
				log.Fatal().Msgf("invalid input: update contains non-numbers line=%d value=%s", i, p)
			}

			update = append(update, n)
		}

		updates = append(updates, update)
	}

	return &day5{
		pageOrders:  orders,
		updatePages: updates,
	}
}

func (d *day5) String() string {
	b := &strings.Builder{}

	for _, p := range d.pageOrders {
		b.WriteString(p.String() + "\n")
	}

	for _, u := range d.updatePages {
		b.WriteString(fmt.Sprintf("%v\n", u))
	}

	return b.String()
}

func (d *day5) Part1() (string, error) {
	sum := 0
	for _, pages := range d.updatePages {
		valid := true
		for i, x := range pages {
			rules := d.getPageRules(x)
			for _, r := range rules {
				if slices.Contains(pages[:i], r.Y) {
					valid = false
					log.Debug().
						Ints("update", pages).
						Str("violation", r.String()).
						Send()
					break
				}
			}

			if !valid {
				break
			}
		}

		if valid {
			i := len(pages) / 2
			sum += pages[i]
		}
	}

	return strconv.Itoa(sum), nil
}

func (d *day5) Part2() (string, error) {
	return "", errors.New("not implemented")
}

func (d *day5) getPageRules(x int) []helpers.Point[int] {
	var result []helpers.Point[int]
	for _, p := range d.pageOrders {
		if p.X == x {
			result = append(result, p)
		}
	}

	return result
}
