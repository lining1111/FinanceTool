package HK

const APIFR = "http://webapi.cninfo.com.cn/api/hk/p_hk4025"

//FinancialRatios 重要指标表（非银行） "http://webapi.cninfo.com.cn/api/hk/p_hk4025"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type FinancialRatios struct {
	SECCODE     string //证券代码	varchar
	SECNAME     string //证券简称	varchar
	DECLAREDATE string //公布日期	date
	STARTDATE   string //起始日期	date
	ENDDATE     string //截止日期	date

	F001V string  //报表类别	varchar
	F002V string  //币种编码	varchar
	F003V string  //币种	varchar
	F004N float64 //汇率(报表货币兑港元)	decimal
	F005N float64 //每股基本盈利（港元）	decimal
	F006N float64 //每股摊薄盈利（港元）	decimal
	F007N float64 //总资产	decimal
	F008N float64 //非流动资产	decimal
	F009N float64 //流动资产	decimal
	F010N float64 //流动负债	decimal
	F011N float64 //非流动负债	decimal
	F012N float64 //总债项	decimal
	F013N float64 //少数股东权益	decimal
	F014N float64 //股东权益/(资产净值)	decimal
	F015N float64 //股本	decimal
	F016N float64 //总储备	decimal
	F017N float64 //营业额/（银行经营收入）	decimal
	F018N float64 //经营溢利	decimal
	F019N float64 //非经常性项目收益	decimal
	F020N float64 //融资成本/（财务费用）	decimal
	F021N float64 //应占联营公司溢利	decimal
	F022N float64 //除税前经营溢利/(亏损)	decimal
	F023N float64 //税项	decimal
	F024N float64 //除税后经营溢利/(亏损)	decimal
	F025N float64 //少数股东损益	decimal
	F026N float64 //经调整股东应占溢利	decimal
	F027N float64 //相对上期增减	decimal
	F028N float64 //股息总额	decimal
	F029N float64 //扣除非经常性项目后溢利	decimal
	F030N float64 //经营业务收到的现金/经营活动之现金流量（经调整）	decimal
	F031N float64 //投资活动之现金流量（经调整）	decimal
	F032N float64 //融资活动之现金流量（经调整）	decimal
	F033N float64 //现金流量净额	decimal
	F034N float64 //年初之现金及现金等价物	decimal
	F035N float64 //外币汇率变动之影响/其他	decimal
	F036N float64 //年终之现金及现金等价物	decimal
	F037N float64 //毛利率	decimal
	F038N float64 //税前利润率	decimal
	F039N float64 //销售净利率	decimal
	F040N float64 //产权比率（负债权益比）	decimal
	F041N float64 //资产报酬率	decimal
	F042N float64 //权益报酬率	decimal
	F043N float64 //流动比率	decimal
	F044N float64 //速动比率	decimal
	F045N float64 //现金与流动资产比率	decimal
	F046N float64 //资产负债率	decimal
	F047N float64 //有形资产负债率	decimal
	F048N float64 //长期负债比率	decimal
	F049N float64 //总资产周转率	decimal
	F050N float64 //流动资产周转率	decimal
	F051N float64 //应收帐款周转率	decimal
	F052N float64 //存货周转率	decimal
	F053N float64 //EBIT	decimal
	F054N float64 //EBITD	decimal
	F055N float64 //EBITDA	decimal
	F056N float64 //息税折旧前利润率	decimal
	F057N float64 //息税折旧摊销前利润率	decimal
	F058N float64 //长期负债权益比率	decimal
	F059N float64 //营业费用比例	decimal
	F060N float64 //管理费用比例	decimal
	F061N float64 //财务费用比例	decimal
	F062N float64 //现金流动负债比	decimal
	F063N float64 //非经常性损益率	decimal
	F064N float64 //短期负债现金保障率	decimal
	F065N float64 //销售现金比率	decimal
	F066N float64 //全部资产现金回收率	decimal
	F067N float64 //经营现金净流量与净利润的比率	decimal
	F068N float64 //业务利润	decimal
	F069N float64 //股东权益合计	decimal
	F070N float64 //未分配利润	decimal
	F071N float64 //毛利润	decimal
	F072V string  //核数师意见	varchar
	F073N float64 //平均股东权益回报率	decimal
	F074N float64 //平均资产回报率	decimal
	F075N float64 //经调整每股现金流量	decimal
	F076N float64 //期末股价/每股资产净值	decimal
	F077N float64 //现金比率 (倍)	decimal
	F078V string  //报表类型编码	char
	F079N float64 //固定资产周转率	decimal
	F080N float64 //流动资产周转天数	decimal
	F081N float64 //总资产周转天数	decimal
	F082N float64 //股东权益周转率	decimal
	F083N float64 //利息保障倍数	decimal
	F084N float64 //应收帐款周转天数	decimal
	F085N float64 //存货占比	decimal
	F086N float64 //股东权益增长率	decimal
	F087N float64 //总负债增长率	decimal
	F088N float64 //经营溢利(营业利润)增长率	decimal
	F089N float64 //固定资产增长率	decimal
	F090N float64 //总资产增长率	decimal
	F091N float64 //毛利润增长率	decimal
	F092N float64 //营业收入增长率	decimal
	F093N float64 //净利润增长率	decimal
	USD   float64 //报告货币兑美元	double		报表截止日期当天汇率
	CNY   float64 //报告货币兑人民币	double		报表截止日期当天汇率
}
