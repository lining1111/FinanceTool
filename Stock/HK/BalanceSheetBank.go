package HK

const apiBSBank = "http://webapi.cninfo.com.cn/api/hk/p_hk4020"

//BalanceSheetBank 资产负债表银行 "http://webapi.cninfo.com.cn/api/hk/p_hk4020"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询报告期	string
//			edate	结束查询报告期	string
//			type	合并类型	string	其中1代表半年报，2代表年报
type BalanceSheetBank struct {
	SECCODE     string //证券代码	varchar
	SECNAME     string //证券简称	varchar
	DECLAREDATE string //公布日期	date
	STARTDATE   string //起始日期	date
	ENDDATE     string //截止日期	date

	F001V string  //报表类别	varchar
	F002V string  //币种编码	varchar
	F003V string  //币种	varchar
	F004N float64 //汇率(报表货币兑港元)	numeric
	F005N float64 //库存现金及短期资金	numeric
	F006N float64 //定期存放银行同业及其他财务机构	numeric
	F007N float64 //交易用途资产	numeric
	F008N float64 //商业票据	numeric
	F009N float64 //持有之存款证	numeric
	F010N float64 //衍生工具	numeric
	F011N float64 //香港政府负债证明书	numeric
	F012N float64 //可供出售之金融资产	numeric
	F013N float64 //持至到期投资	numeric
	F014N float64 //公平订值之金融资产	numeric
	F015N float64 //联营及合资公司权益	numeric
	F016N float64 //金融投资	numeric
	F017N float64 //固定资产	numeric
	F018N float64 //商誉及无形资产	numeric
	F019N float64 //向其他银行托收中之项目	numeric
	F020N float64 //银行同业贷款及垫款	numeric
	F021N float64 //客户贷款及垫款	numeric
	F022N float64 //非交易用途资产	numeric
	F023N float64 //其他资产	numeric
	F024N float64 //总资产	numeric
	F025N float64 //同业及财务机构存款	numeric
	F026N float64 //客户存款	numeric
	F027N float64 //衍生工具1	numeric
	F028N float64 //公平订值之金融负债	numeric
	F029N float64 //香港纸币流通额	numeric
	F030N float64 //向其他银行传送中之项目	numeric
	F031N float64 //已发行之债务项目	numeric
	F032N float64 //其他负债	numeric
	F033N float64 //后偿负债	numeric
	F034N float64 //交易用途负债	numeric
	F035N float64 //总负债	numeric
	F036N float64 //股本	numeric
	F037N float64 //股份溢价	numeric
	F038N float64 //资本储备	numeric
	F039N float64 //其他储备	numeric
	F040N float64 //保留溢利/(累计亏损)	numeric
	F041N float64 //总储备	numeric
	F042N float64 //股东资金	numeric
	F043N float64 //少数股东权益	numeric
	F044N float64 //各类股东权益总额	numeric
	F045N float64 //各类股东权益及负债总额	numeric
	F046V string  //核数师意见	varchar
	F047V string  //报表类型编码	char		其中1代表半年报，2代表年报

	MEMO string  //备注	varchar
	USD  float64 //报告货币兑美元	double		报表截止日期当天汇率
	CNY  float64 //报告货币兑人民币	double		报表截止日期当天汇率
}
