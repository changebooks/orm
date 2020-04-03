package orm

import (
	"database/sql"
	"errors"
)

// result is a pointer of struct (eg. var result User; &result)
func First(db *sql.DB, result interface{}, query string, args ...interface{}) (affectedRows int64, err error, closeErr error) {
	if db == nil {
		err = errors.New("db can't be nil")
		return
	}

	rows, err := db.Query(query, args...)
	if err == nil {
		affectedRows, err = ScanFirst(rows, result)
	}

	if rows != nil {
		closeErr = rows.Close()
	}

	return
}
