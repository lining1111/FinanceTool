package Fund

const apiFR = "http://webapi.cninfo.com.cn/api/fund/p_fund2610"

//FinancialRatios 基金主要财务指标2009版 "http://webapi.cninfo.com.cn/api/fund/p_fund2610"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			rdate	报告期	string	报告期 可为空，为空取所有报告期
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type FinancialRatios struct {
	SECCODE     string //基金代码	varchar
	SECNAME     string //基金简称	varchar
	DECLAREDATE string //公告日期	date
	STARTDATE   string //开始日期	date
	ENDDATE     string //截止日期	date

	F001N float64 //本期已实现收益	decimal		单位：元
	F002N float64 //本期利润	decimal		单位：元
	F004N float64 //加权平均份额本期利润	decimal		单位：元
	F005N float64 //加权平均净值利润率	decimal		单位：%
	F006N float64 //期末可供分配利润	decimal		单位：元
	F007N float64 //期末可供分配份额利润	decimal		单位：元
	F008N float64 //期未基金资产净值	decimal		单位：元
	F009N float64 //期末基金份额净值	decimal		单位：元
	F011N float64 //本期净值增长率	decimal		单位：%
	F012N float64 //累计净值增长率	decimal		单位：%
}
