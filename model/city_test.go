package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"moduledemo/common"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestImportJsonToDb(t *testing.T) {
	common.InitMysql()
	// 打开json文件
	jsonFile, err := os.Open("D:\\data01\\city.json")
	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}
	// 要记得关闭
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//fmt.Println(string(byteValue))
	var cityModelList []CityModel
	jsonParseErr := json.Unmarshal(byteValue, &cityModelList)
	if jsonParseErr != nil {
		fmt.Println(jsonParseErr)
	}

	fmt.Println("cityModelList===")
	fmt.Println(cityModelList)
	for i := 0; i < len(cityModelList); i++ {
		province := cityModelList[i]
		id, _ := strconv.Atoi(province.Code)
		pro := City{Id: id, Code: province.Code, Name: province.Name, FullName: province.FullName, initDate: time.Now()}
		common.DB.Model(&City{}).Create(&pro)
		//fmt.Println("pro id=", pro.Id)
		//fmt.Println("i =", i, "province =", province)
		for j := 0; j < len(province.Children); j++ {
			city := province.Children[j]
			//fmt.Println("j =", j, "city =", city)
			// 如果是直辖市 则不用插入数据
			if city.Code != province.Code {
				id, _ := strconv.Atoi(city.Code)
				parentId, _ := strconv.Atoi(province.Code)
				ct := City{Id: id, Code: city.Code, Name: city.Name, FullName: city.FullName, ParentId: parentId, initDate: time.Now()}
				common.DB.Model(&City{}).Create(&ct)
			}

			for k := 0; k < len(city.Children); k++ {
				area := city.Children[k]
				//fmt.Println("k =", k, "area =", area)
				id, _ := strconv.Atoi(area.Code)
				parentId, _ := strconv.Atoi(city.Code)
				ar := City{Id: id, Code: area.Code, Name: area.Name, FullName: area.FullName, ParentId: parentId, initDate: time.Now()}
				common.DB.Model(&City{}).Create(&ar)
			}
		}
	}

}
