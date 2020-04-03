package orm

import (
	"reflect"
	"testing"
)

func TestReflectSequence1(t *testing.T) {
	_, _, got := ReflectSequence(nil, "")
	want := "t can't be nil"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	type User struct {
	}

	_, _, got2 := ReflectSequence(reflect.TypeOf(User{}), "")
	want2 := "tag can't be empty"
	if got2 != nil && got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	_, _, got3 := ReflectSequence(reflect.TypeOf(&User{}).Elem(), "")
	want3 := "tag can't be empty"
	if got3 != nil && got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	_, _, got4 := ReflectSequence(reflect.TypeOf(User{}), "db")
	want4 := "t's fields can't be empty"
	if got4 != nil && got4.Error() != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	_, _, got5 := ReflectSequence(reflect.TypeOf(&User{}).Elem(), "db")
	want5 := "t's fields can't be empty"
	if got5 != nil && got5.Error() != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}
}

func TestReflectSequence2(t *testing.T) {
	type User struct {
		Id    int
		Name  string
		Age   int
		Phone string
	}

	_, _, got := ReflectSequence(reflect.TypeOf(User{}), "db")
	if !IsEmptyTag(got) {
		t.Errorf("got %q; want %q", got, EmptyTag)
	}

	_, _, got2 := ReflectSequence(reflect.TypeOf(&User{}).Elem(), "db")
	if !IsEmptyTag(got2) {
		t.Errorf("got %q; want %q", got2, EmptyTag)
	}

	type User2 struct {
		Id    int    `table:"id"`
		Name  string `table:"id"`
		Age   int
		Phone string
	}

	_, _, got3 := ReflectSequence(reflect.TypeOf(User2{}), "db")
	if !IsEmptyTag(got3) {
		t.Errorf("got %q; want %q", got3, EmptyTag)
	}

	_, _, got4 := ReflectSequence(reflect.TypeOf(&User2{}).Elem(), "db")
	if !IsEmptyTag(got4) {
		t.Errorf("got %q; want %q", got4, EmptyTag)
	}
}

func TestReflectSequence3(t *testing.T) {
	type User struct {
		Id    int    `db:"id"`
		Name  string `db:"id"`
		age   int
		Phone string `db:"phone"`
	}

	_, _, got := ReflectSequence(reflect.TypeOf(User{}), "db")
	want := "tag 'id' duplicate"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	type User2 struct {
		Id    int    `db:"id"`
		Name  string `db:"phone"`
		age   int
		Phone string `db:"phone"`
	}

	_, _, got2 := ReflectSequence(reflect.TypeOf(&User2{}).Elem(), "db")
	want2 := "tag 'phone' duplicate"
	if got2 != nil && got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}
}

func TestReflectSequence4(t *testing.T) {
	type User struct {
		Id    int    `db:"Id"`
		Name  string `db:"name"`
		Age   int
		phone string `db:"Phone"`
	}

	columns, sequence, _ := ReflectSequence(reflect.TypeOf(User{}), "db")

	if len(columns) != 3 {
		t.Errorf("got len %d; want 3", len(columns))
	}

	if len(sequence) != 3 {
		t.Errorf("got len %d; want 3", len(sequence))
	}

	if columns[0] != "Id" || columns[1] != "name" || columns[2] != "Phone" {
		t.Errorf("got %v; want [Id name Phone]", columns)
	}

	if sequence["Id"] != 0 || sequence["name"] != 1 || sequence["Phone"] != 3 {
		t.Errorf("got %v; want map[Id:0 name:1 Phone:3]", sequence)
	}

	columns2, sequence2, _ := ReflectSequence(reflect.TypeOf(&User{}).Elem(), "db")

	if len(columns2) != 3 {
		t.Errorf("got len %d; want 3", len(columns2))
	}

	if len(sequence2) != 3 {
		t.Errorf("got len %d; want 3", len(sequence2))
	}

	if columns2[0] != "Id" || columns2[1] != "name" || columns2[2] != "Phone" {
		t.Errorf("got %v; want [Id name Phone]", columns2)
	}

	if sequence2["Id"] != 0 || sequence2["name"] != 1 || sequence2["Phone"] != 3 {
		t.Errorf("got %v; want map[Id:0 name:1 Phone:3]", sequence2)
	}

	columns3, sequence3, _ := ReflectSequence(reflect.TypeOf(User{Id: 1001, Name: "abc", Age: 17, phone: "13000000001"}), "db")

	if len(columns3) != 3 {
		t.Errorf("got len %d; want 3", len(columns3))
	}

	if len(sequence3) != 3 {
		t.Errorf("got len %d; want 3", len(sequence3))
	}

	if columns3[0] != "Id" || columns3[1] != "name" || columns3[2] != "Phone" {
		t.Errorf("got %v; want [Id name Phone]", columns3)
	}

	if sequence3["Id"] != 0 || sequence3["name"] != 1 || sequence3["Phone"] != 3 {
		t.Errorf("got %v; want map[Id:0 name:1 Phone:3]", sequence3)
	}

	columns4, sequence4, _ := ReflectSequence(reflect.TypeOf(&User{Id: 1001, Name: "abc", Age: 17, phone: "13000000001"}).Elem(), "db")

	if len(columns4) != 3 {
		t.Errorf("got len %d; want 3", len(columns4))
	}

	if len(sequence4) != 3 {
		t.Errorf("got len %d; want 3", len(sequence4))
	}

	if columns4[0] != "Id" || columns4[1] != "name" || columns4[2] != "Phone" {
		t.Errorf("got %v; want [Id name Phone]", columns4)
	}

	if sequence4["Id"] != 0 || sequence4["name"] != 1 || sequence4["Phone"] != 3 {
		t.Errorf("got %v; want map[Id:0 name:1 Phone:3]", sequence4)
	}
}
