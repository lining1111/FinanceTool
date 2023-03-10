package HK

const APIFR = "http://webapi.cninfo.com.cn/api/hk/p_hk4025"

//FinancialRatios 重要指标表（非银行） "http://webapi.cninfo.com.cn/api/hk/p_hk4025"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type FinancialRatios struct {
	SECCODE     string  `excel:"证券代码"`                     //证券代码	varchar
	SECNAME     string  `excel:"证券简称"`                     //证券简称	varchar
	DECLAREDATE string  `excel:"公布日期"`                     //公布日期	date
	STARTDATE   string  `excel:"起始日期"`                     //起始日期	date
	ENDDATE     string  `excel:"截止日期"`                     //截止日期	date
	F001V       string  `excel:"报表类别"`                     //报表类别	varchar
	F002V       string  `excel:"币种编码"`                     //币种编码	varchar
	F003V       string  `excel:"币种"`                       //币种	varchar
	F004N       float64 `excel:"汇率(报表货币兑港元)"`              //汇率(报表货币兑港元)	decimal
	F005N       float64 `excel:"每股基本盈利（港元）"`               //每股基本盈利（港元）	decimal
	F006N       float64 `excel:"每股摊薄盈利（港元）"`               //每股摊薄盈利（港元）	decimal
	F007N       float64 `excel:"总资产"`                      //总资产	decimal
	F008N       float64 `excel:"非流动资产"`                    //非流动资产	decimal
	F009N       float64 `excel:"流动资产"`                     //流动资产	decimal
	F010N       float64 `excel:"流动负债"`                     //流动负债	decimal
	F011N       float64 `excel:"非流动负债"`                    //非流动负债	decimal
	F012N       float64 `excel:"总债项"`                      //总债项	decimal
	F013N       float64 `excel:"少数股东权益"`                   //少数股东权益	decimal
	F014N       float64 `excel:"股东权益/(资产净值)"`              //股东权益/(资产净值)	decimal
	F015N       float64 `excel:"股本"`                       //股本	decimal
	F016N       float64 `excel:"总储备"`                      //总储备	decimal
	F017N       float64 `excel:"营业额/（银行经营收入）"`             //营业额/（银行经营收入）	decimal
	F018N       float64 `excel:"经营溢利"`                     //经营溢利	decimal
	F019N       float64 `excel:"非经常性项目收益"`                 //非经常性项目收益	decimal
	F020N       float64 `excel:"融资成本/（财务费用）"`              //融资成本/（财务费用）	decimal
	F021N       float64 `excel:"应占联营公司溢利"`                 //应占联营公司溢利	decimal
	F022N       float64 `excel:"除税前经营溢利/(亏损)"`             //除税前经营溢利/(亏损)	decimal
	F023N       float64 `excel:"税项"`                       //税项	decimal
	F024N       float64 `excel:"除税后经营溢利/(亏损)"`             //除税后经营溢利/(亏损)	decimal
	F025N       float64 `excel:"少数股东损益"`                   //少数股东损益	decimal
	F026N       float64 `excel:"经调整股东应占溢利"`                //经调整股东应占溢利	decimal
	F027N       float64 `excel:"相对上期增减"`                   //相对上期增减	decimal
	F028N       float64 `excel:"股息总额"`                     //股息总额	decimal
	F029N       float64 `excel:"扣除非经常性项目后溢利"`              //扣除非经常性项目后溢利	decimal
	F030N       float64 `excel:"经营业务收到的现金/经营活动之现金流量（经调整）"` //经营业务收到的现金/经营活动之现金流量（经调整）	decimal
	F031N       float64 `excel:"投资活动之现金流量（经调整）"`           //投资活动之现金流量（经调整）	decimal
	F032N       float64 `excel:"融资活动之现金流量（经调整）"`           //融资活动之现金流量（经调整）	decimal
	F033N       float64 `excel:"现金流量净额"`                   //现金流量净额	decimal
	F034N       float64 `excel:"年初之现金及现金等价物"`              //年初之现金及现金等价物	decimal
	F035N       float64 `excel:"外币汇率变动之影响/其他"`             //外币汇率变动之影响/其他	decimal
	F036N       float64 `excel:"年终之现金及现金等价物"`              //年终之现金及现金等价物	decimal
	F037N       float64 `excel:"毛利率"`                      //毛利率	decimal
	F038N       float64 `excel:"税前利润率"`                    //税前利润率	decimal
	F039N       float64 `excel:"销售净利率"`                    //销售净利率	decimal
	F040N       float64 `excel:"产权比率（负债权益比）"`              //产权比率（负债权益比）	decimal
	F041N       float64 `excel:"资产报酬率"`                    //资产报酬率	decimal
	F042N       float64 `excel:"权益报酬率"`                    //权益报酬率	decimal
	F043N       float64 `excel:"流动比率"`                     //流动比率	decimal
	F044N       float64 `excel:"速动比率"`                     //速动比率	decimal
	F045N       float64 `excel:"现金与流动资产比率"`                //现金与流动资产比率	decimal
	F046N       float64 `excel:"资产负债率"`                    //资产负债率	decimal
	F047N       float64 `excel:"有形资产负债率"`                  //有形资产负债率	decimal
	F048N       float64 `excel:"长期负债比率"`                   //长期负债比率	decimal
	F049N       float64 `excel:"总资产周转率"`                   //总资产周转率	decimal
	F050N       float64 `excel:"流动资产周转率"`                  //流动资产周转率	decimal
	F051N       float64 `excel:"应收帐款周转率"`                  //应收帐款周转率	decimal
	F052N       float64 `excel:"存货周转率"`                    //存货周转率	decimal
	F053N       float64 `excel:"EBIT"`                     //EBIT	decimal
	F054N       float64 `excel:"EBITD"`                    //EBITD	decimal
	F055N       float64 `excel:"EBITDA"`                   //EBITDA	decimal
	F056N       float64 `excel:"息税折旧前利润率"`                 //息税折旧前利润率	decimal
	F057N       float64 `excel:"息税折旧摊销前利润率"`               //息税折旧摊销前利润率	decimal
	F058N       float64 `excel:"长期负债权益比率"`                 //长期负债权益比率	decimal
	F059N       float64 `excel:"营业费用比例"`                   //营业费用比例	decimal
	F060N       float64 `excel:"管理费用比例"`                   //管理费用比例	decimal
	F061N       float64 `excel:"财务费用比例"`                   //财务费用比例	decimal
	F062N       float64 `excel:"现金流动负债比"`                  //现金流动负债比	decimal
	F063N       float64 `excel:"非经常性损益率"`                  //非经常性损益率	decimal
	F064N       float64 `excel:"短期负债现金保障率"`                //短期负债现金保障率	decimal
	F065N       float64 `excel:"销售现金比率"`                   //销售现金比率	decimal
	F066N       float64 `excel:"全部资产现金回收率"`                //全部资产现金回收率	decimal
	F067N       float64 `excel:"经营现金净流量与净利润的比率"`           //经营现金净流量与净利润的比率	decimal
	F068N       float64 `excel:"业务利润"`                     //业务利润	decimal
	F069N       float64 `excel:"股东权益合计"`                   //股东权益合计	decimal
	F070N       float64 `excel:"未分配利润"`                    //未分配利润	decimal
	F071N       float64 `excel:"毛利润"`                      //毛利润	decimal
	F072V       string  `excel:"核数师意见"`                    //核数师意见	varchar
	F073N       float64 `excel:"平均股东权益回报率"`                //平均股东权益回报率	decimal
	F074N       float64 `excel:"平均资产回报率"`                  //平均资产回报率	decimal
	F075N       float64 `excel:"经调整每股现金流量"`                //经调整每股现金流量	decimal
	F076N       float64 `excel:"期末股价/每股资产净值"`              //期末股价/每股资产净值	decimal
	F077N       float64 `excel:"现金比率 "`                    //现金比率 (倍)
	F078V       string  `excel:"报表类型编码	char "`             //报表类型编码	char
	F079N       float64 `excel:"固定资产周转率"`                  //固定资产周转率	decimal
	F080N       float64 `excel:"流动资产周转天数"`                 //流动资产周转天数	decimal
	F081N       float64 `excel:"总资产周转天数"`                  //总资产周转天数	decimal
	F082N       float64 `excel:"股东权益周转率"`                  //股东权益周转率	decimal
	F083N       float64 `excel:"利息保障倍数"`                   //利息保障倍数	decimal
	F084N       float64 `excel:"应收帐款周转天数"`                 //应收帐款周转天数	decimal
	F085N       float64 `excel:"存货占比"`                     //存货占比	decimal
	F086N       float64 `excel:"股东权益增长率"`                  //股东权益增长率	decimal
	F087N       float64 `excel:"总负债增长率"`                   //总负债增长率	decimal
	F088N       float64 `excel:"经营溢利(营业利润)增长率"`            //经营溢利(营业利润)增长率	decimal
	F089N       float64 `excel:"固定资产增长率"`                  //固定资产增长率	decimal
	F090N       float64 `excel:"总资产增长率"`                   //总资产增长率	decimal
	F091N       float64 `excel:"毛利润增长率"`                   //毛利润增长率	decimal
	F092N       float64 `excel:"营业收入增长率"`                  //营业收入增长率	decimal
	F093N       float64 `excel:"净利润增长率"`                   //净利润增长率	decimal
	USD         float64 `excel:"报告货币兑美元"`                  //报告货币兑美元	double		报表截止日期当天汇率
	CNY         float64 `excel:"报告货币兑人民币"`                 //报告货币兑人民币	double		报表截止日期当天汇率
}
