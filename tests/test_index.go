package main

import (
	"fmt"
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint `gorm:"primaryKey;default:auto_random()"`
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/ginchat"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.UserBasic{})

	user := &models.UserBasic{}
	user.Name = "yyx"
	db.Create(user)
	db.Model(user).Update("Password", "1234")
	//db.First(&readProduct, "code = ?", "D42") // find product with code D42
	fmt.Printf("%#v", db.First(user, 1))
}
