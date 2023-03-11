package HK

const APIFR = "http://webapi.cninfo.com.cn/api/hk/p_hk4025"

//FinancialRatios 重要指标表（非银行） "http://webapi.cninfo.com.cn/api/hk/p_hk4025"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type FinancialRatios struct {
	SECCODE     string  `excel:"证券代码"`                     //证券代码
	SECNAME     string  `excel:"证券简称"`                     //证券简称
	DECLAREDATE string  `excel:"公布日期"`                     //公布日期
	STARTDATE   string  `excel:"起始日期"`                     //起始日期
	ENDDATE     string  `excel:"截止日期"`                     //截止日期
	F001V       string  `excel:"报表类别"`                     //报表类别
	F002V       string  `excel:"币种编码"`                     //币种编码
	F003V       string  `excel:"币种"`                       //币种
	F004N       float64 `excel:"汇率(报表货币兑港元)"`              //汇率(报表货币兑港元)
	F005N       float64 `excel:"每股基本盈利（港元）"`               //每股基本盈利（港元）
	F006N       float64 `excel:"每股摊薄盈利（港元）"`               //每股摊薄盈利（港元）
	F007N       float64 `excel:"总资产"`                      //总资产
	F008N       float64 `excel:"非流动资产"`                    //非流动资产
	F009N       float64 `excel:"流动资产"`                     //流动资产
	F010N       float64 `excel:"流动负债"`                     //流动负债
	F011N       float64 `excel:"非流动负债"`                    //非流动负债
	F012N       float64 `excel:"总债项"`                      //总债项
	F013N       float64 `excel:"少数股东权益"`                   //少数股东权益
	F014N       float64 `excel:"股东权益/(资产净值)"`              //股东权益/(资产净值)
	F015N       float64 `excel:"股本"`                       //股本
	F016N       float64 `excel:"总储备"`                      //总储备
	F017N       float64 `excel:"营业额/（银行经营收入）"`             //营业额/（银行经营收入）
	F018N       float64 `excel:"经营溢利"`                     //经营溢利
	F019N       float64 `excel:"非经常性项目收益"`                 //非经常性项目收益
	F020N       float64 `excel:"融资成本/（财务费用）"`              //融资成本/（财务费用）
	F021N       float64 `excel:"应占联营公司溢利"`                 //应占联营公司溢利
	F022N       float64 `excel:"除税前经营溢利/(亏损)"`             //除税前经营溢利/(亏损)
	F023N       float64 `excel:"税项"`                       //税项
	F024N       float64 `excel:"除税后经营溢利/(亏损)"`             //除税后经营溢利/(亏损)
	F025N       float64 `excel:"少数股东损益"`                   //少数股东损益
	F026N       float64 `excel:"经调整股东应占溢利"`                //经调整股东应占溢利
	F027N       float64 `excel:"相对上期增减"`                   //相对上期增减
	F028N       float64 `excel:"股息总额"`                     //股息总额
	F029N       float64 `excel:"扣除非经常性项目后溢利"`              //扣除非经常性项目后溢利
	F030N       float64 `excel:"经营业务收到的现金/经营活动之现金流量（经调整）"` //经营业务收到的现金/经营活动之现金流量（经调整）
	F031N       float64 `excel:"投资活动之现金流量（经调整）"`           //投资活动之现金流量（经调整）
	F032N       float64 `excel:"融资活动之现金流量（经调整）"`           //融资活动之现金流量（经调整）
	F033N       float64 `excel:"现金流量净额"`                   //现金流量净额
	F034N       float64 `excel:"年初之现金及现金等价物"`              //年初之现金及现金等价物
	F035N       float64 `excel:"外币汇率变动之影响/其他"`             //外币汇率变动之影响/其他
	F036N       float64 `excel:"年终之现金及现金等价物"`              //年终之现金及现金等价物
	F037N       float64 `excel:"毛利率"`                      //毛利率
	F038N       float64 `excel:"税前利润率"`                    //税前利润率
	F039N       float64 `excel:"销售净利率"`                    //销售净利率
	F040N       float64 `excel:"产权比率（负债权益比）"`              //产权比率（负债权益比）
	F041N       float64 `excel:"资产报酬率"`                    //资产报酬率
	F042N       float64 `excel:"权益报酬率"`                    //权益报酬率
	F043N       float64 `excel:"流动比率"`                     //流动比率
	F044N       float64 `excel:"速动比率"`                     //速动比率
	F045N       float64 `excel:"现金与流动资产比率"`                //现金与流动资产比率
	F046N       float64 `excel:"资产负债率"`                    //资产负债率
	F047N       float64 `excel:"有形资产负债率"`                  //有形资产负债率
	F048N       float64 `excel:"长期负债比率"`                   //长期负债比率
	F049N       float64 `excel:"总资产周转率"`                   //总资产周转率
	F050N       float64 `excel:"流动资产周转率"`                  //流动资产周转率
	F051N       float64 `excel:"应收帐款周转率"`                  //应收帐款周转率
	F052N       float64 `excel:"存货周转率"`                    //存货周转率
	F053N       float64 `excel:"EBIT"`                     //EBIT
	F054N       float64 `excel:"EBITD"`                    //EBITD
	F055N       float64 `excel:"EBITDA"`                   //EBITDA
	F056N       float64 `excel:"息税折旧前利润率"`                 //息税折旧前利润率
	F057N       float64 `excel:"息税折旧摊销前利润率"`               //息税折旧摊销前利润率
	F058N       float64 `excel:"长期负债权益比率"`                 //长期负债权益比率
	F059N       float64 `excel:"营业费用比例"`                   //营业费用比例
	F060N       float64 `excel:"管理费用比例"`                   //管理费用比例
	F061N       float64 `excel:"财务费用比例"`                   //财务费用比例
	F062N       float64 `excel:"现金流动负债比"`                  //现金流动负债比
	F063N       float64 `excel:"非经常性损益率"`                  //非经常性损益率
	F064N       float64 `excel:"短期负债现金保障率"`                //短期负债现金保障率
	F065N       float64 `excel:"销售现金比率"`                   //销售现金比率
	F066N       float64 `excel:"全部资产现金回收率"`                //全部资产现金回收率
	F067N       float64 `excel:"经营现金净流量与净利润的比率"`           //经营现金净流量与净利润的比率
	F068N       float64 `excel:"业务利润"`                     //业务利润
	F069N       float64 `excel:"股东权益合计"`                   //股东权益合计
	F070N       float64 `excel:"未分配利润"`                    //未分配利润
	F071N       float64 `excel:"毛利润"`                      //毛利润
	F072V       string  `excel:"核数师意见"`                    //核数师意见
	F073N       float64 `excel:"平均股东权益回报率"`                //平均股东权益回报率
	F074N       float64 `excel:"平均资产回报率"`                  //平均资产回报率
	F075N       float64 `excel:"经调整每股现金流量"`                //经调整每股现金流量
	F076N       float64 `excel:"期末股价/每股资产净值"`              //期末股价/每股资产净值
	F077N       float64 `excel:"现金比率 "`                    //现金比率 (倍)
	F078V       string  `excel:"报表类型编码	char "`             //报表类型编码	char
	F079N       float64 `excel:"固定资产周转率"`                  //固定资产周转率
	F080N       float64 `excel:"流动资产周转天数"`                 //流动资产周转天数
	F081N       float64 `excel:"总资产周转天数"`                  //总资产周转天数
	F082N       float64 `excel:"股东权益周转率"`                  //股东权益周转率
	F083N       float64 `excel:"利息保障倍数"`                   //利息保障倍数
	F084N       float64 `excel:"应收帐款周转天数"`                 //应收帐款周转天数
	F085N       float64 `excel:"存货占比"`                     //存货占比
	F086N       float64 `excel:"股东权益增长率"`                  //股东权益增长率
	F087N       float64 `excel:"总负债增长率"`                   //总负债增长率
	F088N       float64 `excel:"经营溢利(营业利润)增长率"`            //经营溢利(营业利润)增长率
	F089N       float64 `excel:"固定资产增长率"`                  //固定资产增长率
	F090N       float64 `excel:"总资产增长率"`                   //总资产增长率
	F091N       float64 `excel:"毛利润增长率"`                   //毛利润增长率
	F092N       float64 `excel:"营业收入增长率"`                  //营业收入增长率
	F093N       float64 `excel:"净利润增长率"`                   //净利润增长率
	USD         float64 `excel:"报告货币兑美元"`                  //报告货币兑美元 报表截止日期当天汇率
	CNY         float64 `excel:"报告货币兑人民币"`                 //报告货币兑人民币 报表截止日期当天汇率
}
