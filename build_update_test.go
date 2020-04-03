package orm

import "testing"

func TestBuildUpdate(t *testing.T) {
	_, got := BuildUpdate("", nil)
	want := "table can't be empty"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	_, got2 := BuildUpdate("user", nil)
	want2 := "columns can't be nil"
	if got2 != nil && got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	_, got3 := BuildUpdate("user", []string{})
	want3 := "columns can't be empty"
	if got3 != nil && got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	got4, _ := BuildUpdate("user", []string{"id"})
	want4 := "UPDATE user SET id = ?"
	if got4 != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	got5, _ := BuildUpdate("user", []string{"id", "name"})
	want5 := "UPDATE user SET id = ?, name = ?"
	if got5 != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}
}
