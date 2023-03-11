package HK

const APIFRBank = "http://webapi.cninfo.com.cn/api/hk/p_hk4022"

//FinancialRatiosBank 重要指标表（银行） "http://webapi.cninfo.com.cn/api/hk/p_hk4022"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type FinancialRatiosBank struct {
	SECCODE     string `excel:"证券代码"` //证券代码
	SECNAME     string `excel:"证券简称"` //证券简称
	DECLAREDATE string `excel:"公布日期"` //公布日期
	ENDDATE     string `excel:"截止日期"` //截止日期

	F001V string  `excel:"报表类别"`                     //报表类别
	F002V string  `excel:"币种编码"`                     //币种编码
	F003V string  `excel:"币种"`                       //币种
	F004N float64 `excel:"汇率(报表货币兑港元)"`              //汇率(报表货币兑港元)
	F005N float64 `excel:"每股基本盈利（港元）"`               //每股基本盈利（港元）
	F006N float64 `excel:"每股摊薄盈利（港元）"`               //每股摊薄盈利（港元）
	F008N float64 `excel:"客户贷款及垫款"`                  //客户贷款及垫款
	F007N float64 `excel:"银行同业贷款及垫款"`                //银行同业贷款及垫款
	F009N float64 `excel:"定期存放银行同业及其他财务机构"`          //定期存放银行同业及其他财务机构
	F010N float64 `excel:"总资产"`                      //总资产
	F011N float64 `excel:"同业及财务机构存款"`                //同业及财务机构存款
	F012N float64 `excel:"客户存款"`                     //客户存款
	F013N float64 `excel:"总负债"`                      //总负债
	F014N float64 `excel:"少数股东权益"`                   //少数股东权益
	F015N float64 `excel:"股东权益/(资产净值)"`              //股东权益/(资产净值)
	F016N float64 `excel:"股本"`                       //股本
	F017N float64 `excel:"总储备"`                      //总储备
	F018N float64 `excel:"净利息收益"`                    //净利息收益
	F019N float64 `excel:"费用收益净额"`                   //费用收益净额
	F020N float64 `excel:"交易收益净额"`                   //交易收益净额
	F021N float64 `excel:"其他营业收益"`                   //其他营业收益
	F022N float64 `excel:"营业收益总额"`                   //营业收益总额
	F023N float64 `excel:"贷款减值及其他信贷风险准备"`            //贷款减值及其他信贷风险准备
	F024N float64 `excel:"营业利润/(亏损)"`                //营业利润/(亏损)
	F025N float64 `excel:"非盈利/特殊项目"`                 //非盈利/特殊项目
	F026N float64 `excel:"应占联营公司溢利"`                 //应占联营公司溢利
	F027N float64 `excel:"除税前经营溢利/(亏损)"`             //除税前经营溢利/(亏损)
	F028N float64 `excel:"税项"`                       //税项
	F029N float64 `excel:"除税后经营溢利/(亏损)"`             //除税后经营溢利/(亏损)
	F030N float64 `excel:"少数股东损益"`                   //少数股东损益
	F031N float64 `excel:"经调整股东应占溢利"`                //经调整股东应占溢利
	F032N float64 `excel:"相对上期增减"`                   //相对上期增减
	F033N float64 `excel:"股息总额"`                     //股息总额
	F034N float64 `excel:"扣除非经常性项目后溢利"`              //扣除非经常性项目后溢利
	F035N float64 `excel:"经营业务收到的现金/经营活动之现金流量（经调整）"` //经营业务收到的现金/经营活动之现金流量（经调整）
	F036N float64 `excel:"投资活动之现金流量（经调整）"`           //投资活动之现金流量（经调整）
	F037N float64 `excel:"融资活动之现金流量（经调整）"`           //融资活动之现金流量（经调整）
	F038N float64 `excel:"现金流量净额"`                   //现金流量净额
	F039N float64 `excel:"年初之现金及现金等价物"`              //年初之现金及现金等价物
	F040N float64 `excel:"外币汇率变动之影响/其他"`             //外币汇率变动之影响/其他
	F041N float64 `excel:"年终之现金及现金等价物"`              //年终之现金及现金等价物
	F042N float64 `excel:"资本充足比率"`                   //资本充足比率
	F043N float64 `excel:"流动资金比率"`                   //流动资金比率
	F044N float64 `excel:"成本/收入"`                    //成本/收入
	F045N float64 `excel:"流动资金/存款"`                  //流动资金/存款
	F046N float64 `excel:"贷款/存款"`                    //贷款/存款
	F047N float64 `excel:"贷款/股东权益"`                  //贷款/股东权益
	F048N float64 `excel:"贷款/总资产"`                   //贷款/总资产
	F049N float64 `excel:"存款/股东权益"`                  //存款/股东权益
	F050N float64 `excel:"存款/总资产"`                   //存款/总资产
	F051N float64 `excel:"贷款回报率"`                    //贷款回报率
	F052N float64 `excel:"存款回报率"`                    //存款回报率
	F053N float64 `excel:"资产回报率"`                    //资产回报率
	F054N float64 `excel:"权益报酬率"`                    //权益报酬率
	F055N float64 `excel:"股东权益合计"`                   //股东权益合计
	F056N float64 `excel:"未分配利润"`                    //未分配利润
	F057V string  `excel:"核数师意见"`                    //核数师意见
	F058N float64 `excel:"平均股东权益回报率"`                //平均股东权益回报率
	F059N float64 `excel:"平均资产回报率"`                  //平均资产回报率
	F060N float64 `excel:"经调整每股现金流量"`                //经调整每股现金流量
	F061N float64 `excel:"期末股价/每股资产净值"`              //期末股价/每股资产净值
	F062V string  `excel:"报表类型编码"`                   //报表类型编码 其中1代表半年报，2代表年报
	MEMO  string  `excel:"备注"`                       //备注
	USD   float64 `excel:"报告货币兑美元"`                  //报告货币兑美元 报表截止日期当天汇率
	CNY   float64 `excel:"报告货币兑人民币"`                 //报告货币兑人民币 报表截止日期当天汇率
}
