package HK

const APIISBank = "http://webapi.cninfo.com.cn/api/hk/p_hk4021"

//IncomeSheetBank 综合损益表(银行) "http://webapi.cninfo.com.cn/api/hk/p_hk4021"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询时间	string	否	可为空 样例：20160101 或者 2016-01-01 或者 2016/01/01
//			edate	结束查询时间	string	否	可为空 样例：20160101 或者 2016-01-01 或者 2016/01/01
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type IncomeSheetBank struct {
	SECCODE     string `excel:"证券代码"` //证券代码
	SECNAME     string `excel:"证券简称"` //证券简称
	DECLAREDATE string `excel:"公布日期"` //公布日期
	STARTDATE   string `excel:"起始日期"` //起始日期
	ENDDATE     string `excel:"截止日期"` //截止日期

	F001V string  `excel:"报表类别"`                //报表类别
	F002V string  `excel:"币种编码"`                //币种编码
	F003V string  `excel:"币种"`                  //币种
	F004N float64 `excel:"汇率(报表货币兑港元)"`         //汇率(报表货币兑港元)
	F005N float64 `excel:"利息收益"`                //利息收益
	F006N float64 `excel:"利息支出"`                //利息支出
	F007N float64 `excel:"净利息收益"`               //净利息收益
	F008N float64 `excel:"费用收益净额"`              //费用收益净额
	F009N float64 `excel:"交易收益净额"`              //交易收益净额
	F010N float64 `excel:"其他营业收益"`              //其他营业收益
	F011N float64 `excel:"营业收益总额"`              //营业收益总额
	F012N float64 `excel:"保险索偿净额及对保单持有人负债之变动"`  //保险索偿净额及对保单持有人负债之变动
	F013N float64 `excel:"未扣除贷款减值及其他准备之营业收益净额"` //未扣除贷款减值及其他准备之营业收益净额
	F014N float64 `excel:"贷款减值及其他信贷风险准备"`       //贷款减值及其他信贷风险准备
	F015N float64 `excel:"营业收益净额"`              //营业收益净额
	F016N float64 `excel:"营业支出总额"`              //营业支出总额
	F017N float64 `excel:"营业利润/(亏损)"`           //营业利润/(亏损)
	F018N float64 `excel:"非经营/特殊项目"`            //非经营/特殊项目
	F019N float64 `excel:"联营公司"`                //联营公司
	F020N float64 `excel:"除税前经营溢利/(亏损)"`        //除税前经营溢利/(亏损)
	F021N float64 `excel:"税项"`                  //税项
	F022N float64 `excel:"除税后经营溢利/(亏损)"`        //除税后经营溢利/(亏损)
	F023N float64 `excel:"少数股东损益"`              //少数股东损益
	F024N float64 `excel:"优先股股息"`               //优先股股息
	F025N float64 `excel:"经调整股东应占溢利"`           //经调整股东应占溢利
	F026N float64 `excel:"股息总额"`                //股息总额
	F027N float64 `excel:"保留溢利/(亏损)"`           //保留溢利/(亏损)
	F028N float64 `excel:"每股盈利（港元）"`            //每股盈利（港元）
	F029N float64 `excel:"每股派息（港元）"`            //每股派息（港元）
	F030V string  `excel:"核数师意见"`               //核数师意见
	F031N uint64  `excel:"发行股数(年度)"`            //发行股数(年度)
	F032N float64 `excel:"资本证券之应付票息（报表货币兑港元）"`  //资本证券之应付票息（报表货币兑港元）
	F033N float64 `excel:"报告股东应占溢利"`            //报告股东应占溢利
	F034N float64 `excel:"每股盈利（报告货币）"`          //每股盈利（报告货币）
	F035N float64 `excel:"每股派息（报告货币）"`          //每股派息（报告货币）
	F036V string  `excel:"报告每股派息币种编码"`          //报告每股派息币种编码
	F037V string  `excel:"报告每股派息币种"`            //报告每股派息币种
	F038N float64 `excel:"加权平均普通股"`             //加权平均普通股
	F039N float64 `excel:"经营开支总额"`              //经营开支总额
	F040N float64 `excel:"每股资产净值(经调整)"`         //每股资产净值(经调整)
	F041N float64 `excel:"经调整每股资本公积"`           //经调整每股资本公积
	F042V string  `excel:"报表类型编码"`              //报表类型编码 其中1代表半年报，2代表年报

	MEMO string  `excel:"备注"`       //备注
	USD  float64 `excel:"报告货币兑美元"`  //报告货币兑美元 报表截止日期当天汇率
	CNY  float64 `excel:"报告货币兑人民币"` //报告货币兑人民币 报表截止日期当天汇率
}
