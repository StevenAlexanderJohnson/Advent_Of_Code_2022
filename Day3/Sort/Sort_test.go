package Sort

import (
	"testing"
)

func TestInsertion_Sort(t *testing.T) {
	got := Insertion_Sort([]byte("abcd"))
	want := []byte("dcba")

	for i, _ := range got {
		if got[i] != want[i] {
			t.Errorf("got %q want %q", got, want)
		}
	}
}
