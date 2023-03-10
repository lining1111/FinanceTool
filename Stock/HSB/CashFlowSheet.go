package HSB

import (
	"FinanceTool/COM/cninfo"
	"fmt"
)

/**
现金流量表，数据类计算
*/

const APICS = "http://webapi.cninfo.com.cn/api/stock/p_stock2302"

//CashFlowSheet 现金流量表 "http://webapi.cninfo.com.cn/api/stock/p_stock2302"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询报告期	string	否
//			edate	结束查询报告期	string	否
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type CashFlowSheet struct {
	SECNAME     string //证券简称	varchar
	SECCODE     string //证券代码	varchar
	ORGNAME     string //机构名称	varchar
	DECLAREDATE string //公告日期	date
	STARTDATE   string //开始日期	date
	ENDDATE     string //截止日期	date

	F001D string  //报告年度	date
	F002V string  //合并类型编码	varchar
	F003V string  //合并类型	varchar
	F004V string  //报表来源编码	varchar
	F005V string  //报表来源	varchar
	F006N float64 //销售商品、提供劳务收到的现金	decimal		单位：元
	F072N float64 //客户存款和同业存放款项净增加额	decimal		单位：元
	F073N float64 //向中央银行借款净增加额	decimal		单位：元
	F074N float64 //向其他金融机构拆入资金净增加额	decimal		单位：元
	F077N float64 //收到原保险合同保费取得的现金	decimal		单位：元
	F078N float64 //收到再保险业务现金净额	decimal		单位：元
	F079N float64 //保户储金及投资款净增加额	decimal		单位：元
	F080N float64 //处置以公允价值计量且其变动计入当期损益的金融资产净增加额	decimal		单位：元
	F081N float64 //收取利息、手续费及佣金的现金	decimal		单位：元
	F082N float64 //拆入资金净增加额	decimal		单位：元
	F083N float64 //回购业务资金净增加额	decimal		单位：元
	F007N float64 //收到的税费返还	decimal		单位：元
	F008N float64 //收到其他与经营活动有关的现金	decimal		单位：元
	F009N float64 //经营活动现金流入小计	decimal		单位：元
	F010N float64 //购买商品、接受劳务支付的现金	decimal		单位：元
	F084N float64 //客户贷款及垫款净增加额	decimal		单位：元
	F085N float64 //存放中央银行和同业款项净增加额	decimal		单位：元
	F086N float64 //支付原保险合同赔付款项的现金	decimal		单位：元
	F087N float64 //支付利息、手续费及佣金的现金	decimal		单位：元
	F088N float64 //支付保单红利的现金	decimal		单位：元
	F011N float64 //支付给职工以及为职工支付的现金	decimal		单位：元
	F012N float64 //支付的各项税费	decimal		单位：元
	F013N float64 //支付其他与经营活动有关的现金	decimal		单位：元
	F014N float64 //经营活动现金流出小计	decimal		单位：元
	F015N float64 //经营活动产生的现金流量净额	decimal		单位：元
	F016N float64 //收回投资收到的现金	decimal		单位：元
	F017N float64 //取得投资收益收到的现金	decimal		单位：元
	F018N float64 //处置固定资产、无形资产和其他长期资产收回的现金净额	decimal		单位：元
	F019N float64 //处置子公司及其他营业单位收到的现金净额	decimal		单位：元
	F020N float64 //收到其他与投资活动有关的现金	decimal		单位：元
	F021N float64 //投资活动现金流入小计	decimal		单位：元
	F022N float64 //购建固定资产、无形资产和其他长期资产支付的现金	decimal		单位：元
	F023N float64 //投资支付的现金	decimal		单位：元
	F075N float64 //质押贷款净增加额	decimal		单位：元
	F024N float64 //取得子公司及其他营业单位支付的现金净额	decimal		单位：元
	F025N float64 //支付其他与投资活动有关的现金	decimal		单位：元
	F026N float64 //投资活动现金流出小计	decimal		单位：元
	F027N float64 //投资活动产生的现金流量净额	decimal		单位：元
	F028N float64 //吸收投资收到的现金	decimal		单位：元
	F089N float64 //其中：子公司吸收少数股东投资收到的现金	decimal		单位：元
	F029N float64 //取得借款收到的现金	decimal		单位：元
	F076N float64 //发行债券收到的现金	decimal		单位：元
	F030N float64 //收到其他与筹资活动有关的现金	decimal		单位：元
	F031N float64 //筹资活动现金流入小计	decimal		单位：元
	F032N float64 //偿还债务支付的现金	decimal		单位：元
	F033N float64 //分配股利、利润或偿付利息支付的现金	decimal		单位：元
	F090N float64 //其中：子公司支付给少数股东的股利、利润	decimal		单位：元
	F034N float64 //支付其他与筹资活动有关的现金	decimal		单位：元
	F035N float64 //筹资活动现金流出小计	decimal		单位：元
	F036N float64 //筹资活动产生的现金流量净额	decimal		单位：元
	F037N float64 //四、汇率变动对现金的影响	decimal		单位：元
	F038N float64 //四(2)、其他原因对现金的影响	decimal		单位：元
	F039N float64 //五、现金及现金等价物净增加额	decimal		单位：元
	F040N float64 //期初现金及现金等价物余额	decimal		单位：元
	F041N float64 //期末现金及现金等价物余额	decimal		单位：元
	F044N float64 //净利润	decimal		单位：元
	F045N float64 //加：资产减值准备	decimal		单位：元
	F046N float64 //固定资产折旧、油气资产折耗、生产性生物资产折旧	decimal		单位：元
	F091N float64 //投资性房地产的折旧及摊销	decimal		单位：元
	F047N float64 //无形资产摊销	decimal		单位：元
	F048N float64 //长期待摊费用摊销	decimal		单位：元
	F049N float64 //处置固定资产、无形资产和其他长期资产的损失	decimal		单位：元
	F050N float64 //固定资产报废损失	decimal		单位：元
	F051N float64 //公允价值变动损失	decimal		单位：元
	F052N float64 //财务费用	decimal		单位：元
	F053N float64 //投资损失	decimal		单位：元
	F054N float64 //递延所得税资产减少	decimal		单位：元
	F055N float64 //递延所得税负债增加	decimal		单位：元
	F056N float64 //存货的减少	decimal		单位：元
	F057N float64 //经营性应收项目的减少	decimal		单位：元
	F058N float64 //经营性应付项目的增加	decimal		单位：元
	F059N float64 //其他	decimal		单位：元
	F060N float64 //经营活动产生的现金流量净额2	decimal		单位：元
	F062N float64 //债务转为资本	decimal		单位：元
	F063N float64 //一年内到期的可转换公司债券	decimal		单位：元
	F064N float64 //融资租入固定资产	decimal		单位：元
	F066N float64 //现金的期末余额	decimal		单位：元
	F067N float64 //减：现金的期初余额	decimal		单位：元
	F068N float64 //加：现金等价物的期末余额	decimal		单位：元
	F069N float64 //减：现金等价物的期初余额	decimal		单位：元
	F070N float64 //加：其他原因对现金的影响2	decimal		单位：元
	F071N float64 //现金及现金等价物净增加额2	decimal		单位：元
	F096N float64 //信用减值损失	decimal
}

func CSGetFromCNINFByScode_test() {
	cs := make([]CashFlowSheet, 0, 20000)
	url := APICS
	params := map[string]string{
		"scode": "000001",
		"sdate": cninfo.Getrdate("2021", cninfo.Q1),
		"edate": cninfo.Getrdate("2021", cninfo.Q1),
		"type":  "071001",
	}

	err := cninfo.GetInfoByScodeDate(url, params, &cs)
	if err != nil {
		fmt.Println(err)
	}
}
