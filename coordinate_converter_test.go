package coordinate_converter

import "testing"

func TestDmsToDd1(t *testing.T) {

	got := DmsToDd(30, 15, 50)
	want := 30.26388888888889

	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestDmsToDd2(t *testing.T) {

	// 57°54′02″W
	got := DmsToDd(57, 45, 02)
	want := 57.75055555555556

	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestDmsToDd3(t *testing.T) {

	// 63°19′15″S
	got := DmsToDd(63, 19, 15)
	want := 63.32083333333334

	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}
