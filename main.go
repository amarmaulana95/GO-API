package main

import (
	"fmt"
	"log"
	"os/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/db_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("koneksi ok")

	// tampilkan user, masukan ke array of
	var users []user.User
	length := len(users)

	fmt.Println(length)

	db.Find(&users)

	length = len(users)
	fmt.Println(length)

}
