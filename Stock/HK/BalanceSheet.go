package HK

const APIBS = "http://webapi.cninfo.com.cn/api/hk/p_hk4023"

//BalanceSheet 资产负债表 "http://webapi.cninfo.com.cn/api/hk/p_hk4023"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询报告期	string
//			edate	结束查询报告期	string
//			type	合并类型	string	其中1代表半年报，2代表年报
type BalanceSheet struct {
	SECCODE     string //证券代码	varchar
	SECNAME     string //证券简称	varchar
	DECLAREDATE string //公布日期	date
	STARTDATE   string //起始日期	date
	ENDDATE     string //截止日期	date

	F001V string  //报表类别	varchar
	F002V string  //币种编码	varchar
	F003V string  //币种	varchar
	F004N float64 //汇率(报表货币兑港元)	numeric
	F005N float64 //固定资产	numeric
	F006N float64 //投资	numeric
	F007N float64 //其他非流动资产	numeric
	F008N float64 //非流动资产总值	numeric
	F009N float64 //流动资产总值	numeric
	F010N float64 //其中：存货	numeric
	F011N float64 //其中：现金及银行结存	numeric
	F012N float64 //总资产	numeric
	F013N float64 //短期债项	numeric
	F014N float64 //流动负债总值	numeric
	F015N float64 //长期债项	numeric
	F016N float64 //其他非流动负债及少数股东权益	numeric
	F017N float64 //总债项	numeric
	F018N float64 //股本	numeric
	F019N float64 //储备	numeric
	F020N float64 //股东权益	numeric
	F021V string  //核数师意见	varchar
	F022N float64 //营运资金	numeric
	F023V string  //报表类型编码	char		其中1代表半年报，2代表年报

	MEMO string  //备注	varchar
	USD  float64 //报告货币兑美元	double		报表截止日期当天汇率
	CNY  float64 //报告货币兑人民币	double		报表截止日期当天汇率
}
