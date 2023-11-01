package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Email string `gorm:"unique"`
	Age   int
}

func (*User) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "t_old_user"
}
