package orm

import (
	"errors"
	"fmt"
)

func BuildDelete(table string) (string, error) {
	if table == "" {
		return "", errors.New("table can't be empty")
	}

	return fmt.Sprintf("DELETE FROM %s", table), nil
}
