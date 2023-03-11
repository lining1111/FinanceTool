package Fund

const apiFR = "http://webapi.cninfo.com.cn/api/fund/p_fund2610"

//FinancialRatios 基金主要财务指标2009版 "http://webapi.cninfo.com.cn/api/fund/p_fund2610"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			rdate	报告期	string	报告期 可为空，为空取所有报告期
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type FinancialRatios struct {
	SECCODE     string `excel:"基金代码"` //基金代码
	SECNAME     string `excel:"基金简称"` //基金简称
	DECLAREDATE string `excel:"公告日期"` //公告日期
	STARTDATE   string `excel:"开始日期"` //开始日期
	ENDDATE     string `excel:"截止日期"` //截止日期

	F001N float64 `excel:"本期已实现收益"`    //本期已实现收益 单位：元
	F002N float64 `excel:"本期利润"`       //本期利润 单位：元
	F004N float64 `excel:"加权平均份额本期利润"` //加权平均份额本期利润 单位：元
	F005N float64 `excel:"加权平均净值利润率"`  //加权平均净值利润率 单位：%
	F006N float64 `excel:"期末可供分配利润"`   //期末可供分配利润 单位：元
	F007N float64 `excel:"期末可供分配份额利润"` //期末可供分配份额利润 单位：元
	F008N float64 `excel:"期未基金资产净值"`   //期未基金资产净值 单位：元
	F009N float64 `excel:"期末基金份额净值"`   //期末基金份额净值 单位：元
	F011N float64 `excel:"本期净值增长率"`    //本期净值增长率 单位：%
	F012N float64 `excel:"累计净值增长率"`    //累计净值增长率 单位：%
}
