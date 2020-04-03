package orm

import (
	"errors"
	"fmt"
	"strings"
)

func BuildInsert(table string, columns []string) (string, error) {
	if table == "" {
		return "", errors.New("table can't be empty")
	}

	if columns == nil {
		return "", errors.New("columns can't be nil")
	}

	size := len(columns)
	if size == 0 {
		return "", errors.New("columns can't be empty")
	}

	fields := strings.Join(columns, ", ")
	values := strings.TrimRight(strings.Repeat(PlaceHolder+", ", size), ", ")

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, fields, values), nil
}
