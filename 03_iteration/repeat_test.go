package iteration

import "testing"

func TestRepeating(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}