package orm

import "testing"

func TestMapValue1(t *testing.T) {
	_, _, got := MapValue(nil)
	want := "attributes can't be nil"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	_, _, got2 := MapValue(map[string]interface{}{})
	want2 := "attributes can't be empty"
	if got2 != nil && got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	columns, values, _ := MapValue(map[string]interface{}{"id": 1})

	if len(columns) != 1 {
		t.Errorf("got len %d; want 1", len(columns))
	}

	if len(values) != 1 {
		t.Errorf("got len %d; want 1", len(values))
	}

	if columns[0] != "id" {
		t.Errorf("got %v; want [id]", columns)
	}

	if values[0] != 1 {
		t.Errorf("got %v; want [1]", values)
	}
}
