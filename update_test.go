package orm

import (
	"database/sql"
	"testing"
)

func TestUpdate(t *testing.T) {
	_, got, _, _, _ := Update(nil, "", nil, nil)
	want := "db can't be nil"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	db := &sql.DB{}

	_, got2, _, _, _ := Update(db, "", nil, nil)
	want2 := "table can't be empty"
	if got2 != nil && got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	_, got3, _, _, _ := Update(db, "user", nil, nil)
	want3 := "attributes can't be nil"
	if got3 != nil && got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	_, got4, _, _, _ := Update(db, "user", map[string]interface{}{}, nil)
	want4 := "id can't be nil"
	if got4 != nil && got4.Error() != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	_, got5, _, _, _ := Update(db, "user", map[string]interface{}{}, 1)
	want5 := "attributes can't be empty"
	if got5 != nil && got5.Error() != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}
}
