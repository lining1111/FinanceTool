package Fund

const apiBS = "http://webapi.cninfo.com.cn/api/fund/p_fund2615"

//BalanceSheet 基金资产负债表2009版 "http://webapi.cninfo.com.cn/api/fund/p_fund2615"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询报告期	string
//			edate	结束查询报告期	string
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type BalanceSheet struct {
	SECCODE     string `excel:"基金代码"` //基金代码	varchar
	SECNAME     string `excel:"基金简称"` //基金简称	varchar
	DECLAREDATE string `excel:"公告日期"` //公告日期	date
	STARTDATE   string `excel:"开始日期"` //开始日期	date
	ENDDATE     string `excel:"截止日期"` //截止日期	date

	F001V string  `excel:"报表来源编码"`      //报表来源编码	varchar		通过p_public0006可获取，对应的总类编码为033
	F002V string  `excel:"报表来源"`        //报表来源	varchar
	F004N float64 `excel:"银行存款"`        //银行存款	decimal		单位：元
	F005N float64 `excel:"结算备付金"`       //结算备付金	decimal		单位：元
	F006N float64 `excel:"存出保证金"`       //存出保证金	decimal		单位：元
	F007N float64 `excel:"交易性金融资产"`     //交易性金融资产	decimal		单位：元
	F008N float64 `excel:"其中：股票投资"`     //其中：股票投资	decimal		单位：元
	F009N float64 `excel:"其中：基金投资"`     //其中：基金投资	decimal		单位：元
	F010N float64 `excel:"其中：债券投资"`     //其中：债券投资	decimal		单位：元
	F011N float64 `excel:"其中：权证投资"`     //其中：权证投资	decimal		单位：元
	F012N float64 `excel:"其中：资产支持证券投资"` //其中：资产支持证券投资	decimal		单位：元
	F013N float64 `excel:"其中：其他投资"`     //其中：其他投资	decimal		单位：元
	F014N float64 `excel:"衍生金融资产"`      //衍生金融资产	decimal		单位：元
	F015N float64 `excel:"买入返售金融资产"`    //买入返售金融资产	decimal		单位：元
	F016N float64 `excel:"应收证券清算款"`     //应收证券清算款	decimal		单位：元
	F017N float64 `excel:"应收利息"`        //应收利息	decimal		单位：元
	F018N float64 `excel:"应收股利"`        //应收股利	decimal		单位：元
	F019N float64 `excel:"应收申购款"`       //应收申购款	decimal		单位：元
	F020N float64 `excel:"其它应收款"`       //其它应收款	decimal		单位：元
	F021N float64 `excel:"递延所得税资产"`     //递延所得税资产	decimal		单位：元
	F022N float64 `excel:"其他资产"`        //其他资产	decimal		单位：元
	F023N float64 `excel:"资产总计"`        //资产总计	decimal		单位：元
	F026N float64 `excel:"短期借款"`        //短期借款	decimal		单位：元
	F027N float64 `excel:"交易性金融负债"`     //交易性金融负债	decimal		单位：元
	F028N float64 `excel:"衍生金融负债"`      //衍生金融负债	decimal		单位：元
	F029N float64 `excel:"卖出回购金融资产款"`   //卖出回购金融资产款	decimal		单位：元
	F030N float64 `excel:"应付证券清算款"`     //应付证券清算款	decimal		单位：元
	F031N float64 `excel:"应付赎回款"`       //应付赎回款	decimal		单位：元
	F032N float64 `excel:"应付管理人报酬"`     //应付管理人报酬	decimal		单位：元
	F033N float64 `excel:"应付托管费"`       //应付托管费	decimal		单位：元
	F034N float64 `excel:"应付销售服务费"`     //应付销售服务费	decimal		单位：元
	F035N float64 `excel:"应付交易费用"`      //应付交易费用	decimal		单位：元
	F036N float64 `excel:"应交税费"`        //应交税费	decimal		单位：元
	F037N float64 `excel:"应付利息"`        //应付利息	decimal		单位：元
	F038N float64 `excel:"应付利润"`        //应付利润	decimal		单位：元
	F039N float64 `excel:"递延所得税负债"`     //递延所得税负债	decimal		单位：元
	F040N float64 `excel:"其他负债"`        //其他负债	decimal		单位：元
	F041N float64 `excel:"负债合计"`        //负债合计	decimal		单位：元
	F043N float64 `excel:"实收基金"`        //实收基金	decimal		单位：元
	F044N float64 `excel:"未分配利润"`       //未分配利润	decimal		单位：元
	F045N float64 `excel:"所有者权益合计"`     //所有者权益合计	decimal		单位：元
	F046N float64 `excel:"负债和所有者权益合计"`  //负债和所有者权益合计	decimal		单位：元
	F047N float64 `excel:"基金份额总额"`      //基金份额总额	decimal		单位：份
	F048N float64 `excel:"基金份额净值"`      //基金份额净值	decimal		单位：元
	F049N float64 `excel:"其中：贵金属投资"`    //其中：贵金属投资	decimal
}
