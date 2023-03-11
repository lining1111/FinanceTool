package HK

/**
现金流量表，数据类计算
*/

const APICS = "http://webapi.cninfo.com.cn/api/hk/p_hk4019"

//CashFlowSheet 现金流量表 "http://webapi.cninfo.com.cn/api/hk/p_hk4019"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			rdate	报告期	string	报告期 可为空，为空取所有报告期
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type CashFlowSheet struct {
	SECCODE     string `excel:"证券代码:"` //证券代码
	SECNAME     string `excel:"证券简称:"` //证券简称
	DECLAREDATE string `excel:"公布日期"`  //公布日期
	STARTDATE   string `excel:"起始日期"`  //起始日期
	ENDDATE     string `excel:"截止日期"`  //截止日期

	F001v string  `excel:"报表类别"`                     //报表类别
	F002V string  `excel:"币种编码"`                     //币种编码
	F003V string  `excel:"币种"`                       //币种
	F004N float64 `excel:"汇率(报表货币兑港元)"`              //汇率(报表货币兑港元)
	F005N float64 `excel:"经营业务收到的现金/经营活动之现金流量(经调整)"` //经营业务收到的现金/经营活动之现金流量(经调整)
	F006N float64 `excel:"已收利息"`                     //已收利息
	F007N float64 `excel:"已付利息"`                     //已付利息
	F008N float64 `excel:"已收股息"`                     //已收股息
	F009N float64 `excel:"已派股息"`                     //已派股息
	F010N float64 `excel:"其他(投资回报)"`                 //其他(投资回报)
	F011N float64 `excel:"投资回报及融资费用之现金流量净额"`         //投资回报及融资费用之现金流量净额
	F012N float64 `excel:"退回/(已缴)税项"`                //退回/(已缴)税项
	F013N float64 `excel:"增添固定资产"`                   //增添固定资产
	F014N float64 `excel:"投资增加"`                     //投资增加
	F015N float64 `excel:"出售固定资产"`                   //出售固定资产
	F016N float64 `excel:"投资减少"`                     //投资减少
	F017N float64 `excel:"与关连人士之现金流量净额(投资活动)"`       //与关连人士之现金流量净额(投资活动)
	F018N float64 `excel:"其他(投资活动)"`                 //其他(投资活动)
	F019N float64 `excel:"投资活动之现金流量(经调整)"`           //投资活动之现金流量(经调整)
	F020N float64 `excel:"融资活动前之现金流量净额"`             //融资活动前之现金流量净额
	F021N float64 `excel:"新增贷款"`                     //新增贷款
	F022N float64 `excel:"偿还贷款"`                     //偿还贷款
	F023N float64 `excel:"债务融资收到的现金"`                //债务融资收到的现金
	F024N float64 `excel:"股权融资收到的现金"`                //股权融资收到的现金
	F025N float64 `excel:"与关连人士之现金流量净额(融资活动)"`       //与关连人士之现金流量净额(融资活动)
	F026N float64 `excel:"偿还债务支付的现金"`                //偿还债务支付的现金
	F027N float64 `excel:"其他(融资活动)"`                 //其他(融资活动)
	F028N float64 `excel:"融资活动之现金流量(经调整)"`           //融资活动之现金流量(经调整)
	F029N float64 `excel:"现金及现金等价物之增加/(减少)"`         //现金及现金等价物之增加/(减少)
	F030N float64 `excel:"年初之现金及现金等价物"`              //年初之现金及现金等价物
	F031N float64 `excel:"外币汇率变动之影响/其他"`             //外币汇率变动之影响/其他
	F032N float64 `excel:"年终之现金及现金等价物"`              //年终之现金及现金等价物
	F033V string  `excel:"核数师意见:"`                   //核数师意见
	F034N float64 `excel:"经营活动产生的现金流量净额(报告)"`        //经营活动产生的现金流量净额(报告)
	F035N float64 `excel:"投资活动产生的现金流量净额(报告)"`        //投资活动产生的现金流量净额(报告)
	F036N float64 `excel:"融资活动产生的现金流量净额(报告)"`        //融资活动产生的现金流量净额(报告)
	F037V string  `excel:"报表类型编码"`                   //报表类型编码 其中1 代表半年报，2 代表年报  	MEMO string  `excel:"备注:""` //备注

	USD float64 `excel:"报告货币兑美元"`  //报告货币兑美元 报表截止日期当天汇率
	CNY float64 `excel:"报告货币兑人民币"` //报告货币兑人民币 报表截止日期当天汇率
}
