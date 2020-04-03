package orm

import "testing"

func TestBuildDelete(t *testing.T) {
	_, got := BuildDelete("")
	want := "table can't be empty"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2, _ := BuildDelete("user")
	want2 := "DELETE FROM user"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}
}
