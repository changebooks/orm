package orm

import (
	"errors"
	"fmt"
	"strings"
)

func BuildSelect(table string, columns []string) (string, error) {
	if table == "" {
		return "", errors.New("table can't be empty")
	}

	if columns == nil {
		return "", errors.New("columns can't be nil")
	}

	if len(columns) == 0 {
		return "", errors.New("columns can't be empty")
	}

	fields := strings.Join(columns, ", ")

	return fmt.Sprintf("SELECT %s FROM %s", fields, table), nil
}
