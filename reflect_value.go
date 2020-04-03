package orm

import (
	"errors"
	"fmt"
	"reflect"
)

// s is a struct or a pointer of struct (eg. struct or &struct)
func ReflectValue(s interface{}, tag string) (columns []string, values []interface{}, attributes map[string]interface{}, err error) {
	if s == nil {
		return nil, nil, nil, errors.New("s can't be nil")
	}

	if tag == "" {
		return nil, nil, nil, errors.New("tag can't be empty")
	}

	isPtr := false // struct
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		isPtr = true // &struct
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, nil, nil, errors.New("s must be a struct or a pointer of struct (eg. struct or &struct)")
	}

	size := t.NumField()
	if size <= 0 {
		return nil, nil, nil, errors.New("s's fields can't be empty")
	}

	v := reflect.ValueOf(s)
	if isPtr {
		v = v.Elem()
	}

	attributes = make(map[string]interface{})
	for i := 0; i < size; i++ {
		f := v.Field(i)
		if f.CanInterface() {
			n := t.Field(i).Tag.Get(tag)
			if n == "" {
				continue
			}

			if _, ok := attributes[n]; ok {
				return nil, nil, nil, errors.New("tag '" + n + "' duplicate")
			}

			o := f.Interface()

			columns = append(columns, n)
			values = append(values, o)
			attributes[n] = o
		}
	}

	if columns == nil || values == nil {
		return nil, nil, nil, NewEmptyTag()
	}

	columnSize := len(columns)
	valueSize := len(values)
	attributeSize := len(attributes)

	if columnSize == 0 || valueSize == 0 || attributeSize == 0 {
		return nil, nil, nil, NewEmptyTag()
	}

	if columnSize != valueSize {
		return nil, nil, nil, fmt.Errorf("columns's size %d must be equal than values's size %d", columnSize, valueSize)
	}

	if columnSize != attributeSize {
		return nil, nil, nil, fmt.Errorf("columns's size %d must be equal than attributes's size %d", columnSize, attributeSize)
	}

	return columns, values, attributes, nil
}
