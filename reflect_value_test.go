package orm

import "testing"

func TestReflectValue1(t *testing.T) {
	_, _, _, got := ReflectValue(nil, "")
	want := "s can't be nil"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	_, _, _, got2 := ReflectValue(1, "")
	want2 := "tag can't be empty"
	if got2 != nil && got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	_, _, _, got3 := ReflectValue(1, "db")
	want3 := "s must be a struct or a pointer of struct (eg. struct or &struct)"
	if got3 != nil && got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	_, _, _, got4 := ReflectValue("", "db")
	want4 := "s must be a struct or a pointer of struct (eg. struct or &struct)"
	if got4 != nil && got4.Error() != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	_, _, _, got5 := ReflectValue(map[int]string{}, "db")
	want5 := "s must be a struct or a pointer of struct (eg. struct or &struct)"
	if got5 != nil && got5.Error() != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}

	type User struct {
	}

	_, _, _, got6 := ReflectValue(User{}, "db")
	want6 := "s's fields can't be empty"
	if got6 != nil && got6.Error() != want6 {
		t.Errorf("got %q; want %q", got6, want6)
	}

	_, _, _, got7 := ReflectValue(&User{}, "db")
	want7 := "s's fields can't be empty"
	if got7 != nil && got7.Error() != want7 {
		t.Errorf("got %q; want %q", got7, want7)
	}
}

func TestReflectValue2(t *testing.T) {
	type User struct {
		Id    int
		Name  string
		Age   int
		Phone string
	}

	_, _, _, got := ReflectValue(User{}, "db")
	if !IsEmptyTag(got) {
		t.Errorf("got %q; want %q", got, EmptyTag)
	}

	_, _, _, got2 := ReflectValue(&User{}, "db")
	if !IsEmptyTag(got2) {
		t.Errorf("got %q; want %q", got2, EmptyTag)
	}

	type User2 struct {
		Id    int    `table:"id"`
		Name  string `table:"id"`
		Age   int
		Phone string
	}

	_, _, _, got3 := ReflectValue(User2{}, "db")
	if !IsEmptyTag(got3) {
		t.Errorf("got %q; want %q", got3, EmptyTag)
	}

	_, _, _, got4 := ReflectValue(&User2{}, "db")
	if !IsEmptyTag(got4) {
		t.Errorf("got %q; want %q", got4, EmptyTag)
	}
}

func TestReflectValue3(t *testing.T) {
	type User struct {
		Id    int    `db:"id"`
		Name  string `db:"id"`
		age   int
		Phone string `db:"Phone"`
	}

	_, _, _, got := ReflectValue(User{}, "db")
	want := "tag 'id' duplicate"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	type User2 struct {
		Id    int    `db:"id"`
		Name  string `db:"name"`
		age   int
		Phone string `db:"name"`
	}

	_, _, _, got2 := ReflectValue(&User2{}, "db")
	want2 := "tag 'name' duplicate"
	if got2 != nil && got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}
}

func TestReflectValue4(t *testing.T) {
	type User struct {
		Id    int    `db:"Id"`
		Name  string `db:"name"`
		Age   int
		Phone string `db:"Phone"`
	}

	columns, values, attributes, _ := ReflectValue(User{}, "db")

	if len(columns) != 3 {
		t.Errorf("got len %d; want 3", len(columns))
	}

	if len(values) != 3 {
		t.Errorf("got len %d; want 3", len(values))
	}

	if len(attributes) != 3 {
		t.Errorf("got len %d; want 3", len(attributes))
	}

	if columns[0] != "Id" || columns[1] != "name" || columns[2] != "Phone" {
		t.Errorf("got %v; want [Id name Phone]", columns)
	}

	if values[0] != 0 || values[1] != "" || values[2] != "" {
		t.Errorf("got %v; want [0  ]", values)
	}

	if attributes["Id"] != 0 || attributes["name"] != "" || attributes["Phone"] != "" {
		t.Errorf("got %v; want map[Id:0 Name: Phone:]", attributes)
	}

	columns2, values2, attributes2, _ := ReflectValue(&User{}, "db")

	if len(columns2) != 3 {
		t.Errorf("got len %d; want 3", len(columns2))
	}

	if len(values2) != 3 {
		t.Errorf("got len %d; want 3", len(values2))
	}

	if len(attributes2) != 3 {
		t.Errorf("got len %d; want 3", len(attributes2))
	}

	if columns2[0] != "Id" || columns2[1] != "name" || columns2[2] != "Phone" {
		t.Errorf("got %v; want [Id name Phone]", columns2)
	}

	if values2[0] != 0 || values2[1] != "" || values2[2] != "" {
		t.Errorf("got %v; want map[0  ]", values2)
	}

	if attributes2["Id"] != 0 || attributes2["name"] != "" || attributes2["Phone"] != "" {
		t.Errorf("got %v; want map[Id:0 name: Phone:]", attributes2)
	}

	columns3, values3, attributes3, _ := ReflectValue(User{Id: 1001, Name: "abc", Age: 17, Phone: "13000000001"}, "db")

	if len(columns3) != 3 {
		t.Errorf("got len %d; want 3", len(columns3))
	}

	if len(values3) != 3 {
		t.Errorf("got len %d; want 3", len(values3))
	}

	if len(attributes3) != 3 {
		t.Errorf("got len %d; want 3", len(attributes3))
	}

	if columns3[0] != "Id" || columns3[1] != "name" || columns3[2] != "Phone" {
		t.Errorf("got %v; want [Id name Phone]", columns3)
	}

	if values3[0] != 1001 || values3[1] != "abc" || values3[2] != "13000000001" {
		t.Errorf("got %v; want [1001 abc 13000000001]", values3)
	}

	if attributes3["Id"] != 1001 || attributes3["name"] != "abc" || attributes3["Phone"] != "13000000001" {
		t.Errorf("got %v; want map[Phone:13000000001 Id:1001 name:abc]", attributes3)
	}

	columns4, values4, attributes4, _ := ReflectValue(&User{Id: 1001, Name: "abc", Age: 17, Phone: "13000000001"}, "db")

	if len(columns4) != 3 {
		t.Errorf("got len %d; want 3", len(columns4))
	}

	if len(values4) != 3 {
		t.Errorf("got len %d; want 3", len(values4))
	}

	if len(attributes4) != 3 {
		t.Errorf("got len %d; want 3", len(attributes4))
	}

	if columns4[0] != "Id" || columns4[1] != "name" || columns4[2] != "Phone" {
		t.Errorf("got %v; want [Id name Phone]", columns4)
	}

	if values4[0] != 1001 || values4[1] != "abc" || values4[2] != "13000000001" {
		t.Errorf("got %v; want [1001 abc 13000000001]", values4)
	}

	if attributes4["Id"] != 1001 || attributes4["name"] != "abc" || attributes4["Phone"] != "13000000001" {
		t.Errorf("got %v; want map[Phone:13000000001 Id:1001 name:abc]", attributes4)
	}
}

