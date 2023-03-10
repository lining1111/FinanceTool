package HK

const APIISBank = "http://webapi.cninfo.com.cn/api/hk/p_hk4021"

//IncomeSheetBank 综合损益表(银行) "http://webapi.cninfo.com.cn/api/hk/p_hk4021"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询时间	string	否	可为空 样例：20160101 或者 2016-01-01 或者 2016/01/01
//			edate	结束查询时间	string	否	可为空 样例：20160101 或者 2016-01-01 或者 2016/01/01
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type IncomeSheetBank struct {
	SECCODE     string `excel:"证券代码"` //证券代码	varchar
	SECNAME     string `excel:"证券简称"` //证券简称	varchar
	DECLAREDATE string `excel:"公布日期"` //公布日期	date
	STARTDATE   string `excel:"起始日期"` //起始日期	date
	ENDDATE     string `excel:"截止日期"` //截止日期	date

	F001V string  `excel:"报表类别"`                //报表类别	varchar
	F002V string  `excel:"币种编码"`                //币种编码	varchar
	F003V string  `excel:"币种"`                  //币种	varchar
	F004N float64 `excel:"汇率(报表货币兑港元)"`         //汇率(报表货币兑港元)	numeric
	F005N float64 `excel:"利息收益"`                //利息收益	numeric
	F006N float64 `excel:"利息支出"`                //利息支出	numeric
	F007N float64 `excel:"净利息收益"`               //净利息收益	numeric
	F008N float64 `excel:"费用收益净额"`              //费用收益净额	numeric
	F009N float64 `excel:"交易收益净额"`              //交易收益净额	numeric
	F010N float64 `excel:"其他营业收益"`              //其他营业收益	numeric
	F011N float64 `excel:"营业收益总额"`              //营业收益总额	numeric
	F012N float64 `excel:"保险索偿净额及对保单持有人负债之变动"`  //保险索偿净额及对保单持有人负债之变动	numeric
	F013N float64 `excel:"未扣除贷款减值及其他准备之营业收益净额"` //未扣除贷款减值及其他准备之营业收益净额	numeric
	F014N float64 `excel:"贷款减值及其他信贷风险准备"`       //贷款减值及其他信贷风险准备	numeric
	F015N float64 `excel:"营业收益净额"`              //营业收益净额	numeric
	F016N float64 `excel:"营业支出总额"`              //营业支出总额	numeric
	F017N float64 `excel:"营业利润/(亏损)"`           //营业利润/(亏损)	numeric
	F018N float64 `excel:"非经营/特殊项目"`            //非经营/特殊项目	numeric
	F019N float64 `excel:"联营公司"`                //联营公司	numeric
	F020N float64 `excel:"除税前经营溢利/(亏损)"`        //除税前经营溢利/(亏损)	numeric
	F021N float64 `excel:"税项"`                  //税项	numeric
	F022N float64 `excel:"除税后经营溢利/(亏损)"`        //除税后经营溢利/(亏损)	numeric
	F023N float64 `excel:"少数股东损益"`              //少数股东损益	numeric
	F024N float64 `excel:"优先股股息"`               //优先股股息	numeric
	F025N float64 `excel:"经调整股东应占溢利"`           //经调整股东应占溢利	numeric
	F026N float64 `excel:"股息总额"`                //股息总额	numeric
	F027N float64 `excel:"保留溢利/(亏损)"`           //保留溢利/(亏损)	numeric
	F028N float64 `excel:"每股盈利（港元）"`            //每股盈利（港元）	numeric
	F029N float64 `excel:"每股派息（港元）"`            //每股派息（港元）	numeric
	F030V string  `excel:"核数师意见"`               //核数师意见	varchar
	F031N uint64  `excel:"发行股数(年度)"`            //发行股数(年度)	numeric
	F032N float64 `excel:"资本证券之应付票息（报表货币兑港元）"`  //资本证券之应付票息（报表货币兑港元）	numeric
	F033N float64 `excel:"报告股东应占溢利"`            //报告股东应占溢利	numeric
	F034N float64 `excel:"每股盈利（报告货币）"`          //每股盈利（报告货币）	numeric
	F035N float64 `excel:"每股派息（报告货币）"`          //每股派息（报告货币）	numeric
	F036V string  `excel:"报告每股派息币种编码"`          //报告每股派息币种编码	varchar
	F037V string  `excel:"报告每股派息币种"`            //报告每股派息币种	varchar
	F038N float64 `excel:"加权平均普通股"`             //加权平均普通股	numeric
	F039N float64 `excel:"经营开支总额"`              //经营开支总额	numeric
	F040N float64 `excel:"每股资产净值(经调整)"`         //每股资产净值(经调整)	numeric
	F041N float64 `excel:"经调整每股资本公积"`           //经调整每股资本公积	numeric
	F042V string  `excel:"报表类型编码"`              //报表类型编码	char		其中1代表半年报，2代表年报

	MEMO string  `excel:"备注"`       //备注	varchar
	USD  float64 `excel:"报告货币兑美元"`  //报告货币兑美元	double		报表截止日期当天汇率
	CNY  float64 `excel:"报告货币兑人民币"` //报告货币兑人民币	double		报表截止日期当天汇率
}
