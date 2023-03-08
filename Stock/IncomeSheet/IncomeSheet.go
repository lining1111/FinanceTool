package IncomeSheet

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

const urlApi = "http://webapi.cninfo.com.cn/api/stock/p_stock2301"

//IncomeSheet 利润表 "http://webapi.cninfo.com.cn/api/stock/p_stock2301"
type IncomeSheet struct {
	SECNAME     string //证券简称
	SECCODE     string //证券代码
	ORGNAME     string //机构名称
	DECLAREDATE string //公告日期
	ENDDATE     string //截止日期
	F001D       string //报告年度
	F002V       string //合并类型编码
	F003V       string //合并类型
	F004V       string //报表来源编码
	F005V       string //报表来源

	F035N float64 //一、营业总收入
	F006N float64 //其中：营业收入
	F033N float64 //利息收入
	F034N float64 //已赚保费
	F042N float64 //手续费及佣金收入
	F036N float64 //二、营业总成本
	F007N float64 //其中：营业成本
	F043N float64 //利息支出
	F044N float64 //手续费及佣金支出
	F045N float64 //退保金
	F046N float64 //赔付支出净额
	F047N float64 //提取保险合同准备金净额
	F048N float64 //保单红利支出
	F049N float64 //分保费用
	F008N float64 //营业税金及附加
	F009N float64 //销售费用
	F010N float64 //管理费用
	F011N float64 //勘探费用
	F012N float64 //财务费用
	F056N float64 //研发费用
	F013N float64 //资产减值损失
	F014N float64 //加：公允价值变动净收益
	F015N float64 //投资收益
	F016N float64 //其中：对联营企业和合营企业的投资收益
	F037N float64 //汇兑收益
	F051N float64 //其它收入
	F057N float64 //信用减值损失
	F058N float64 //净敞口套期收益
	F059N float64 //资产处置收益
	F017N float64 //影响营业利润的其他科目
	F018N float64 //三、营业利润
	F019N float64 //加：补贴收入
	F020N float64 //营业外收入
	F050N float64 //其中：非流动资产处置利得
	F021N float64 //减：营业外支出
	F022N float64 //其中：非流动资产处置损失
	F023N float64 //加：影响利润总额的其他科目
	F024N float64 //四、利润总额
	F025N float64 //减：所得税
	F026N float64 //加：影响净利润的其他科目
	F027N float64 //五、净利润
	F060N float64 //持续经营净利润
	F061N float64 //终止经营净利润
	F028N float64 //归属于母公司所有者的净利润
	F029N float64 //少数股东损益
	F031N float64 //(一）基本每股收益
	F032N float64 //（二）稀释每股收益
	F038N float64 //七、其他综合收益
	F039N float64 //八、综合收益总额
	F040N float64 //其中：归属于母公司
	F041N float64 //其中：归属于少数股东
	MEMO  string  //备注
	F062N float64 //其中：利息费用
	F063N float64 //其中：利息收入
	F064N float64 //信用减值损失（2019格式）
	F065N float64 //资产减值损失（2019格式）
	F066N float64 //其中：归属母公司所有者的其他综合收益的税后净额
	F067N float64 //其中：归属于少数股东的其他综合收益的税后净额
}

//GetFromCNINFByScode 从巨潮资讯平台以股票代码和日期为输入获取
func (is *IncomeSheet) GetFromCNINFByScode(scode string, year string, q string) (bool, error) {
	url := urlApi
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
				err3 := mapstructure.Decode(result["records"].([]interface{})[0], is)
				if err3 != nil {
					return false, err2
				}
			}
		}
	}
	return true, nil
}

func GetFromCNINFByScode_test() {
	is := new(IncomeSheet)
	_, err := is.GetFromCNINFByScode("000001", "2021", cninfo.Q1)
	if err != nil {
		fmt.Println(err)
	}
}
