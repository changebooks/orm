package orm

import (
	"database/sql"
	"errors"
)

// attributes is a struct or a pointer of struct (eg. struct or &struct)
func Insert(db *sql.DB, table string, attributes interface{}) (result sql.Result, err error, query string, args []interface{}) {
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

	columns, args, _, err := ReflectValue(attributes, TagKey)
	if err != nil {
		return
	}

	query, err = BuildInsert(table, columns)
	if err != nil {
		return
	}

	result, err = db.Exec(query, args...)
	return
}
