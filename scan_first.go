package orm

import (
	"database/sql"
	"errors"
	"reflect"
)

// result is a pointer of struct (eg. var result User; &result)
// return affected rows
func ScanFirst(rows *sql.Rows, result interface{}) (int64, error) {
	if rows == nil {
		return 0, errors.New("rows can't be nil")
	}

	if result == nil {
		return 0, errors.New("result can't be nil")
	}

	columns, err := rows.Columns()
	if err != nil {
		return 0, err
	}

	size := len(columns)
	if size == 0 {
		return 0, errors.New("columns's size can't be equal than 0")
	}

	t := reflect.TypeOf(result)
	if t.Kind() != reflect.Ptr {
		return 0, errors.New("result must be a pointer (eg. &struct)")
	}

	sequence, err := SqlSequence(columns, t.Elem())
	if err != nil {
		return 0, err
	}

	if len(sequence) == 0 {
		return 0, errors.New("no sequence, sql rows's columns no mapping result's tag")
	}

	if rows.Next() {
		attributes := make([]interface{}, size)
		elements := reflect.ValueOf(result).Elem()

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
			return 0, err
		}

		return 1, nil
	}

	return 0, nil
}
