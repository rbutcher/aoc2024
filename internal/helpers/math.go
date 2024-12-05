package helpers

import "fmt"

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Min[T Signed](right, left T) T {
	if left > right {
		return right
	}

	return left
}

func Max[T Signed](right, left T) T {
	if left < right {
		return right
	}

	return left
}

func Abs[T Signed](value T) T {
	if value < 0 {
		return value * -1
	}

	return value
}

func Clamp[T Signed](value, min, max T) T {
	if value > max {
		return max
	}

	if value < min {
		return min
	}

	return value
}

type Point[T Signed] struct {
	X T
	Y T
}

func (p Point[T]) String() string {
	return fmt.Sprintf("(%v,%v)", p.X, p.Y)
}
