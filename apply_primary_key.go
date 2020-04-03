package orm

func ApplyPrimaryKey(query string) string {
	return query + " WHERE " + PrimaryKey + " = ?"
}
