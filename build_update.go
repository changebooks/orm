package orm

import (
	"errors"
	"fmt"
	"strings"
)

func BuildUpdate(table string, columns []string) (string, error) {
	if table == "" {
		return "", errors.New("table can't be empty")
	}

	if columns == nil {
		return "", errors.New("columns can't be nil")
	}

	if len(columns) == 0 {
		return "", errors.New("columns can't be empty")
	}

	fields := strings.Join(columns, " = "+PlaceHolder+", ") + " = " + PlaceHolder

	return fmt.Sprintf("UPDATE %s SET %s", table, fields), nil
}
