package solution

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/rbutcher/aoc2024/internal/helpers"
	"strconv"
	"strings"
)

//go:embed input/day6.txt
var day6Input string

type day6 struct {
	data  []string
	start helpers.Point[int]
}

func NewDay6() Solution {
	data := strings.Split(day6Input, "\n")
	start, ok := findStart(data)
	if !ok {
		panic("failed to parse input")
	}

	return &day6{
		data:  strings.Split(day6Input, "\n"),
		start: start,
	}
}

func (d *day6) Part1() (string, error) {
	allVisited := make(map[helpers.Point[int]]bool)
	dir := helpers.DirectionUp
	visited, last, ok := walkToObstacle(d.data, dir, d.start)
	allVisited = mapZip(allVisited, visited)
	for ok {
		dir = dir.Rotate90CW()
		visited, last, ok = walkToObstacle(d.data, dir, last)
		allVisited = mapZip(allVisited, visited)
	}

	return strconv.Itoa(len(allVisited)), nil
}

func (d *day6) Part2() (string, error) {
	return "", errors.New("not implemented")
}

func walkToObstacle(data []string, dir helpers.Direction, start helpers.Point[int]) (map[helpers.Point[int]]bool, helpers.Point[int], bool) {
	delta := dir.GetMoveDelta()
	visited := map[helpers.Point[int]]bool{
		start: true,
	}

	current := start
	if delta.X != 0 {
		for x := start.X + delta.X; x >= 0 && x < len(data[start.Y]); x += delta.X {
			c := data[start.Y][x]
			if c == '#' {
				return visited, current, true
			}

			current = helpers.Point[int]{X: x, Y: start.Y}
			visited[current] = true
		}
	} else if delta.Y != 0 {
		for y := start.Y + delta.Y; y >= 0 && y < len(data); y += delta.Y {
			c := data[y][start.X]
			if c == '#' {
				return visited, current, true
			}

			current = helpers.Point[int]{X: start.X, Y: y}
			visited[current] = true
		}

	} else {
		msg := fmt.Sprintf("invalid move direction: %s", dir)
		panic(msg)
	}

	return visited, current, false
}

func mapZip[K comparable, V any](l map[K]V, r map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range l {
		result[k] = v
	}

	for k, v := range r {
		result[k] = v
	}

	return result
}

func findStart(data []string) (helpers.Point[int], bool) {
	for y := 0; y < len(data); y++ {
		for x, v := range data[y] {
			if v == '^' {
				return helpers.Point[int]{X: x, Y: y}, true
			}
		}
	}

	return helpers.PointZero[int](), false
}
