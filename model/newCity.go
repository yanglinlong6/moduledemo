package model

import "time"

type NewCity struct {
	Id       int       `gorm:"primary_key column:id"`
	CityCode string    `gorm:"column:city_code"`
	AdCode   string    `gorm:"column:area_code"`
	Name     string    `gorm:"column:area_name"`
	Center   string    `gorm:"column:center"`
	Level    string    `gorm:"column:level"`
	ParentId int       `gorm:"column:parent_id"`
	initDate time.Time `gorm:"column:init_date"`
}

func (m *NewCity) TableName() string {
	return "d_city"
}
