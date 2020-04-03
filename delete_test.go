package orm

import (
	"database/sql"
	"testing"
)

func TestDelete(t *testing.T) {
	_, got, _, _ := Delete(nil, "", nil)
	want := "db can't be nil"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	db := &sql.DB{}

	_, got2, _, _ := Delete(db, "", nil)
	want2 := "table can't be empty"
	if got2 != nil && got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	_, got3, _, _ := Delete(db, "user", nil)
	want3 := "id can't be nil"
	if got3 != nil && got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}
}
