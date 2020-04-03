package orm

import (
	"database/sql"
	"errors"
)

// attributes is a struct or a pointer of struct (eg. struct or &struct)
// id is primary key's value
func Update(db *sql.DB, table string, attributes map[string]interface{}, id interface{}) (affectedRows int64, err error, affectedRowsErr error, query string, args []interface{}) {
	if db == nil {
		err = errors.New("db can't be nil")
		return
	}

	if table == "" {
		err = errors.New("table can't be empty")
		return
	}

	if attributes == nil {
		err = errors.New("attributes can't be nil")
		return
	}

	if id == nil {
		err = errors.New("id can't be nil")
		return
	}

	columns, args, err := MapValue(attributes)
	if err != nil {
		return
	}

	query, err = BuildUpdate(table, columns)
	if err != nil {
		return
	}

	query = ApplyPrimaryKey(query)
	args = append(args, id)

	result, err := db.Exec(query, args...)
	if err != nil {
		return
	}

	affectedRows, affectedRowsErr = result.RowsAffected()
	return
}
