package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

var db *sql.DB

func initDB() (err error) {
	dsn := "homestead:secret@tcp(192.168.10.10:3306)/homestead"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		defer db.Close()
		return
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)
	return nil
}

type User struct {
	Id   int
	Name string
	Age  int
}

//查询单条语句
func queryOne(sql string, args ...interface{}) {
	u := &User{}
	err := db.QueryRow(sql, args...).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", u)
}

//查询多条语句
func queryMore(sql string, args ...interface{}) {
	allUser := make([]User, 0, 10)
	u := User{}
	rows, err := db.Query(sql, args...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&u.Id, &u.Name, &u.Age)
		allUser = append(allUser, u)
	}
	fmt.Println("--------------")
	fmt.Printf("%#v", allUser)
}

//插入sql语句
func insertSql(sql string, args ...interface{}) {
	r, err := db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.LastInsertId())
}

func updateSql(sql string, args ...interface{}) (n int64, err error) {
	r, err := db.Exec(sql, args...)
	if err != nil {
		return 0, nil
	}
	n, err = r.RowsAffected()
	fmt.Println(r.RowsAffected())
	return n, nil
}

func updateSqlByTx(tx *sql.Tx, sql string, args ...interface{}) (n int64, err error) {
	r, err := tx.Exec(sql, args...)
	if err != nil {
		return 0, nil
	}
	n, err = r.RowsAffected()
	fmt.Println(r.RowsAffected())
	return n, nil
}

func deleteSql(sql string, args ...interface{}) {
	r, err := db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.RowsAffected())
}

func prepareSql(sql string, args ...interface{}) {
	allUser := make([]User, 0)
	u := User{}
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rows, _ := stmt.Query(args...)
	for rows.Next() {
		err = rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			panic(err)
		}
		allUser = append(allUser, u)
	}
	fmt.Printf("%#v", allUser)
}

func transactionSql() {
	tx, err := db.Begin()
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

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	//sql := "select * from user where age>?"
	//queryOne(sql, 15)
	//queryMore(sql, 0)

	//sql := "INSERT INTO `user`(`name`,`age`) VALUES (?,?)"
	//insertSql(sql, "why", 5)

	//sql := "update user set age=? where name=?"
	//updateSql(sql,7,"why")

	//sql := "delete from user where id=?"
	//deleteSql(sql, 1)

	//sql := "select * from user where age>?"
	//prepareSql(sql,1)

	transactionSql()
}
