package model

import "time"

type DAlarmSecondData struct {
	ID           int       `json:"id" gorm:"column:id"`                       // id
	Sn           string    `json:"sn" gorm:"column:sn"`                       // 设备号
	Type         string    `json:"type" gorm:"column:type"`                   // 二押区域规则类型(二押地停留)
	Name         string    `json:"name" gorm:"column:name"`                   // 二押区域规则名称(二押地停留)
	AlarmTime    time.Time `json:"alarm_time" gorm:"column:alarm_time"`       // 二押区域报警时间
	AlarmAddress string    `json:"alarm_address" gorm:"column:alarm_address"` // 二押区域报警位置
	LeaveTime    time.Time `json:"leave_time" gorm:"column:leave_time"`       // 离开二押区域时间(当前还未动身的数据不记录到数据库)
	StayTime     int       `json:"stay_time" gorm:"column:stay_time"`         // 二押区域停留时长
	Lng          string    `json:"lng" gorm:"column:lng"`                     // 二押区域停留经度
	Lat          string    `json:"lat" gorm:"column:lat"`                     // 二押区域停留纬度
	EventTime    time.Time `json:"event_time" gorm:"column:event_time"`       // 事件时间，写入事件
	DelFlag      int8      `json:"del_flag" gorm:"column:del_flag"`           // 删除标识 1 是最新插入 2 历史数据
	UniqueFlag   string    `json:"unique_flag" gorm:"column:unique_flag"`     // 唯一标识 设备号加上报警时间
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`     // 创建时间
	UpdateTime   time.Time `json:"update_time" gorm:"column:update_time"`     // 更新时间
}

func (m *DAlarmSecondData) TableName() string {
	return "d_alarm_second_data"
}
