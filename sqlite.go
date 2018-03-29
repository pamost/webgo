package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)
// http://tdm-gcc.tdragon.net/download GCC for WIN64

//CREATE TABLE `userinfo` (
//`uid` INTEGER PRIMARY KEY AUTOINCREMENT,
//`username` VARCHAR(64) NULL,
//`email` VARCHAR(64) NULL,
//`created` DATE NULL
//);

func main() {

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	checkErr(err)

	// вставка
	stmt, err := db.Prepare("INSERT INTO userinfo(username, email, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("pavel", "pavel@pavel.ru", "2018-01-03")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	// обновление
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("pavelupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// запрос
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var email string
		var created string
		err = rows.Scan(&uid, &username, &email, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(email)
		fmt.Println(created)
	}

	// удаление
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
