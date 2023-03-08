package HK

const apiIS = "http://webapi.cninfo.com.cn/api/hk/p_hk4024"

//IncomeSheet 综合损益表(非银行) "http://webapi.cninfo.com.cn/api/hk/p_hk4024"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询时间	string	否	可为空 样例：20160101 或者 2016-01-01 或者 2016/01/01
//			edate	结束查询时间	string	否	可为空 样例：20160101 或者 2016-01-01 或者 2016/01/01
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type IncomeSheet struct {
	SECCODE     string //证券代码	varchar
	SECNAME     string //证券简称	varchar
	DECLAREDATE string //公布日期	date
	STARTDATE   string //起始日期	date
	ENDDATE     string //截止日期	date

	F001V string  //报表类别	varchar
	F002V string  //币种编码	varchar
	F003V string  //币种	varchar
	F004N float64 //汇率(报表货币兑港元)	numeric
	F005N float64 //营业额	numeric
	F006N float64 //经营溢利	numeric
	F007N float64 //特殊项目	numeric
	F008N float64 //联营公司	numeric
	F009N float64 //除税前经营溢利	numeric
	F010N float64 //税项	numeric
	F011N float64 //少数股东损益	numeric
	F012N float64 //经调整股东应占溢利	numeric
	F013N float64 //折旧	numeric
	F014N float64 //贷款利息	numeric
	F015N float64 //利息拨作发展资本	numeric
	F016N float64 //营业额增长	numeric
	F017N float64 //经调整股东应占溢利增长	numeric
	F018N float64 //税率	numeric
	F019N float64 //每股盈利（港元）	numeric
	F020N float64 //每股派息（港元）	numeric
	F021V string  //核数师意见	varchar
	F022N float64 //发行股数(年度)	numeric
	F023N float64 //优先股股息	numeric
	F024N float64 //报告股东应占溢利	numeric
	F025N float64 //每股盈利（报告货币）	numeric
	F026N float64 //每股派息（报告货币）	numeric
	F027V string  //报告每股派息币种编码	varchar
	F028V string  //报告每股派息币种	varchar
	F029N float64 //加权平均普通股	numeric
	F030N float64 //经营开支总额	numeric
	F031N float64 //销售及分销成本	numeric
	F032N float64 //行政费用	numeric
	F033N float64 //财务成本	numeric
	F034N float64 //每股资产净值(经调整)	numeric
	F035N float64 //经调整每股资本公积	numeric
	F036V string  //报表类型编码	char		其中1代表半年报，2代表年报

	MEMO string  //备注	varchar
	USD  float64 //报告货币兑美元	double		报表截止日期当天汇率
	CNY  float64 //报告货币兑人民币	double		报表截止日期当天汇率
}
