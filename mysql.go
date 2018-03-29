package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//CREATE TABLE `world`.`userinfo` (
//`uid` INT(10) NOT NULL AUTO_INCREMENT,
//`username` VARCHAR(64) NULL DEFAULT NULL,
//`email` VARCHAR(64) NULL DEFAULT NULL,
//`created` DATE NULL DEFAULT NULL,
//PRIMARY KEY (`uid`)
//)
//ENGINE = InnoDB
//DEFAULT CHARACTER SET = utf8;

func main() {
	db, err := sql.Open("mysql", "monstr:Qwertys!23@tcp(localhost:3306)/world?charset=utf8")
	checkErr(err)

	// вставка
	stmt, err := db.Prepare("INSERT userinfo SET username=?,email=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("pavel", "pavel@pavel.ru", "2018-01-02")
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
