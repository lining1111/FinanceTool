package HK

const APIBS = "http://webapi.cninfo.com.cn/api/hk/p_hk4023"

//BalanceSheet 资产负债表 "http://webapi.cninfo.com.cn/api/hk/p_hk4023"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询报告期	string
//			edate	结束查询报告期	string
//			type	合并类型	string	其中1代表半年报，2代表年报
type BalanceSheet struct {
	SECCODE     string `excel:"证券代码"` //证券代码
	SECNAME     string `excel:"证券简称"` //证券简称
	DECLAREDATE string `excel:"公布日期"` //公布日期
	STARTDATE   string `excel:"起始日期"` //起始日期
	ENDDATE     string `excel:"截止日期"` //截止日期

	F001V string  `excel:"报表类别"`           //报表类别
	F002V string  `excel:"币种编码"`           //币种编码
	F003V string  `excel:"币种"`             //币种
	F004N float64 `excel:"汇率(报表货币兑港元)"`    //汇率(报表货币兑港元)
	F005N float64 `excel:"固定资产"`           //固定资产
	F006N float64 `excel:"投资"`             //投资
	F007N float64 `excel:"其他非流动资产"`        //其他非流动资产
	F008N float64 `excel:"非流动资产总值"`        //非流动资产总值
	F009N float64 `excel:"流动资产总值"`         //流动资产总值
	F010N float64 `excel:"其中：存货"`          //其中：存货
	F011N float64 `excel:"其中：现金及银行结存"`     //其中：现金及银行结存
	F012N float64 `excel:"总资产"`            //总资产
	F013N float64 `excel:"短期债项"`           //短期债项
	F014N float64 `excel:"流动负债总值"`         //流动负债总值
	F015N float64 `excel:"长期债项"`           //长期债项
	F016N float64 `excel:"其他非流动负债及少数股东权益"` //其他非流动负债及少数股东权益
	F017N float64 `excel:"总债项"`            //总债项
	F018N float64 `excel:"股本"`             //股本
	F019N float64 `excel:"储备"`             //储备
	F020N float64 `excel:"股东权益"`           //股东权益
	F021V string  `excel:"核数师意见"`          //核数师意见
	F022N float64 `excel:"营运资金"`           //营运资金
	F023V string  `excel:"报表类型编码"`         //报表类型编码 其中1代表半年报，2代表年报

	MEMO string  `excel:"备注"`       //备注
	USD  float64 `excel:"报告货币兑美元"`  //报告货币兑美元 报表截止日期当天汇率
	CNY  float64 `excel:"报告货币兑人民币"` //报告货币兑人民币 报表截止日期当天汇率
}
