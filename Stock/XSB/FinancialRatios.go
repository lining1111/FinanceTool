package XSB

const APIFR = "http://webapi.cninfo.com.cn/api/neeq/p_neeq6017"

//FinancialRatios 财务指标 "http://webapi.cninfo.com.cn/api/neeq/p_neeq6017"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询日期	string	否
//			edate	结束查询日期	string	否
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type FinancialRatios struct {
	ORGNAME     string `excel:"机构名称"` //机构名称	varchar
	SECCODE     string `excel:"证券代码"` //证券代码	varchar
	SECNAME     string `excel:"证券简称"` //证券简称	varchar
	DECLAREDATE string `excel:"公告日期"` //公告日期	Date
	STARTDATE   string `excel:"开始日期"` //开始日期	Date
	ENDDATE     string `excel:"截止日期"` //截止日期	Date

	F001V string  `excel:"数据来源编码"`                //数据来源编码	varchar		033003定期报告033004业绩快报033006转让说明书033013其他
	F002V string  `excel:"数据来源"`                  //数据来源	varchar		详见数据来源编码备注
	F003N float64 `excel:"资产总计"`                  //资产总计	decimal		单位：元
	F004N float64 `excel:"归属于母公司所有者权益合计"`         //归属于母公司所有者权益合计	decimal		单位：元
	F005N float64 `excel:"归属于母公司所有者每股净资产"`        //归属于母公司所有者每股净资产	decimal		单位：元
	F006N float64 `excel:"归属于母公司所有者净利润"`          //归属于母公司所有者净利润	decimal		单位：元
	F007N float64 `excel:"扣除非经常性损益归属于母公司所有者的净利润"` //扣除非经常性损益归属于母公司所有者的净利润	decimal		单位：元
	F008N float64 `excel:"净资产收益率-加权"`             //净资产收益率-加权	decimal
	F009N float64 `excel:"净资产收益率-加权(扣除非经常性损益"`    //净资产收益率-加权(扣除非经常性损益	decimal
	F010N float64 `excel:"基本每股收益"`                //基本每股收益	decimal		单位：元
	F011N float64 `excel:"稀释每股收益"`                //稀释每股收益	decimal		单位：元
	F012N float64 `excel:"扣除非经常性损益基本每股收益"`        //扣除非经常性损益基本每股收益	decimal		单位：元
	F013N float64 `excel:"经营活动产生的现金流量净额"`         //经营活动产生的现金流量净额	decimal		单位：元
	F014N float64 `excel:"每股经营活动产生的现金流量净额"`       //每股经营活动产生的现金流量净额	decimal		单位：元
	F015N float64 `excel:"净资产收益率-摊薄"`             //净资产收益率-摊薄	decimal
	F016N float64 `excel:"净资产收益率_摊薄(扣除非经常性损益)"`   //净资产收益率_摊薄(扣除非经常性损益)	decimal
	MEMO  string  `excel:"备注"`                    //备注	varchar
}
