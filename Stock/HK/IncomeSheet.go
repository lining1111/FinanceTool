package HK

const APIIS = "http://webapi.cninfo.com.cn/api/hk/p_hk4024"

//IncomeSheet 综合损益表(非银行) "http://webapi.cninfo.com.cn/api/hk/p_hk4024"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询时间	string	否	可为空 样例：20160101 或者 2016-01-01 或者 2016/01/01
//			edate	结束查询时间	string	否	可为空 样例：20160101 或者 2016-01-01 或者 2016/01/01
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type IncomeSheet struct {
	SECCODE     string `excel:"证券代码"` //证券代码	varchar
	SECNAME     string `excel:"证券简称"` //证券简称	varchar
	DECLAREDATE string `excel:"公布日期"` //公布日期	date
	STARTDATE   string `excel:"起始日期"` //起始日期	date
	ENDDATE     string `excel:"截止日期"` //截止日期	date

	F001V string  `excel:"报表类别"`        //报表类别	varchar
	F002V string  `excel:"币种编码"`        //币种编码	varchar
	F003V string  `excel:"币种"`          //币种	varchar
	F004N float64 `excel:"汇率(报表货币兑港元)"` //汇率(报表货币兑港元)	numeric
	F005N float64 `excel:"营业额"`         //营业额	numeric
	F006N float64 `excel:"经营溢利"`        //经营溢利	numeric
	F007N float64 `excel:"特殊项目"`        //特殊项目	numeric
	F008N float64 `excel:"联营公司"`        //联营公司	numeric
	F009N float64 `excel:"除税前经营溢利"`     //除税前经营溢利	numeric
	F010N float64 `excel:"税项"`          //税项	numeric
	F011N float64 `excel:"少数股东损益"`      //少数股东损益	numeric
	F012N float64 `excel:"经调整股东应占溢利"`   //经调整股东应占溢利	numeric
	F013N float64 `excel:"折旧"`          //折旧	numeric
	F014N float64 `excel:"贷款利息"`        //贷款利息	numeric
	F015N float64 `excel:"利息拨作发展资本"`    //利息拨作发展资本	numeric
	F016N float64 `excel:"营业额增长"`       //营业额增长	numeric
	F017N float64 `excel:"经调整股东应占溢利增长"` //经调整股东应占溢利增长	numeric
	F018N float64 `excel:"税率"`          //税率	numeric
	F019N float64 `excel:"每股盈利（港元）"`    //每股盈利（港元）	numeric
	F020N float64 `excel:"每股派息（港元）"`    //每股派息（港元）	numeric
	F021V string  `excel:"核数师意见"`       //核数师意见	varchar
	F022N float64 `excel:"发行股数(年度)"`    //发行股数(年度)	numeric
	F023N float64 `excel:"优先股股息"`       //优先股股息	numeric
	F024N float64 `excel:"报告股东应占溢利"`    //报告股东应占溢利	numeric
	F025N float64 `excel:"每股盈利（报告货币）"`  //每股盈利（报告货币）	numeric
	F026N float64 `excel:"每股派息（报告货币）"`  //每股派息（报告货币）	numeric
	F027V string  `excel:"报告每股派息币种编码"`  //报告每股派息币种编码	varchar
	F028V string  `excel:"报告每股派息币种"`    //报告每股派息币种	varchar
	F029N float64 `excel:"加权平均普通股"`     //加权平均普通股	numeric
	F030N float64 `excel:"经营开支总额"`      //经营开支总额	numeric
	F031N float64 `excel:"销售及分销成本"`     //销售及分销成本	numeric
	F032N float64 `excel:"行政费用"`        //行政费用	numeric
	F033N float64 `excel:"财务成本"`        //财务成本	numeric
	F034N float64 `excel:"每股资产净值(经调整)"` //每股资产净值(经调整)	numeric
	F035N float64 `excel:"经调整每股资本公积"`   //经调整每股资本公积	numeric
	F036V string  `excel:"报表类型编码"`      //报表类型编码	char		其中1代表半年报，2代表年报

	MEMO string  `excel:"备注"`       //备注	varchar
	USD  float64 `excel:"报告货币兑美元"`  //报告货币兑美元	double		报表截止日期当天汇率
	CNY  float64 `excel:"报告货币兑人民币"` //报告货币兑人民币	double		报表截止日期当天汇率
}
