package interfaces

import "testing"

func TestInterfaces_RectangleMeasure(t *testing.T) {
	r := Rectangle{width: 5, height: 5}

	area := measureArea(r)
	if area != 25 {
		t.Errorf("value %f is wrong", area)
	}

	perimeter := measurePerimeter(r)
	if perimeter != 20 {
		t.Errorf("value %f is wrong", area)
	}
}

func TestInterfaces_CircleMeasure(t *testing.T) {
	c := Circle{radius: 5}

	area := measureArea(c)
	if area != 78.539818 {
		t.Errorf("value %f is wrong", area)
	}

	perimeter := measurePerimeter(c)
	if perimeter != 31.415928 {
		t.Errorf("value %f is wrong", area)
	}
}
