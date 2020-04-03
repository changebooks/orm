package orm

const (
	TagKey         = "db"        // tag: `db:"sql.Rows's column name"`
	PrimaryKey     = "id"        // table's primary key column name
	PlaceHolder    = "?"         // 填充SQL字符
	AggregateAlias = "aggregate" // 统计字段别名
	EmptyTag       = "empty tag" // struct tag empty
)
