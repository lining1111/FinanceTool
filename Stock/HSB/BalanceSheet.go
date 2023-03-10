package HSB

import (
	"FinanceTool/COM/cninfo"
	"fmt"
)

const APIBS = "http://webapi.cninfo.com.cn/api/stock/p_stock2300"

//BalanceSheet 资产负债表 "http://webapi.cninfo.com.cn/api/stock/p_stock2300"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询报告期	string
//			edate	结束查询报告期	string
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type BalanceSheet struct {
	SECNAME     string //证券简称	varchar
	SECCODE     string //证券代码	varchar
	ORGNAME     string //机构名称	varchar
	DECLAREDATE string //公告日期	date
	ENDDATE     string //截止日期	date

	F001D string  //报告年度	date
	F002V string  //合并类型编码	varchar
	F003V string  //合并类型	varchar
	F004V string  //报表来源编码	varchar
	F005V string  //报表来源	varchar
	F006N float64 //货币资金	decimal		单位：元
	F077N float64 //结算备付金	decimal		单位：元
	F078N float64 //拆出资金	decimal		单位：元
	F007N float64 //以公允价值计量且其变动计入当期损益的金融资产(20190322弃用)	decimal		单位：元
	F080N float64 //衍生金融资产	decimal		单位：元
	F008N float64 //应收票据	decimal		单位：元
	F009N float64 //应收账款	decimal		单位：元
	F010N float64 //预付款项	decimal		单位：元
	F081N float64 //应收保费	decimal		单位：元
	F082N float64 //应收分保账款	decimal		单位：元
	F083N float64 //应收分保合同准备金	decimal		单位：元
	F013N float64 //其中：应收利息	decimal		单位：元
	F014N float64 //其中：应收股利	decimal		单位：元
	F011N float64 //其他应收款	decimal		单位：元
	F012N float64 //应收关联公司款	decimal		单位：元
	F084N float64 //买入返售金融资产	decimal		单位：元
	F015N float64 //存货	decimal		单位：元
	F016N float64 //其中：消耗性生物资产	decimal		单位：元
	F085N float64 //划分为持有待售的资产	decimal		单位：元
	F079N float64 //发放贷款及垫款-流动资产	decimal		单位：元
	F017N float64 //一年内到期的非流动资产	decimal		单位：元
	F117N float64 //交易性金融资产	DECIMAL(18,2)		单位：元
	F118N float64 //应收票据及应收账款	DECIMAL(18,2)		单位：元
	F119N float64 //合同资产	DECIMAL(18,2)		单位：元
	F018N float64 //其他流动资产	decimal		单位：元
	F019N float64 //流动资产合计	decimal		单位：元
	F086N float64 //发放贷款及垫款-非流动资产	decimal		单位：元
	F020N float64 //可供出售金融资产	decimal		单位：元
	F021N float64 //持有至到期投资	decimal		单位：元
	F022N float64 //长期应收款	decimal		单位：元
	F023N float64 //长期股权投资	decimal		单位：元
	F024N float64 //投资性房地产	decimal		单位：元
	F025N float64 //固定资产	decimal		单位：元
	F026N float64 //在建工程	decimal		单位：元
	F027N float64 //工程物资	decimal		单位：元
	F028N float64 //固定资产清理	decimal		单位：元
	F029N float64 //生产性生物资产	decimal		单位：元
	F030N float64 //油气资产	decimal		单位：元
	F031N float64 //无形资产	decimal		单位：元
	F032N float64 //开发支出	decimal		单位：元
	F033N float64 //商誉	decimal		单位：元
	F034N float64 //长期待摊费用	decimal		单位：元
	F035N float64 //递延所得税资产	decimal		单位：元
	F116N float64 //债权投资	DECIMAL(18,2)		单位：元
	F110N float64 //其他债权投资	DECIMAL(18,2)		单位：元
	F111N float64 //其他权益工具投资	DECIMAL(18,2)		单位：元
	F112N float64 //其他非流动金融资产	DECIMAL(18,2)		单位：元
	F036N float64 //其他非流动资产	decimal		单位：元
	F037N float64 //非流动资产合计	decimal		单位：元
	F038N float64 //资产总计	decimal		单位：元
	F039N float64 //短期借款	decimal		单位：元
	F087N float64 //向中央银行借款	decimal		单位：元
	F088N float64 //吸收存款及同业存放	decimal		单位：元
	F089N float64 //拆入资金	decimal		单位：元
	F040N float64 //以公允价值计量且其变动计入当期损益的金融负债（20190322弃用）	decimal		单位：元
	F090N float64 //衍生金融负债	decimal		单位：元
	F041N float64 //应付票据	decimal		单位：元
	F042N float64 //应付账款	decimal		单位：元
	F043N float64 //预收款项	decimal		单位：元
	F091N float64 //卖出回购金融资产款	decimal		单位：元
	F092N float64 //应付手续费及佣金	decimal		单位：元
	F044N float64 //应付职工薪酬	decimal		单位：元
	F045N float64 //应交税费	decimal		单位：元
	F046N float64 //其中：应付利息	decimal		单位：元
	F047N float64 //其中：应付股利	decimal		单位：元
	F048N float64 //其他应付款	decimal		单位：元
	F049N float64 //应付关联公司款	decimal		单位：元
	F093N float64 //应付分保账款	decimal		单位：元
	F094N float64 //保险合同准备金	decimal		单位：元
	F095N float64 //代理买卖证券款	decimal		单位：元
	F096N float64 //代理承销证券款	decimal		单位：元
	F097N float64 //划分为持有待售的负债	decimal		单位：元
	F050N float64 //一年内到期的非流动负债	decimal		单位：元
	F098N float64 //预计负债-流动负债	decimal		单位：元
	F099N float64 //递延收益-流动负债	decimal		单位：元
	F113N float64 //交易性金融负债	DECIMAL(18,2)		单位：元
	F114N float64 //应付票据及应付账款	DECIMAL(18,2)		单位：元
	F115N float64 //合同负债	DECIMAL(18,2)		单位：元
	F051N float64 //其他流动负债	decimal		单位：元
	F052N float64 //流动负债合计	decimal		单位：元
	F053N float64 //长期借款	decimal		单位：元
	F054N float64 //应付债券	decimal		单位：元
	F100N float64 //其中：优先股-非流动负债	decimal		单位：元
	F101N float64 //永续债-非流动负债	decimal		单位：元
	F055N float64 //长期应付款	decimal		单位：元
	F102N float64 //长期应付职工薪酬	decimal		单位：元
	F056N float64 //专项应付款	decimal		单位：元
	F057N float64 //预计负债	decimal		单位：元
	F075N float64 //递延收益-非流动负债	decimal		单位：元
	F058N float64 //递延所得税负债	decimal		单位：元
	F059N float64 //其他非流动负债	decimal		单位：元
	F060N float64 //非流动负债合计	decimal		单位：元
	F061N float64 //负债合计	decimal		单位：元
	F062N float64 //实收资本（或股本）	decimal		单位：元
	F103N float64 //其他权益工具	decimal		单位：元
	F104N float64 //其中：优先股-所有者权益	decimal		单位：元
	F105N float64 //永续债-所有者权益	decimal		单位：元
	F063N float64 //资本公积	decimal		单位：元
	F066N float64 //减：库存股	decimal		单位：元
	F074N float64 //其他综合收益	decimal		单位：元
	F072N float64 //专项储备	decimal		单位：元
	F064N float64 //盈余公积	decimal		单位：元
	F076N float64 //一般风险准备	decimal		单位：元
	F065N float64 //未分配利润	decimal		单位：元
	F068N float64 //外币报表折算价差	decimal		单位：元
	F073N float64 //归属于母公司所有者权益	decimal		单位：元
	F067N float64 //少数股东权益	decimal		单位：元
	F069N float64 //非正常经营项目收益调整	decimal		单位：元
	F070N float64 //所有者权益（或股东权益）合计	decimal		单位：元
	F071N float64 //负债和所有者（或股东权益）合计	decimal		单位：元
	MEMO  string  //备注	varchar
	F120N float64 //应收款项融资	decimal(18,2)		2019年8月新增
	F121N float64 //使用权资产	decimal(18,2)		2019年8月新增
	F122N float64 //租赁负债	decimal(18,2)		2019年8月新增
}

//BSGetFromCNINFByScode_test 合并本期 定期报告
func BSGetFromCNINFByScode_test() {
	bs := make([]BalanceSheet, 0, 20000)
	url := APIBS
	params := map[string]string{
		"scode": "000001",
		"sdate": cninfo.Getrdate("2021", cninfo.Q1),
		"edate": cninfo.Getrdate("2021", cninfo.Q2),
		"type":  cninfo.PublicCode071["合并本期"],
	}

	err := cninfo.GetInfoByScodeDate(url, params, &bs)
	if err != nil {
		fmt.Println(err)
	}
}

type BSU []BalanceSheet

//SetExcel 将资产负债表写入excel 返回 excel表格的文件名
func (u *BSU) SetExcel() (string, error) {

	return "", nil
}
