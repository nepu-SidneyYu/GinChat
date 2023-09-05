package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	ID    int64  //`gorm:"column:id"`
	Name  string //`gorm:"column:name"`
	Phone string //`gorm:"column:phone"`
}

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "123456",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "test",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	//users := []*User{
	//	&User{ID: 9385, Name: "sjgs", Phone: "hjsdjf"},
	//	&User{ID: 8752, Name: "dskmf", Phone: "skgs"},
	//}
	//db.Exec("insert into user values (9385,'sjgs','hjsdjf')")
	//db.Exec("insert into user values (8752,'dskmf','skgs')")
	var user User
	//deletedb, err := db.Prepare("delete from user where id=?")
	//deletedb.Exec(9385)
	//deletedb.Exec(8752)
	rows, e := db.Query("select * from user")
	if e != nil {
		errors.New("query incur error")
	}
	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Phone)
		fmt.Printf("%#v\n", user)
	}
	rows.Close()
}
