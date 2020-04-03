package orm

import (
	"database/sql"
	"errors"
	"reflect"
)

// result is a pointer (eg. &[]struct or &[]*struct)
// return affected rows
func Scan(rows *sql.Rows, result interface{}) (int64, error) {
	if rows == nil {
		return 0, errors.New("rows can't be nil")
	}

	if result == nil {
		return 0, errors.New("result can't be nil")
	}

	out := reflect.ValueOf(result)
	if out.Kind() != reflect.Ptr {
		return 0, errors.New("result must be a pointer (eg. &[]struct or &[]*struct)")
	}

	out = out.Elem()
	if out.Kind() != reflect.Slice {
		return 0, errors.New("result must be a slice (eg. &[]struct or &[]*struct)")
	}

	outType := out.Type()
	outType = outType.Elem()
	isPtr := false // &[]struct
	if outType.Kind() == reflect.Ptr {
		isPtr = true // &[]*struct
		outType = outType.Elem()
	}

	columns, err := rows.Columns()
	if err != nil {
		return 0, err
	}

	size := len(columns)
	if size == 0 {
		return 0, errors.New("columns's size can't be equal than 0")
	}

	sequence, err := SqlSequence(columns, outType)
	if err != nil {
		return 0, err
	}

	if len(sequence) == 0 {
		return 0, errors.New("no sequence, sql rows's columns no mapping result's tag")
	}

	var affectedRows int64 = 0
	for rows.Next() {
		attributes := make([]interface{}, size)
		elements := reflect.New(outType).Elem()

		for i, n := range columns {
			if j, ok := sequence[n]; ok {
				f := elements.Field(j)
				if f.IsValid() {
					attributes[i] = f.Addr().Interface()
					continue
				}
			}

			attributes[i] = new(interface{})
		}

		if err := rows.Scan(attributes...); err != nil {
			return affectedRows, err
		}

		if isPtr {
			out.Set(reflect.Append(out, elements.Addr()))
		} else {
			out.Set(reflect.Append(out, elements))
		}

		affectedRows++
	}

	return affectedRows, nil
}
