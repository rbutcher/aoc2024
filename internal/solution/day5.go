package solution

import (
	_ "embed"
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
	for i := 0; i < len(d.updatePages); i++ {
		if d.isValidPageOrder(i) {
			sum += d.updatePages[i][len(d.updatePages[i])/2]
		}
	}

	return strconv.Itoa(sum), nil
}

func (d *day5) Part2() (string, error) {
	sum := 0
	for i := 0; i < len(d.updatePages); i++ {
		if !d.isValidPageOrder(i) {
			ro := d.reorderPages(i)

			log.Debug().
				Ints("orig", d.updatePages[i]).
				Ints("reorder", ro).
				Int("mid", ro[len(ro)/2]).
				Send()

			sum += ro[len(ro)/2]
		}
	}

	return strconv.Itoa(sum), nil
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

func (d *day5) isValidPageOrder(i int) bool {
	var pageOrders = d.updatePages[i]
	for ip, p := range pageOrders {
		rules := d.getPageRules(p)
		for _, r := range rules {
			if slices.Contains(pageOrders[:ip], r.Y) {
				return false
			}
		}
	}

	return true
}

func (d *day5) reorderPages(i int) []int {
	result := []int{d.updatePages[i][0]}
	for _, x := range d.updatePages[i][1:] {
		before := len(result)
		rules := d.getPageRules(x)
		for _, r := range rules {
			fi := slices.Index(result, r.Y)
			if fi != -1 {
				log.Debug().
					Ints("orig", d.updatePages[i]).
					Ints("ro", result).
					Str("rule", r.String()).
					Int("curr", x).
					Int("prev", before).
					Int("after", fi).
					Send()
				before = helpers.Min(fi, before)
			}
		}

		result = slices.Insert(result, before, x)
		log.Debug().Int("before", before).Ints("ro", result).Send()
	}

	return result
}
