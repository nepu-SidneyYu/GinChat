package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	ID    int64
	Name  string
	Email string
}

func main() {
	cfg := mysql.Config{
		User:                 "root",   //os.Getenv("DBUSER"),
		Passwd:               "123456", //os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "test",
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
		e := rows.Scan(user.ID, user.Name, user.Email)

		if e != nil {
			fmt.Printf("%#v\n", user)

			//fmt.Println(json.Marshal(user))
		}
	}
	rows.Close()
}
