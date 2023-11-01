package model

type NewCityModel struct {
	CityCode  string         `json:"citycode"`
	AdCode    string         `json:"adcode" `
	Name      string         `json:"name" `
	Center    string         `json:"center" `
	Level     string         `json:"level" `
	Districts []NewCityModel `json:"districts"`
}
