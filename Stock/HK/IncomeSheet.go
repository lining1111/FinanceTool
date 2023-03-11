package HK

const APIIS = "http://webapi.cninfo.com.cn/api/hk/p_hk4024"

//IncomeSheet 综合损益表(非银行) "http://webapi.cninfo.com.cn/api/hk/p_hk4024"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询时间	string	否	可为空 样例：20160101 或者 2016-01-01 或者 2016/01/01
//			edate	结束查询时间	string	否	可为空 样例：20160101 或者 2016-01-01 或者 2016/01/01
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type IncomeSheet struct {
	SECCODE     string `excel:"证券代码"` //证券代码
	SECNAME     string `excel:"证券简称"` //证券简称
	DECLAREDATE string `excel:"公布日期"` //公布日期
	STARTDATE   string `excel:"起始日期"` //起始日期
	ENDDATE     string `excel:"截止日期"` //截止日期

	F001V string  `excel:"报表类别"`        //报表类别
	F002V string  `excel:"币种编码"`        //币种编码
	F003V string  `excel:"币种"`          //币种
	F004N float64 `excel:"汇率(报表货币兑港元)"` //汇率(报表货币兑港元)
	F005N float64 `excel:"营业额"`         //营业额
	F006N float64 `excel:"经营溢利"`        //经营溢利
	F007N float64 `excel:"特殊项目"`        //特殊项目
	F008N float64 `excel:"联营公司"`        //联营公司
	F009N float64 `excel:"除税前经营溢利"`     //除税前经营溢利
	F010N float64 `excel:"税项"`          //税项
	F011N float64 `excel:"少数股东损益"`      //少数股东损益
	F012N float64 `excel:"经调整股东应占溢利"`   //经调整股东应占溢利
	F013N float64 `excel:"折旧"`          //折旧
	F014N float64 `excel:"贷款利息"`        //贷款利息
	F015N float64 `excel:"利息拨作发展资本"`    //利息拨作发展资本
	F016N float64 `excel:"营业额增长"`       //营业额增长
	F017N float64 `excel:"经调整股东应占溢利增长"` //经调整股东应占溢利增长
	F018N float64 `excel:"税率"`          //税率
	F019N float64 `excel:"每股盈利（港元）"`    //每股盈利（港元）
	F020N float64 `excel:"每股派息（港元）"`    //每股派息（港元）
	F021V string  `excel:"核数师意见"`       //核数师意见
	F022N float64 `excel:"发行股数(年度)"`    //发行股数(年度)
	F023N float64 `excel:"优先股股息"`       //优先股股息
	F024N float64 `excel:"报告股东应占溢利"`    //报告股东应占溢利
	F025N float64 `excel:"每股盈利（报告货币）"`  //每股盈利（报告货币）
	F026N float64 `excel:"每股派息（报告货币）"`  //每股派息（报告货币）
	F027V string  `excel:"报告每股派息币种编码"`  //报告每股派息币种编码
	F028V string  `excel:"报告每股派息币种"`    //报告每股派息币种
	F029N float64 `excel:"加权平均普通股"`     //加权平均普通股
	F030N float64 `excel:"经营开支总额"`      //经营开支总额
	F031N float64 `excel:"销售及分销成本"`     //销售及分销成本
	F032N float64 `excel:"行政费用"`        //行政费用
	F033N float64 `excel:"财务成本"`        //财务成本
	F034N float64 `excel:"每股资产净值(经调整)"` //每股资产净值(经调整)
	F035N float64 `excel:"经调整每股资本公积"`   //经调整每股资本公积
	F036V string  `excel:"报表类型编码"`      //报表类型编码 其中1代表半年报，2代表年报

	MEMO string  `excel:"备注"`       //备注
	USD  float64 `excel:"报告货币兑美元"`  //报告货币兑美元 报表截止日期当天汇率
	CNY  float64 `excel:"报告货币兑人民币"` //报告货币兑人民币 报表截止日期当天汇率
}
