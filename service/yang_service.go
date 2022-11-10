package service

import (
	"github.com/gin-gonic/gin"
	"moduledemo/common"
	"moduledemo/model"
	"moduledemo/model/request"
	"moduledemo/model/response"
)

func YangService(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, _ := req.(*request.Yang)
	_ = c
	var list []*model.TNesDevice
	var yangList []*model.TNesDevice
	db := common.DB.Model(&model.TNesDevice{}).Order("sn desc")
	db.Find(&list)
	common.DB.Exec("select * from t_nes_device").Where("sn = ?", "1500101471").Offset(0).Limit(100).Find(&yangList)
	return response.LinLong{Name: r.Name, Yang: r.Yang, List: list, YangList: yangList}, nil
	//return response.LinLong{Name: r.Name, Yang: r.Yang, YangList: yangList}, nil
}
