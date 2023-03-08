package HK

const apiFRBank = "http://webapi.cninfo.com.cn/api/hk/p_hk4022"

//FinancialRatiosBank 重要指标表（银行） "http://webapi.cninfo.com.cn/api/hk/p_hk4022"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type FinancialRatiosBank struct {
	SECCODE     string //证券代码	varchar
	SECNAME     string //证券简称	varchar
	DECLAREDATE string //公布日期	date
	ENDDATE     string //截止日期	date

	F001V string  //报表类别	varchar
	F002V string  //币种编码	varchar
	F003V string  //币种	varchar
	F004N float64 //汇率(报表货币兑港元)	numeric
	F005N float64 //每股基本盈利（港元）	numeric
	F006N float64 //每股摊薄盈利（港元）	numeric
	F008N float64 //客户贷款及垫款	numeric
	F007N float64 //银行同业贷款及垫款	numeric
	F009N float64 //定期存放银行同业及其他财务机构	numeric
	F010N float64 //总资产	numeric
	F011N float64 //同业及财务机构存款	numeric
	F012N float64 //客户存款	numeric
	F013N float64 //总负债	numeric
	F014N float64 //少数股东权益	numeric
	F015N float64 //股东权益/(资产净值)	numeric
	F016N float64 //股本	numeric
	F017N float64 //总储备	numeric
	F018N float64 //净利息收益	numeric
	F019N float64 //费用收益净额	numeric
	F020N float64 //交易收益净额	numeric
	F021N float64 //其他营业收益	numeric
	F022N float64 //营业收益总额	numeric
	F023N float64 //贷款减值及其他信贷风险准备	numeric
	F024N float64 //营业利润/(亏损)	numeric
	F025N float64 //非盈利/特殊项目	numeric
	F026N float64 //应占联营公司溢利	numeric
	F027N float64 //除税前经营溢利/(亏损)	numeric
	F028N float64 //税项	numeric
	F029N float64 //除税后经营溢利/(亏损)	numeric
	F030N float64 //少数股东损益	numeric
	F031N float64 //经调整股东应占溢利	numeric
	F032N float64 //相对上期增减	numeric
	F033N float64 //股息总额	numeric
	F034N float64 //扣除非经常性项目后溢利	numeric
	F035N float64 //经营业务收到的现金/经营活动之现金流量（经调整）	numeric
	F036N float64 //投资活动之现金流量（经调整）	numeric
	F037N float64 //融资活动之现金流量（经调整）	numeric
	F038N float64 //现金流量净额	numeric
	F039N float64 //年初之现金及现金等价物	numeric
	F040N float64 //外币汇率变动之影响/其他	numeric
	F041N float64 //年终之现金及现金等价物	numeric
	F042N float64 //资本充足比率	numeric
	F043N float64 //流动资金比率	numeric
	F044N float64 //成本/收入	numeric
	F045N float64 //流动资金/存款	numeric
	F046N float64 //贷款/存款	numeric
	F047N float64 //贷款/股东权益	numeric
	F048N float64 //贷款/总资产	numeric
	F049N float64 //存款/股东权益	numeric
	F050N float64 //存款/总资产	numeric
	F051N float64 //贷款回报率	numeric
	F052N float64 //存款回报率	numeric
	F053N float64 //资产回报率	numeric
	F054N float64 //权益报酬率	numeric
	F055N float64 //股东权益合计	numeric
	F056N float64 //未分配利润	numeric
	F057V string  //核数师意见	varchar
	F058N float64 //平均股东权益回报率	numeric
	F059N float64 //平均资产回报率	numeric
	F060N float64 //经调整每股现金流量	numeric
	F061N float64 //期末股价/每股资产净值	numeric
	F062V string  //报表类型编码	char		其中1代表半年报，2代表年报

	MEMO string  //备注	varchar
	USD  float64 //报告货币兑美元	double		报表截止日期当天汇率
	CNY  float64 //报告货币兑人民币	double		报表截止日期当天汇率
}
