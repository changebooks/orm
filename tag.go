package orm

import (
	"errors"
	"reflect"
)

// struct tag attributes (tag: define.Tag)
// type User struct {
//	 Id    int    `db:"id"`
//	 Phone string
//	 Name  string `db:"name"`
// }
type Tag struct {
	columns  []string       // column name list (eg. [id name])
	sequence map[string]int // column name => column sequence num (eg. [id:0 name:2])
}

// t is the reflect.Type (eg. reflect.TypeOf(User{}))
// or
// t is the pointer's reflect.Type's element type (eg. reflect.TypeOf(&User{}).Elem())
//
// type User struct {
//   Id   int    `db:"id"`
//   Name string `db:"name"`
// }
//
// t := reflect.TypeOf(User{})
// or
// t := reflect.TypeOf(&User{}).Elem()
func NewTag(t reflect.Type) (*Tag, error) {
	columns, sequence, err := ReflectSequence(t, TagKey)
	if err != nil {
		return nil, err
	}

	return &Tag{
		columns:  columns,
		sequence: sequence,
	}, nil
}

// columns is the sql.Rows's columns name
// return column name => column sequence num (eg. [id:0 name:2])
func (x *Tag) SqlSequence(columns []string) (map[string]int, error) {
	if columns == nil {
		return nil, errors.New("columns can't be nil")
	}

	if len(columns) == 0 {
		return nil, errors.New("columns's size can't be equal than 0")
	}

	r := make(map[string]int)
	for _, n := range columns {
		if n == "" {
			continue
		}

		if i, ok := x.sequence[n]; ok {
			r[n] = i
		}
	}

	return r, nil
}

func (x *Tag) GetColumns() []string {
	return x.columns
}

func (x *Tag) GetSequence() map[string]int {
	return x.sequence
}
