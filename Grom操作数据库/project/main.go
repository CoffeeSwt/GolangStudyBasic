package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint      `gorm:"comment:用户ID"`
	UserName string    `gorm:"comment:用户名"`
	Gender   string    `gorm:"comment:性别"`
	Birthday time.Time `gorm:"comment:出生日期"`
}

func read(db *gorm.DB) (User, error) {
	var user User
	db = db.Model(&User{})
	err := db.Select("name,gender").Find(&user).Error
	return user, err
}

func main() {
	fmt.Println("Hello Golang")
	dataSource := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dataSource), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err := read(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
}
