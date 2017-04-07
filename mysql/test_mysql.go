//http://www.cnblogs.com/tsiangleo/p/4483657.html
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

var db = &sql.DB{}

func main() {
	Init()
	Insert()
	Update()
	Query()
	//	Delete()
}

func Init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/gc")
	if err != nil {
		panic(err.Error())
	}
}

func Insert() {
	start := time.Now()
	for i := 1001; i <= 1010; i++ {
		db.Exec("insert into book(id, title, pages) values(?,?,?)", i, strconv.Itoa(i), ((i * 243) % 200))
	}
	end := time.Now()
	fmt.Println("insert1 total time : ", end.Sub(start).Seconds())
}

func Delete() {
	start := time.Now()
	for i := 1001; i <= 1010; i++ {
		db.Exec("delete from book where id = ?", i)
	}
	end := time.Now()
	fmt.Println("Delete1 total time : ", end.Sub(start).Seconds())
}

func Update() {
	start := time.Now()
	for i := 1001; i <= 1010; i++ {
		//
		db.Exec("Update book set title = ? where id = ?", i+100, i)
	}
	end := time.Now()
	fmt.Println("Update1 total time : ", end.Sub(start).Seconds())

	start = time.Now()
	stm, _ := db.Prepare("update book set title=? where id =?")
	for i := 1001; i <= 1010; i++ {
		stm.Exec(i-1000, i)
	}
	end = time.Now()
	fmt.Println("Update2 total time : ", end.Sub(start).Seconds())

	start = time.Now()
	tx, _ := db.Begin()
	for i := 1001; i <= 1010; i++ {
		tx.Exec("update book set title=? where id=?", i-123, i)
	}
	tx.Commit() // commit 1 times //quick
	end = time.Now()
	fmt.Println("Update3 total time : ", end.Sub(start).Seconds())
}

func Query() {
	start := time.Now()
	rows, _ := db.Query("select * from book")
	defer rows.Close()
	end := time.Now()
	fmt.Println("Query1 total time : ", end.Sub(start).Seconds())
}
