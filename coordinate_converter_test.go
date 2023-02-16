package coordinate_converter

import "testing"

func TestDmsToDd(t *testing.T) {

	want := 30.2638888888
	got := DmsToDd(30, 15, 50)

	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}
