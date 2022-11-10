package response

import "moduledemo/model"

type LinLong struct {
	Name     string              `json:"name"`
	Yang     string              `json:"yang"`
	List     []*model.TNesDevice `json:"list"`
	YangList []*model.TNesDevice `json:"yangList"`
}
