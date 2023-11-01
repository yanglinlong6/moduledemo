package model

import (
	"time"
)

type City struct {
	Id       int       `gorm:"primary_key column:id"`
	Code     string    `json:"code" gorm:"column:area_code"`
	Name     string    `json:"name" gorm:"column:area_name"`
	FullName string    `json:"fullname" gorm:"column:full_name"`
	ParentId int       `gorm:"column:parent_id"`
	initDate time.Time `gorm:"column:init_date"`
}

func (m *City) TableName() string {
	return "t_city"
}
