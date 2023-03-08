package AuditOpinion

import (
	"FinanceTool/COM/cninfo"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/goinggo/mapstructure"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

const urlApi = "http://webapi.cninfo.com.cn/api/stock/p_stock2239"

//AuditOpinion 审计意见 "http://webapi.cninfo.com.cn/api/stock/p_stock2239"
type AuditOpinion struct {
	SECNAME     string //证券简称
	SECCODE     string //证券代码
	DECLAREDATE string //公告日期

	F001D string //报告年度
	F002C string //是否经审计
	F003V string //境内会计师事务所ID
	F004V string //境内会计师事务所名称
	F005V string //境内会计师事务所签名注册会计师
	F006V string //境内会计师事务所审计意见类型编码
	F007V string //境内会计师事务所审计意见类型
	F008V string //非标准审计意见的事项内容

	F009V string //境外会计师事务所ID
	F010V string //境外会计师事务所名称
	F011V string //境外会计师事务所签名注册会计师
	F012V string //境外审计报告会计师事务所审计意见类型编码
	F013V string //境外审计报告会计师事务所审计意见类型
	F014V string //境外审计意见非标准的事项说明

	F015V string //董事会对会计师事务所出具的有保留意见书、解释说明
	F016V string //监事会对董事会就上述事项的说明所表示的意见
}

//GetFromCNINFByScode 从巨潮资讯平台以股票代码和日期为输入获取
func (ao *AuditOpinion) GetFromCNINFByScode(scode string, year string, q string) (bool, error) {
	url := urlApi
	params := map[string]string{
		"scode": scode,
		"sdate": cninfo.Getrdate(year, q),
		"edate": cninfo.Getrdate(year, q),
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
				if result["resultcode"].(float64) != http.StatusOK {
					glog.Error("result:%s", string(body))
					return false, errors.New("http req err:" + string(body))
				}

				//打印下数组数量
				fmt.Println("回复结果数为 ", result["total"])
				if result["total"].(float64) == 0 {
					return false, errors.New("http result total 0")
				}
				//反序列化
				err3 := mapstructure.Decode(result["records"].([]interface{})[0], ao)
				if err3 != nil {
					return false, err2
				}
			}
		}
	}
	return true, nil
}

func GetFromCNINFByScode_test() {
	ao := new(AuditOpinion)
	_, err := ao.GetFromCNINFByScode("000001", "2021", cninfo.Q1)
	if err != nil {
		fmt.Println(err)
	}
}
