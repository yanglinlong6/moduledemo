package model

type OldProvinceModel struct {
	Id   int    `json:"id"`
	Name string `json:"name" `
}

type OldCityModel struct {
	Province string `json:"province"`
	Name     string `json:"name"`
	Id       int    `json:"id"`
}

type OldAreaModel struct {
	City string `json:"city"`
	Name string `json:"name"`
	Id   string `json:"id"`
}
