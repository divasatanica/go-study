package main

import (
	"testing"
	"math"
)

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %.2f want %.2f.", shape, got, want)
		}
	}

	areaTests := []struct {
		shape Shape
		want  float64
		name  string
	}{
		{Rectangle{12, 6}, 72.0, "Rectangle"},
		{Circle{10}, math.Pi * 100, "Circle"},
		{Triangle{10, 5}, 25, "Triangle"},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}
