package model

import "github.com/jinzhu/gorm"

/*
*	nickname允许重复，mobile当做用户的唯一标识
 */
type User struct {
	gorm.Model
	//Id       int64
	Mobile   string `gorm:"varchar(110);notnull;unique"`
	Password string `gorm:"size:255;notnull"`
	//Online   int
	//Token    string
	Nickname string `gorm:"varchar(20);notnull"`
}
