package raycasting_test

import (
	"github.com/parsaaes/go-ray-casting"
	"testing"
)

func TestPolygon_Contains(t *testing.T) {
	polygon := raycasting.Polygon{
		Points: [][]float64{
			{1.0, 1.0},
			{2.0, 3.0},
			{3.0, 2.0},
		},
	}

	tests := []struct {
		point      [2]float64
		expectedIn bool
	}{
		{
			point:      [2]float64{2.0, 2.0},
			expectedIn: true,
		},
		{
			point:      [2]float64{2.0, 4.0},
			expectedIn: false,
		},
	}
	for _, test := range tests {
		if polygon.Contains(test.point) != test.expectedIn {
			t.Errorf("Point %v is expected to be in the polygon %v (%v)",
				test.point, polygon, test.expectedIn)
		}
	}
}
