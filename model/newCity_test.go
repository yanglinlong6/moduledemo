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

func TestImportJsonToNewDb(t *testing.T) {
	common.InitMysql()
	// 打开json文件
	jsonFile, err := os.Open("D:\\data01\\newcity.json")
	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}
	// 要记得关闭
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//fmt.Println(string(byteValue))
	var newCityModelList []NewCityModel
	jsonParseErr := json.Unmarshal(byteValue, &newCityModelList)
	if jsonParseErr != nil {
		fmt.Println(jsonParseErr)
	}

	fmt.Println("newCityModelList===")
	fmt.Println(newCityModelList)
	for i := 0; i < len(newCityModelList); i++ {
		province := newCityModelList[i]
		id, _ := strconv.Atoi(province.AdCode)
		pro := NewCity{Id: id, CityCode: province.CityCode, AdCode: province.AdCode, Name: province.Name, Center: province.Center, Level: province.Level, initDate: time.Now()}
		common.DB.Model(&NewCity{}).Create(&pro)
		//fmt.Println("pro id=", pro.Id)
		//fmt.Println("i =", i, "province =", province)
		for j := 0; j < len(province.Districts); j++ {
			city := province.Districts[j]
			//fmt.Println("j =", j, "city =", city)

			id, _ := strconv.Atoi(city.AdCode)
			parentId, _ := strconv.Atoi(province.AdCode)
			ct := NewCity{Id: id, CityCode: city.CityCode, AdCode: city.AdCode, Name: city.Name, Center: city.Center, Level: province.Level, ParentId: parentId, initDate: time.Now()}
			common.DB.Model(&NewCity{}).Create(&ct)

			for k := 0; k < len(city.Districts); k++ {
				area := city.Districts[k]
				//fmt.Println("k =", k, "area =", area)
				id, _ := strconv.Atoi(area.AdCode)
				parentId, _ := strconv.Atoi(city.AdCode)
				ar := NewCity{Id: id, CityCode: area.CityCode, AdCode: area.AdCode, Name: area.Name, Center: area.Center, Level: province.Level, ParentId: parentId, initDate: time.Now()}
				common.DB.Model(&NewCity{}).Create(&ar)
			}
		}
	}
}
