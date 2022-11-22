package tests

import (
	"TDD/model/interfaces"
	"TDD/model/mock"
	"testing"
)

/*func TestPerimeter(t *testing.T) {
	rectangle := mock.Rectangle{Width: 10.0, Height: 10.0}
	got := mock.Rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %2.f", got, want)
	}
}*/

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape interfaces.Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %2.f", got, want)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := mock.Rectangle{Width: 12.0, Height: 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := mock.Circle{10}
		checkArea(t, circle, 314.15922653589793)
	})

}

func TestAreaList(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   interfaces.Shape
		hasArea float64
	}{
		{name: "Rectagle", shape: mock.Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: mock.Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Trangle", shape: mock.Triangle{Base: 12.0, Height: 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}
