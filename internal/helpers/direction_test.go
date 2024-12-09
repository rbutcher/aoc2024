package helpers

import "testing"

func TestDirection_Rotate90CW(t *testing.T) {
	testCases := []struct {
		input    Direction
		expected Direction
	}{
		{DirectionUp, DirectionRight},
		{DirectionRight, DirectionDown},
		{DirectionDown, DirectionLeft},
		{DirectionLeft, DirectionUp},
		{DirectionUpRight, DirectionDownRight},
		{DirectionDownRight, DirectionDownLeft},
		{DirectionDownLeft, DirectionUpLeft},
		{DirectionUpLeft, DirectionUpRight},
	}

	for _, tc := range testCases {
		t.Run(tc.input.String(), func(t *testing.T) {
			actual := tc.input.Rotate90CW()
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
