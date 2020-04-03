package orm

import (
	"database/sql"
	"testing"
)

func TestInsert(t *testing.T) {
	_, got, _, _ := Insert(nil, "", nil)
	want := "db can't be nil"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	db := &sql.DB{}

	_, got2, _, _ := Insert(db, "", nil)
	want2 := "table can't be empty"
	if got2 != nil && got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	_, got3, _, _ := Insert(db, "user", nil)
	want3 := "attributes can't be nil"
	if got3 != nil && got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	_, got4, _, _ := Insert(db, "user", nil)
	want4 := "attributes can't be nil"
	if got4 != nil && got4.Error() != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}
}
