package BalanceSheet

import (
	"FinanceTool/FinancialStatement/cninfo"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/goinggo/mapstructure"
	"github.com/golang/glog"
	"io/ioutil"
)

/**
资产负债表格式化处理，包括但不限于 xls html 数据库
这里数据库选项选择 sqlite3
*/

//GetFromCNINFByScode 从巨潮资讯平台以股票代码和日期为输入获取
func (bs *BalanceSheet) GetFromCNINFByScode(scode string, year string, q string) (bool, error) {
	url := "http://webapi.cninfo.com.cn/api/stock/p_stock2300"
	params := map[string]string{
		"scode": scode,
		"rdate": cninfo.Getrdate(year, q),
		"type":  "071001",
	}
	resp, err := cninfo.Post(url, nil, params, nil)
	defer resp.Body.Close()
	if err != nil {
		return false, err
	} else {
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			return false, err1
		} else {
			//打印结果
			//fmt.Println(string(body))
			var result map[string]interface{}
			err2 := json.Unmarshal(body, &result)
			if err2 != nil {
				return false, err2
			} else {
				if result["resultmsg"].(string) != "success" {
					glog.Error("result:%s", string(body))
					return false, errors.New("http req err")
				}

				//打印下数组数量
				fmt.Println("回复结果数为 ", result["total"])
				//fmt.Println(result["records"].([]interface{})[0])
				//反序列化
				err3 := mapstructure.Decode(result["records"].([]interface{})[0], bs)
				if err3 != nil {
					return false, err2
				}
			}
		}
	}
	return true, nil
}

func GetFromCNINFByScode_test() {
	bs := new(BalanceSheet)
	_, err := bs.GetFromCNINFByScode("000001", "2021", cninfo.Q1)
	if err != nil {
		fmt.Println(err)
	}
}
