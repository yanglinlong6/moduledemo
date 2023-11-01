package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"moduledemo/common"
	"os"
	"testing"
)

func OldTestImportJsonToDb(t *testing.T) {
	common.InitMysql()
	// 打开json文件
	jsonFile, err := os.Open("D:\\data01\\老E盾province.json")
	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}
	// 要记得关闭
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//fmt.Println(string(byteValue))
	var oldProvinceModelList []OldProvinceModel
	jsonParseErr := json.Unmarshal(byteValue, &oldProvinceModelList)
	if jsonParseErr != nil {
		fmt.Println(jsonParseErr)
	}

	fmt.Println("oldProvinceModelList===")
	fmt.Println(oldProvinceModelList)
	for i := 0; i < len(oldProvinceModelList); i++ {
		province := oldProvinceModelList[i]
		fmt.Println("province===", province)
		//id, _ := strconv.Atoi(province.Code)
		//pro := City{Id: id, Code: province.Code, Name: province.Name, FullName: province.FullName, initDate: time.Now()}
		//common.DB.Model(&City{}).Create(&pro)

	}
}
