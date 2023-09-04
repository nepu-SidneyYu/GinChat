package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	ID    int64  `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:phone"`
}

func main() {
	cfg := mysql.Config{
		User:                 "root",   //os.Getenv("DBUSER"),
		Passwd:               "123456", //os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "users",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	var user User
	rows, e := db.Query("select * from user")
	if e == nil {
		errors.New("query incur error")
	}
	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email)
		fmt.Printf("%#v\n", user)
	}
	rows.Close()

}
