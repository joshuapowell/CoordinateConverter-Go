package coordinate_converter

import "testing"

func TestDmsToDd1(t *testing.T) {

	got := ConvertDmsToDd(30, 15, 50, North)
	want := 30.26388888888889

	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestDmsToDd2(t *testing.T) {

	// 57°54′02″W
	got := ConvertDmsToDd(57, 45, 02, East)
	want := 57.75055555555556

	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestDmsToDd3(t *testing.T) {

	// 57°54′02″W
	got := ConvertDmsToDd(57, 45, 02, West)
	want := -57.75055555555556

	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestDmsToDd4(t *testing.T) {

	// 63°19′15″S
	got := ConvertDmsToDd(63, 19, 15, South)
	want := -63.32083333333334

	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}
