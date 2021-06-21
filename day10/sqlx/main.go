package main

import (
	"database/sql/driver"
	"fmt"
	"github.com/jmoiron/sqlx"
)
import _ "github.com/go-sql-driver/mysql"

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

var db *sqlx.DB

func initDb() (err error) {
	dsn := "homestead:secret@(192.168.10.10)/homestead?charset=utf8mb4"
	//dsn := "root:root@(127.0.0.1)/homestead?charset=utf8mb4"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)
	return nil

}

//查询单条语句
func queryOne(sql string, args ...interface{}) {
	u := &User{}
	err := db.Get(u, sql, args...)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", u)
}

//根据搜索条件进行查询
func queryByName(sql string) {
	search := map[string]interface{}{
		"age": 0,
	}
	rows, _ := db.NamedQuery(sql, search)
	allUser := make([]User, 0)
	u := User{}
	for rows.Next() {
		rows.StructScan(&u)
		allUser = append(allUser, u)
	}
	fmt.Println(allUser)
}

//查询多条语句
func queryMore(sql string, args ...interface{}) {
	user := make([]User, 0)
	err := db.Select(&user, sql, args...)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

}

//插入sql语句
func insertSql(sql string, args ...interface{}) {
	r, err := db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.LastInsertId())
}

//更新sql语句
func updateSql(sql string, args ...interface{}) (n int64, err error) {
	r, err := db.Exec(sql, args...)
	if err != nil {
		return 0, nil
	}
	n, err = r.RowsAffected()
	fmt.Println(r.RowsAffected())
	return n, nil
}

//删除sql语句
func deleteSql(sql string, args ...interface{}) {
	r, err := db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.RowsAffected())
}

//根据查询条件执行
func execBySearch(sql string, args interface{}) {
	r, _ := db.NamedExec(sql, args)
	n, _ := r.RowsAffected()
	if n > 0 {
		fmt.Println("执行成功,影响的行数：", n)
	} else {
		fmt.Println("执行失败")
	}
}

//事务
func updateSqlByTx(tx *sqlx.Tx, sql string, args ...interface{}) (n int64, err error) {
	r, err := tx.Exec(sql, args...)
	if err != nil {
		return 0, nil
	}
	n, err = r.RowsAffected()
	fmt.Println(r.RowsAffected())
	return n, nil
}

func transactionSql() {
	tx, err := db.Beginx()
	if err != nil {
		fmt.Println("tranaction err")
	}
	sql1 := "update user set age=? where name=?"
	n, _ := updateSqlByTx(tx, sql1, 1, "why")
	if n == 0 {
		tx.Rollback()
	}

	sql2 := "update user set age=? where name=?"
	n, err = updateSqlByTx(tx, sql2, 1, "lml")
	if n == 0 {
		tx.Rollback()
	}

	tx.Commit()
}

func batchInsert(sql string, u []interface{}) {
	str, args, _ := sqlx.In(sql, u...)
	fmt.Println(sql, u, args)
	//return
	r, _ := db.Exec(str, args...)
	fmt.Println(r.RowsAffected())
}

func batchInsertByNamed(u []interface{}) {
	sql := "INSERT INTO `user`(`name`,`age`) VALUES (:name, :age)"
	r, _ := db.NamedExec(sql, u)
	fmt.Println(r.RowsAffected())
}

//根据in查询
func queryByIn(ids []int) (user []User, err error) {
	sql := "SELECT * FROM user WHERE id IN (?)"
	query, args, err := sqlx.In(sql, ids)
	query = db.Rebind(query)
	fmt.Println("批量查询sql", query)
	db.Select(&user, query, args...)
	return
}

func main() {
	err := initDb()
	if err != nil {
		panic(err)
	}
	//sql := "select * from user where age>?"
	//queryOne(sql, 2)
	//queryMore(sql, 0)

	//sql := "select * from user where age>:age"
	//queryByName(sql)

	//sql := "INSERT INTO `user`(`name`,`age`) VALUES (?,?)"
	//insertSql(sql, "wyw", 5)

	//sql := "update user set age=? where name=?"
	//updateSql(sql,7,"wyw")

	//sql := "delete from user where id=?"
	//deleteSql(sql, 5)

	//sql := "update user set age=:age where name=:name"
	//search := User{
	//	Age:19,
	//	Name:"wyw",
	//}
	//execBySearch(sql, search)

	//transactionSql()
	//u := make([]interface{}, 0)
	//sql := "INSERT INTO `user`(`name`,`age`) VALUES ";
	//for i := 1; i <= 10; i++ {
	//	name := fmt.Sprintf("N%d", i)
	//	u = append(u, &User{ Age: i, Name: name})
	//	sql += " (?), "
	//}
	// sql = strings.TrimRight(sql, ", ")
	// fmt.Println(sql)
	//batchInsert(sql, u)

	//u := make([]interface{}, 0)
	//for i := 1; i <= 10; i++ {
	//	name := fmt.Sprintf("Q%d", i)
	//	u = append(u, &User{Age: i, Name: name})
	//}
	//batchInsertByNamed(u)

	ids := []int{3, 4}
	user, _ := queryByIn(ids)
	fmt.Println(user)

}
