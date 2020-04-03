package orm

import (
	"reflect"
	"testing"
)

func TestSqlSequence(t *testing.T) {
	type User struct {
		Id    int    `db:"id"`
		Name  string `db:"name"`
		Age   int
		Phone string `db:"phone"`
	}

	tag, _ := NewTag(reflect.TypeOf(User{}))

	_, got := tag.SqlSequence(nil)
	want := "columns can't be nil"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	_, got2 := tag.SqlSequence([]string{})
	want2 := "columns's size can't be less or equal than 0"
	if got2 != nil && got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	sequence, _ := tag.SqlSequence([]string{"id", "phone"})
	if len(sequence) != 2 {
		t.Errorf("got len %d; want 2", len(sequence))
	}

	if sequence["id"] != 0 || sequence["phone"] != 3 {
		t.Errorf("got %v; want map[id:0 phone:3]", sequence)
	}
}
