package helpers

import "fmt"

type Direction int

const (
	DirectionUp Direction = iota
	DirectionUpRight
	DirectionRight
	DirectionDownRight
	DirectionDown
	DirectionDownLeft
	DirectionLeft
	DirectionUpLeft
)

func DirectionMin() Direction {
	return DirectionUp
}

func DirectionMax() Direction {
	return DirectionUpLeft
}

func (d Direction) GetMoveDelta() Point[int] {
	switch d {
	case DirectionUp:
		return Point[int]{X: 0, Y: -1}

	case DirectionUpRight:
		return Point[int]{X: 1, Y: -1}

	case DirectionRight:
		return Point[int]{X: 1, Y: 0}

	case DirectionDownRight:
		return Point[int]{X: 1, Y: 1}

	case DirectionDown:
		return Point[int]{X: 0, Y: 1}

	case DirectionDownLeft:
		return Point[int]{X: -1, Y: 1}

	case DirectionLeft:
		return Point[int]{X: -1, Y: 0}

	case DirectionUpLeft:
		return Point[int]{X: -1, Y: -1}

	default:
		msg := fmt.Sprintf("value is not a valid direction")
		panic(msg)
	}
}

func (d Direction) Rotate90CW() Direction {
	return Direction((d.Int() + 2) % 8)
}

func (d Direction) Int() int {
	return int(d)
}

func (d Direction) String() string {
	switch d {
	case DirectionUp:
		return "up"

	case DirectionUpRight:
		return "up_right"

	case DirectionRight:
		return "right"

	case DirectionDownRight:
		return "down_right"

	case DirectionDown:
		return "down"

	case DirectionDownLeft:
		return "down_left"

	case DirectionLeft:
		return "left"

	case DirectionUpLeft:
		return "up_left"

	default:
		return "not_a_direction"
	}
}
