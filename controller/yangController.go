package controller

import (
	"github.com/gin-gonic/gin"
	"moduledemo/model/request"
	"moduledemo/service"
	"moduledemo/tools"
	"net/http"
)

type YangController struct{}

var (
	Yang = &YangController{}
)

// List 记录列表
func (m *YangController) Hello(c *gin.Context) {
	req := new(request.Yang)
	Run(c, req, func() (interface{}, interface{}) {
		return service.YangService(c, req)
	})

}

func (m *YangController) Save(c *gin.Context) {
	req := new(request.Yang)
	Run(c, req, func() (interface{}, interface{}) {
		return service.Save(c, req)
	})

}

func Demo(c *gin.Context) {
	c.JSON(http.StatusOK, tools.H{"code": 200, "msg": "ok", "data": "pong"})
}

func Run(c *gin.Context, req interface{}, fn func() (interface{}, interface{})) {
	var err error
	// bind struct
	err = c.ShouldBind(req)
	if err != nil {
		tools.Err(c, tools.NewValidatorError(err), nil)
		return
	}

	data, err1 := fn()
	if err1 != nil {
		tools.Err(c, tools.ReloadErr(err1), data)
		return
	}
	tools.Success(c, data)
}
