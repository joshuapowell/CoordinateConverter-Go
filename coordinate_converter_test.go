package coordinate_converter

import "testing"

func TestDmsToDd(t *testing.T) {

	got := DmsToDd(30, 15, 50)
	want := 30.26388888888889

	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}
