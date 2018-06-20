package main

import (
	_ "github.com/Go-SQL-Driver/mysql"
	"database/sql"
	"fmt"
	//"time"
)

//CREATE TABLE `userinfo` (
//`uid` INT(10) NOT NULL AUTO_INCREMENT,
//`username` VARCHAR(64) NULL DEFAULT NULL,
//`departname` VARCHAR(64) NULL DEFAULT NULL,
//`created` DATE NULL DEFAULT NULL,
//PRIMARY KEY (`uid`)
//);
//CREATE TABLE `userdetail` (
//`uid` INT(10) NOT NULL DEFAULT '0',
//`intro` TEXT NULL,
//`profile` TEXT NULL,
//PRIMARY KEY (`uid`)
//);


func main() {
	db, err := sql.Open("mysql", "root:emc123123@tcp(192.168.191.150:3306)/test?charset=utf8")
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("start")
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
		fmt.Println("end")
	}
	//删除数据
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
