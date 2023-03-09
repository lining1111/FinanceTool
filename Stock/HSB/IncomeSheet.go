package HSB

import (
	"FinanceTool/COM/cninfo"
	"fmt"
)

const APIIS = "http://webapi.cninfo.com.cn/api/stock/p_stock2301"

//IncomeSheet 利润表 "http://webapi.cninfo.com.cn/api/stock/p_stock2301"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	查询开始报告期	string	否
//			edate	查询结束报告期	string	否
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type IncomeSheet struct {
	SECCODE     string //证券代码	varchar
	SECNAME     string //证券简称	varchar
	ORGNAME     string //机构名称	varchar
	DECLAREDATE string //公告日期	date
	STARTDATE   string //开始日期	date
	ENDDATE     string //截止日期	date

	F001D string  //报告年度	date
	F002V string  //合并类型编码	varchar
	F003V string  //合并类型	varchar
	F004V string  //报表来源编码	varchar
	F005V string  //报表来源	varchar
	F035N float64 //一、营业总收入	decimal		单位：元
	F006N float64 //其中：营业收入	decimal		单位：元
	F033N float64 //利息收入	decimal		单位：元
	F034N float64 //已赚保费	decimal		单位：元
	F042N float64 //手续费及佣金收入	decimal		单位：元
	F036N float64 //二、营业总成本	decimal		单位：元
	F007N float64 //其中：营业成本	decimal		单位：元
	F043N float64 //利息支出	decimal		单位：元
	F044N float64 //手续费及佣金支出	decimal		单位：元
	F045N float64 //退保金	decimal		单位：元
	F046N float64 //赔付支出净额	decimal		单位：元
	F047N float64 //提取保险合同准备金净额	decimal		单位：元
	F048N float64 //保单红利支出	decimal		单位：元
	F049N float64 //分保费用	decimal		单位：元
	F008N float64 //营业税金及附加	decimal		单位：元
	F009N float64 //销售费用	decimal		单位：元
	F010N float64 //管理费用	decimal		单位：元
	F011N float64 //勘探费用	decimal		单位：元
	F012N float64 //财务费用	decimal		单位：元
	F056N float64 //研发费用	DECIMAL(18,2)
	F013N float64 //资产减值损失	decimal		单位：元
	F014N float64 //加：公允价值变动净收益	decimal		单位：元
	F015N float64 //投资收益	decimal		单位：元
	F016N float64 //其中：对联营企业和合营企业的投资收益	decimal		单位：元
	F037N float64 //汇兑收益	decimal		单位：元
	F051N float64 //其它收入	DECIMAL(18,2)
	F057N float64 //信用减值损失	DECIMAL(18,2)
	F058N float64 //净敞口套期收益	DECIMAL(18,2)
	F059N float64 //资产处置收益	DECIMAL(18,2)
	F017N float64 //影响营业利润的其他科目	decimal		单位：元
	F018N float64 //三、营业利润	decimal		单位：元
	F019N float64 //加：补贴收入	decimal		单位：元
	F020N float64 //营业外收入	decimal		单位：元
	F050N float64 //其中：非流动资产处置利得	decimal		单位：元
	F021N float64 //减：营业外支出	decimal		单位：元
	F022N float64 //其中：非流动资产处置损失	decimal		单位：元
	F023N float64 //加：影响利润总额的其他科目	decimal		单位：元
	F024N float64 //四、利润总额	decimal		单位：元
	F025N float64 //减：所得税	decimal		单位：元
	F026N float64 //加：影响净利润的其他科目	decimal		单位：元
	F027N float64 //五、净利润	decimal		单位：元
	F060N float64 //持续经营净利润	DECIMAL(18,2)
	F061N float64 //终止经营净利润	DECIMAL(18,2)
	F028N float64 //归属于母公司所有者的净利润	decimal(18,2)		单位：元
	F029N float64 //少数股东损益	decimal(18,2)		单位：元
	F031N float64 //（一）基本每股收益	decimal(18,2)
	F032N float64 //（二）稀释每股收益	decimal(18,2)
	F038N float64 //七、其他综合收益	decimal(18,2)		单位：元
	F039N float64 //八、综合收益总额	decimal(18,2)		单位：元
	F040N float64 //其中：归属于母公司	decimal(18,2)		单位：元
	F041N float64 //其中：归属于少数股东	decimal(18,2)		单位：元
	MEMO  string  //备注	VARCHAR
	F062N float64 //其中：利息费用	decimal(18,2)		2019-03-15新增
	F063N float64 //其中：利息收入	decimal(18,2)		2019-03-15新增
	F064N float64 //信用减值损失（2019格式）	decimal(18,2)
	F065N float64 //资产减值损失（2019格式）	decimal(18,2)
	F066N float64 //其中：归属母公司所有者的其他综合收益的税后净额	decimal(18,2)
	F067N float64 //其中：归属于少数股东的其他综合收益的税后净额	decimal(18,2)
}

func ISGetFromCNINFByScode_test() {
	is := make([]IncomeSheet, 1, 20000)
	url := APIIS
	params := map[string]string{
		"scode": "000001",
		"sdate": cninfo.Getrdate("2021", cninfo.Q1),
		"edate": cninfo.Getrdate("2021", cninfo.Q1),
		"type":  "071001",
	}

	err := cninfo.GetInfoByScodeDate(url, params, &is, cap(is))
	if err != nil {
		fmt.Println(err)
	}
}
