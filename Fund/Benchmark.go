package Fund

const apiBm = "http://webapi.cninfo.com.cn/api/fund/p_fund2614"

//Benchmark 基金经营业绩及收益分配表2009版 "http://webapi.cninfo.com.cn/api/fund/p_fund2614"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询报告期	string
//			edate	结束查询报告期	string
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type Benchmark struct {
	SECCODE     string //基金代码	varchar
	SECNAME     string //基金简称	varchar
	DECLAREDATE string //公告日期	date
	STARTDATE   string //开始日期	date
	ENDDATE     string //截止日期	date

	F001V string  //报表来源编码	varchar		通过p_public0006可获取，对应的总类编码为033
	f002V string  //报表来源	varchar
	F004N float64 //利息收入	deciaml		单位：元
	F005N float64 //其中：存款利息收入	deciaml		单位：元
	F006N float64 //其中：债券利息收入	deciaml		单位：元
	F007N float64 //其中：资产支持证券利息收入	deciaml		单位：元
	F008N float64 //其中：买入返售金融资产收入	deciaml		单位：元
	F009N float64 //其中：其它利息收入	deciaml		单位：元
	F010N float64 //投资收益	deciaml		单位：元
	F011N float64 //其中：股票投资收益	deciaml		单位：元
	F012N float64 //其中：债券投资收益	deciaml		单位：元
	F013N float64 //其中：基金投资收益	deciaml		单位：元
	F014N float64 //其中：资产支持证券投资收益	deciaml		单位：元
	F015N float64 //其中：衍生工具收益	deciaml		单位：元
	F016N float64 //其中：股利收益	deciaml		单位：元
	F017N float64 //公允价值变动收益	deciaml		单位：元
	F018N float64 //买入返售证券收益	deciaml		单位：元
	F019N float64 //汇兑收益	deciaml		单位：元
	F020N float64 //其它收入	deciaml		单位：元
	F021N float64 //收入合计	deciaml		单位：元
	F023N float64 //管理人报酬	deciaml		单位：元
	F024N float64 //托管费	deciaml		单位：元
	F025N float64 //销售服务费	deciaml		单位：元
	F026N float64 //交易费用	deciaml		单位：元
	F027N float64 //利息支出	deciaml		单位：元
	F028N float64 //其中：卖出回购金融资产支出	deciaml		单位：元
	F029N float64 //其它费用	deciaml		单位：元
	F030N float64 //费用合计	deciaml		单位：元
	F031N float64 //三、利润总额	deciaml		单位：元
	F032N float64 //所得税费用	deciaml		单位：元
	F033N float64 //四、净利润	deciaml		单位：元
	F034N float64 //其中：贵金属投资收益	deciaml
}
