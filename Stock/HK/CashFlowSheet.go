package HK

/**
现金流量表，数据类计算
*/

const apiCS = "http://webapi.cninfo.com.cn/api/hk/p_hk4019"

//CashFlowSheet 现金流量表 "http://webapi.cninfo.com.cn/api/hk/p_hk4019"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			rdate	报告期	string	报告期 可为空，为空取所有报告期
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type CashFlowSheet struct {
	SECCODE     string //证券代码	varchar
	SECNAME     string //证券简称	varchar
	DECLAREDATE string //公布日期	date
	STARTDATE   string //起始日期	date
	ENDDATE     string //截止日期	date

	F001v string  //报表类别	varchar
	F002V string  //币种编码	varchar
	F003V string  //币种	varchar
	F004N float64 //汇率(报表货币兑港元)	numeric
	F005N float64 //经营业务收到的现金/经营活动之现金流量(经调整)	numeric
	F006N float64 //已收利息	numeric
	F007N float64 //已付利息	numeric
	F008N float64 //已收股息	numeric
	F009N float64 //已派股息	numeric
	F010N float64 //其他(投资回报)	numeric
	F011N float64 //投资回报及融资费用之现金流量净额	numeric
	F012N float64 //退回/(已缴)税项	numeric
	F013N float64 //增添固定资产	numeric
	F014N float64 //投资增加	numeric
	F015N float64 //出售固定资产	numeric
	F016N float64 //投资减少	numeric
	F017N float64 //与关连人士之现金流量净额(投资活动)	numeric
	F018N float64 //其他(投资活动)	numeric
	F019N float64 //投资活动之现金流量(经调整)	numeric
	F020N float64 //融资活动前之现金流量净额	numeric
	F021N float64 //新增贷款	numeric
	F022N float64 //偿还贷款	numeric
	F023N float64 //债务融资收到的现金	numeric
	F024N float64 //股权融资收到的现金	numeric
	F025N float64 //与关连人士之现金流量净额(融资活动)	numeric
	F026N float64 //偿还债务支付的现金	numeric
	F027N float64 //其他(融资活动)	numeric
	F028N float64 //融资活动之现金流量(经调整)	numeric
	F029N float64 //现金及现金等价物之增加/(减少)	numeric
	F030N float64 //年初之现金及现金等价物	numeric
	F031N float64 //外币汇率变动之影响/其他	numeric
	F032N float64 //年终之现金及现金等价物	numeric
	F033V string  //核数师意见	varchar
	F034N float64 //经营活动产生的现金流量净额(报告)	numeric
	F035N float64 //投资活动产生的现金流量净额(报告)	numeric
	F036N float64 //融资活动产生的现金流量净额(报告)	numeric
	F037V string  //报表类型编码	char		其中1 代表半年报，2 代表年报

	MEMO string  //备注	varchar
	USD  float64 //报告货币兑美元	double		报表截止日期当天汇率
	CNY  float64 //报告货币兑人民币	double		报表截止日期当天汇率
}
