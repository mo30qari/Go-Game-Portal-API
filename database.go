package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
)

var dbAddress = "root:@tcp(127.0.0.1:3306)/game_portal?charset=utf8mb4&parseTime=True&loc=Local"

func openDb() *gorm.DB{

	db, err := gorm.Open("mysql", dbAddress)

	if err != nil {
		fmt.Println("Can't Connect to the Database!")
		panic(err.Error())
	}

	db.AutoMigrate(&User{}, &Game{}, &Genre{})

	return db

}

func closeDb(db *gorm.DB){

	db.Close()

}