func TestReflectValue5(t *testing.T) {
	type User struct {
		Id    int    `db:"id"`
		name  string `db:"name"`
		age   int
		phone string `db:"phone"`
	}

	columns, values, attributes, _ := ReflectValue(User{}, "db")

	if len(columns) != 1 {
		t.Errorf("got len %d; want 1", len(columns))
	}

	if len(values) != 1 {
		t.Errorf("got len %d; want 1", len(values))
	}

	if len(attributes) != 1 {
		t.Errorf("got len %d; want 1", len(attributes))
	}

	if columns[0] != "id" {
		t.Errorf("got %v; want [id]", columns)
	}

	if values[0] != 0 {
		t.Errorf("got %v; want [0]", values)
	}

	if attributes["id"] != 0 {
		t.Errorf("got %v; want map[id:0]", attributes)
	}

	columns2, values2, attributes2, _ := ReflectValue(&User{}, "db")

	if len(columns2) != 1 {
		t.Errorf("got len %d; want 1", len(columns2))
	}

	if len(values2) != 1 {
		t.Errorf("got len %d; want 1", len(values2))
	}

	if len(attributes2) != 1 {
		t.Errorf("got len %d; want 1", len(attributes2))
	}

	if columns2[0] != "id" {
		t.Errorf("got %v; want [id]", columns2)
	}

	if values2[0] != 0 {
		t.Errorf("got %v; want [0]", values2)
	}

	if attributes2["id"] != 0 {
		t.Errorf("got %v; want map[id:0]", attributes2)
	}

	columns3, values3, attributes3, _ := ReflectValue(User{Id: 1001, name: "abc", age: 17, phone: "13000000001"}, "db")

	if len(columns3) != 1 {
		t.Errorf("got len %d; want 1", len(columns3))
	}

	if len(values3) != 1 {
		t.Errorf("got len %d; want 1", len(values3))
	}

	if len(attributes3) != 1 {
		t.Errorf("got len %d; want 1", len(attributes3))
	}

	if columns3[0] != "id" {
		t.Errorf("got %v; want [id]", columns3)
	}

	if values3[0] != 1001 {
		t.Errorf("got %v; want [1001]", values3)
	}

	if attributes3["id"] != 1001 {
		t.Errorf("got %v; want map[id:1001]", attributes3)
	}

	columns4, values4, attributes4, _ := ReflectValue(&User{Id: 1001, name: "abc", age: 17, phone: "13000000001"}, "db")

	if len(columns4) != 1 {
		t.Errorf("got len %d; want 1", len(columns4))
	}

	if len(values4) != 1 {
		t.Errorf("got len %d; want 1", len(values4))
	}

	if len(attributes4) != 1 {
		t.Errorf("got len %d; want 1", len(attributes4))
	}

	if columns4[0] != "id" {
		t.Errorf("got %v; want [id]", columns4)
	}

	if values4[0] != 1001 {
		t.Errorf("got %v; want [1001]", values4)
	}

	if attributes4["id"] != 1001 {
		t.Errorf("got %v; want map[id:1001]", attributes4)
	}
}
