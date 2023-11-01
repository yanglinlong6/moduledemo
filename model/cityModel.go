package model

type CityModel struct {
	Id       int        `json:"id"`
	Code     string     `json:"code" `
	Name     string     `json:"name" `
	FullName string     `json:"fullname" `
	Children []Children `json:"children"`
}

type Children struct {
	Code     string     `gorm:"column:code" json:"code"`
	Name     string     `gorm:"column:name" json:"name"`
	FullName string     `gorm:"column:full_name" json:"fullname"`
	ParentId int        `gorm:"column:parent_id"`
	Children []Children `json:"children"  sql:"-"`
}
