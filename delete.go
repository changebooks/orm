package orm

import (
	"database/sql"
	"errors"
)

// id is primary key's value
func Delete(db *sql.DB, table string, id interface{}) (affectedRows int64, err error, affectedRowsErr error, query string) {
	if db == nil {
		err = errors.New("db can't be nil")
		return
	}

	if table == "" {
		err = errors.New("table can't be empty")
		return
	}

	if id == nil {
		err = errors.New("id can't be nil")
		return
	}

	query, err = BuildDelete(table)
	if err != nil {
		return
	}

	query = ApplyPrimaryKey(query)

	result, err := db.Exec(query, id)
	if err != nil {
		return
	}

	affectedRows, affectedRowsErr = result.RowsAffected()
	return
}
