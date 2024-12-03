package solution

import (
	_ "embed"
	"github.com/rbutcher/aoc2024/internal/helpers"
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
			continue
		}

		for i := 0; i < len(row); i++ {
			numbers := removeAt(row, i)
			if isSafe(numbers) {
				count += 1
				break
			}
		}
	}

	return strconv.Itoa(count), nil
}

func isSafe(numbers []int) bool {
	asc := numbers[0] < numbers[len(numbers)-1]
	for i := 1; i < len(numbers); i++ {
		current := numbers[i]
		last := numbers[i-1]

		diff := current - last
		if (diff > 0 && !asc) || (diff < 0 && asc) ||
			(diff == 0) || (helpers.Abs(diff) > 3) {
			return false
		}
	}

	return true
}

func removeAt(numbers []int, i int) []int {
	var result []int
	result = append(result, numbers[:i]...)
	result = append(result, numbers[i+1:]...)
	return result
}
