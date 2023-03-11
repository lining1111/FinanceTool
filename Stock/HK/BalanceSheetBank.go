package HK

const APIBSBank = "http://webapi.cninfo.com.cn/api/hk/p_hk4020"

//BalanceSheetBank 资产负债表银行 "http://webapi.cninfo.com.cn/api/hk/p_hk4020"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询报告期	string
//			edate	结束查询报告期	string
//			type	合并类型	string	其中1代表半年报，2代表年报
type BalanceSheetBank struct {
	SECCODE     string `excel:"证券代码"` //证券代码
	SECNAME     string `excel:"证券简称"` //证券简称
	DECLAREDATE string `excel:"公布日期"` //公布日期
	STARTDATE   string `excel:"起始日期"` //起始日期
	ENDDATE     string `excel:"截止日期"` //截止日期

	F001V string  `excel:"报表类别"`            //报表类别
	F002V string  `excel:"币种编码"`            //币种编码
	F003V string  `excel:"币种"`              //币种
	F004N float64 `excel:"汇率(报表货币兑港元)"`     //汇率(报表货币兑港元)
	F005N float64 `excel:"库存现金及短期资金"`       //库存现金及短期资金
	F006N float64 `excel:"定期存放银行同业及其他财务机构"` //定期存放银行同业及其他财务机构
	F007N float64 `excel:"交易用途资产"`          //交易用途资产
	F008N float64 `excel:"商业票据"`            //商业票据
	F009N float64 `excel:"持有之存款证"`          //持有之存款证
	F010N float64 `excel:"衍生工具"`            //衍生工具
	F011N float64 `excel:"香港政府负债证明书"`       //香港政府负债证明书
	F012N float64 `excel:"可供出售之金融资产"`       //可供出售之金融资产
	F013N float64 `excel:"持至到期投资"`          //持至到期投资
	F014N float64 `excel:"公平订值之金融资产"`       //公平订值之金融资产
	F015N float64 `excel:"联营及合资公司权益"`       //联营及合资公司权益
	F016N float64 `excel:"金融投资"`            //金融投资
	F017N float64 `excel:"固定资产"`            //固定资产
	F018N float64 `excel:"商誉及无形资产"`         //商誉及无形资产
	F019N float64 `excel:"向其他银行托收中之项目"`     //向其他银行托收中之项目
	F020N float64 `excel:"银行同业贷款及垫款"`       //银行同业贷款及垫款
	F021N float64 `excel:"客户贷款及垫款"`         //客户贷款及垫款
	F022N float64 `excel:"非交易用途资产"`         //非交易用途资产
	F023N float64 `excel:"其他资产"`            //其他资产
	F024N float64 `excel:"总资产"`             //总资产
	F025N float64 `excel:"同业及财务机构存款"`       //同业及财务机构存款
	F026N float64 `excel:"客户存款"`            //客户存款
	F027N float64 `excel:"衍生工具1"`           //衍生工具1
	F028N float64 `excel:"公平订值之金融负债"`       //公平订值之金融负债
	F029N float64 `excel:"香港纸币流通额"`         //香港纸币流通额
	F030N float64 `excel:"向其他银行传送中之项目"`     //向其他银行传送中之项目
	F031N float64 `excel:"已发行之债务项目"`        //已发行之债务项目
	F032N float64 `excel:"其他负债"`            //其他负债
	F033N float64 `excel:"后偿负债"`            //后偿负债
	F034N float64 `excel:"交易用途负债"`          //交易用途负债
	F035N float64 `excel:"总负债"`             //总负债
	F036N float64 `excel:"股本"`              //股本
	F037N float64 `excel:"股份溢价"`            //股份溢价
	F038N float64 `excel:"资本储备"`            //资本储备
	F039N float64 `excel:"其他储备"`            //其他储备
	F040N float64 `excel:"保留溢利/(累计亏损)"`     //保留溢利/(累计亏损)
	F041N float64 `excel:"总储备"`             //总储备
	F042N float64 `excel:"股东资金"`            //股东资金
	F043N float64 `excel:"少数股东权益"`          //少数股东权益
	F044N float64 `excel:"各类股东权益总额"`        //各类股东权益总额
	F045N float64 `excel:"各类股东权益及负债总额"`     //各类股东权益及负债总额
	F046V string  `excel:"核数师意见"`           //核数师意见
	F047V string  `excel:"报表类型编码"`          //报表类型编码 其中1代表半年报，2代表年报

	MEMO string  `excel:"备注"`       //备注
	USD  float64 `excel:"报告货币兑美元"`  //报告货币兑美元 报表截止日期当天汇率
	CNY  float64 `excel:"报告货币兑人民币"` //报告货币兑人民币 报表截止日期当天汇率
}
