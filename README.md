# orm
Object Relational Mapping
==

<pre>
CREATE TABLE user  (
  id bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL DEFAULT '' COMMENT '姓名',
  age int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '年龄',
  phone varchar(32) NOT NULL DEFAULT '' COMMENT '手机号',
  PRIMARY KEY (id)
) ENGINE=InnoDB CHARACTER SET=utf8mb4;

type User struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Age   int    `db:"age"`
	Phone string `db:"phone"`
}
</pre>

<pre>
db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
if err != nil {
    fmt.Println(err)
    return
}

var result []*User
affectedRows, err2, closeErr := orm.Find(db, &result, "SELECT id, name, phone FROM user")
if err2 != nil {
    fmt.Println(err2)
    return
}

if closeErr != nil {
    fmt.Println(closeErr)
    return
}

fmt.Println("affectedRows: ", affectedRows)

for _, user := range result {
    fmt.Println(user)
}

_ = db.Close()
</pre>

<pre>
db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
if err != nil {
    fmt.Println(err)
    return
}

var result User
affectedRows, err2, closeErr := orm.First(db, &result, "SELECT id, name, phone FROM user")
if err2 != nil {
    fmt.Println(err2)
    return
}

if closeErr != nil {
    fmt.Println(closeErr)
    return
}

fmt.Println("affectedRows: ", affectedRows)
fmt.Println(result)

_ = db.Close()
</pre>

<pre>
db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
if err != nil {
    fmt.Println(err)
    return
}

user := &User{
    Name: "张三",
    Age: 17,
    Phone: "13000000001",
}
result, err2, query, args := orm.Insert(db, "user", user)
if err2 != nil {
    fmt.Println(err2)
    return
}

affectedRows, _ :=  result.RowsAffected()
lastInsertId, _ :=  result.LastInsertId()
fmt.Println("affectedRows: ", affectedRows)
fmt.Println("lastInsertId: ", lastInsertId)
fmt.Println("query: ", query)
fmt.Println("args: ", args)

_ = db.Close()
</pre>

<pre>
db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
if err != nil {
    fmt.Println(err)
    return
}

affectedRows, err2, affectedRowsErr, query, args := orm.Update(db, "user", map[string]interface{}{"age": 18}, 2)
if err2 != nil {
    fmt.Println(err2)
    return
}

if affectedRowsErr != nil {
    fmt.Println(affectedRowsErr)
    return
}
	
fmt.Println("affectedRows: ", affectedRows)
fmt.Println("query: ", query)
fmt.Println("args: ", args)

_ = db.Close()
</pre>

<pre>
db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
if err != nil {
    fmt.Println(err)
    return
}

affectedRows, err2, affectedRowsErr, query := orm.Delete(db, "user", 1)
if err2 != nil {
    fmt.Println(err2)
    return
}

if affectedRowsErr != nil {
    fmt.Println(affectedRowsErr)
    return
}

fmt.Println("affectedRows: ", affectedRows)
fmt.Println("query: ", query)

_ = db.Close()
</pre>
