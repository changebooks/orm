package orm

import (
	"errors"
	"reflect"
	"sync"
)

type Tags struct {
	mu   sync.Mutex            // protects following fields
	data map[reflect.Type]*Tag // reflect.Type => Tag
}

func NewTags() *Tags {
	return &Tags{
		data: make(map[reflect.Type]*Tag),
	}
}

func (x *Tags) Get(key reflect.Type) (*Tag, error) {
	if key == nil {
		return nil, errors.New("key can't be nil")
	}

	if r, ok := x.data[key]; ok {
		return r, nil
	}

	r, err := NewTag(key)
	if err != nil {
		return nil, err
	}

	if err2 := x.Set(key, r); err2 != nil {
		return nil, err2
	}

	return r, nil
}

func (x *Tags) Set(key reflect.Type, value *Tag) error {
	if key == nil {
		return errors.New("key can't be nil")
	}

	if value == nil {
		return errors.New("value can't be nil")
	}

	if _, ok := x.data[key]; ok {
		return nil
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	if _, ok := x.data[key]; ok {
		return nil
	}

	x.data[key] = value
	return nil
}

var tags = NewTags()

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
func GetTag(t reflect.Type) (*Tag, error) {
	return tags.Get(t)
}

// columns is the sql.Rows's columns name
// return column name => column sequence num (eg. [id:0 name:2])
func SqlSequence(columns []string, t reflect.Type) (map[string]int, error) {
	tag, err := GetTag(t)
	if err != nil {
		return nil, err
	}

	return tag.SqlSequence(columns)
}
