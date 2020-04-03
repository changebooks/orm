package orm

import (
	"errors"
	"fmt"
	"reflect"
)

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
func ReflectSequence(t reflect.Type, tag string) (columns []string, sequence map[string]int, err error) {
	if t == nil {
		return nil, nil, errors.New("t can't be nil")
	}

	if tag == "" {
		return nil, nil, errors.New("tag can't be empty")
	}

	size := t.NumField()
	if size <= 0 {
		return nil, nil, errors.New("t's fields can't be empty")
	}

	sequence = make(map[string]int)
	for i := 0; i < size; i++ {
		n := t.Field(i).Tag.Get(tag)
		if n == "" {
			continue
		}

		if _, ok := sequence[n]; ok {
			return nil, nil, errors.New("tag '" + n + "' duplicate")
		}

		columns = append(columns, n)
		sequence[n] = i
	}

	if columns == nil {
		return nil, nil, NewEmptyTag()
	}

	columnSize := len(columns)
	sequenceSize := len(sequence)

	if columnSize == 0 || sequenceSize == 0 {
		return nil, nil, NewEmptyTag()
	}

	if columnSize != sequenceSize {
		return nil, nil, fmt.Errorf("columns's size %d must be equal than sequence's size %d", columnSize, sequenceSize)
	}

	return columns, sequence, nil
}
