package HSB

import (
	"FinanceTool/COM/cninfo"
	"fmt"
)

const APIFR = "http://webapi.cninfo.com.cn/api/stock/p_stock2303"

//FinancialRatios 财务指标 "http://webapi.cninfo.com.cn/api/stock/p_stock2303"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询报告期	string	否
//			edate	结束查询报告期	string	否
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type FinancialRatios struct {
	ORGNAME   string //机构名称	VARCHAR
	SECCODE   string //证券代码	DECIMAL
	SECNAME   string //证券简称	VARCHAR
	STARTDATE string //开始日期	DATE
	ENDDATE   string //截止日期	DATE

	F001V string  //数据来源编码	VARCHAR
	F002V string  //数据来源	VARCHAR
	F003N float64 //每股收益	DECIMAL(14,6)		单位：元
	F004N float64 //基本每股收益	DECIMAL(14,6)		单位：元
	F005N float64 //稀释每股收益	DECIMAL(14,6)		单位：元
	F006N float64 //扣除非经常性损益每股收益	DECIMAL(14,6)		单位：元
	F007N float64 //每股未分配利润	DECIMAL(12,4)		单位：元
	F008N float64 //每股净资产	DECIMAL(12,4)		净资产/股本
	F009N float64 //调整后每股净资产	DECIMAL(12,4)
	F010N float64 //每股资本公积金	DECIMAL(12,4)		资本公积/股本
	F011N float64 //营业利润率	DECIMAL(10,3)		营业利润/营业收入*100
	F012N float64 //营业税金率	DECIMAL(10,3)		营业税金/营业收入*100
	F013N float64 //营业成本率	DECIMAL(10,3)		营业成本/营业收入*100
	F014N float64 //净资产收益率	DECIMAL(10,3)		净利润/股东权益*100
	F015N float64 //投资收益率	DECIMAL(10,3)		投资收益/(交易性金融资产+可供出售金融资产+持有至到期投资+长期股权投资)*100
	F016N float64 //总资产报酬率	DECIMAL(10,3)		净利润/(期末资产总额+期初资产总额)/2*100
	F017N float64 //净利润率	DECIMAL(15,3)		净利润/营业收入*100
	F018N float64 //管理费用率	DECIMAL(15,3)		管理费用/营业收入*100
	F019N float64 //财务费用率	DECIMAL(15,3)		财务费用/营业收入*100
	F020N float64 //成本费用利润率	DECIMAL(15,3)		利润总额/成本费用总额×100
	F021N float64 //三费比重	DECIMAL(15,3)		（财务费用+管理费用+营业费用）/营业收入*100
	F022N float64 //应收账款周转率	DECIMAL(10,3)		营业收入/(期初应收帐款+期末应收帐款)/2
	F023N float64 //存货周转率	DECIMAL(10,3)		营业成本/(期初存货+期末存货)/2
	F024N float64 //运营资金周转率	DECIMAL(10,3)		营业收入/(期初流动资产-期初流动负债)+(期末流动资产+期末流动负债)/2
	F025N float64 //总资产周转率	DECIMAL(10,3)		营业收入/((期初资产总计+期末资产总计)/2)
	F026N float64 //固定资产周转率	DECIMAL(10,3)		营业收入/((期初固定资产+期末固定资产)/2)
	F027N float64 //应收账款周转天数	DECIMAL(10,3)		360/应收账款周转率 或（期初应收账款+期末应收账款）/2] / 营业收入
	F028N float64 //存货周转天数	DECIMAL(10,3)		360/存货周转率；存货周转率=[360*（期初存货+期末存货）/2]/ 营业成本
	F029N float64 //流动资产周转率	DECIMAL(10,3)		营业收入/[（期初流动资产+期末流动资产）/2]
	F030N float64 //流动资产周转天数	DECIMAL(10,3)		360/流动资产周转率
	F031N float64 //总资产周转天数	DECIMAL(10,3)		360/总资产周转率
	F032N float64 //股东权益周转率	DECIMAL(10,3)		营业收入*2/(期初股东权益+期末股东权益)
	F033N float64 //流动资产比率	DECIMAL(10,3)		流动资产/总资产*100
	F034N float64 //货币资金比率	DECIMAL(10,3)		货币资金/流动资产*100
	F035N float64 //交易性金融资产比率	DECIMAL(10,3)		短期投资/流动资产*100
	F036N float64 //存货比率	DECIMAL(10,3)		存货/流动资产*100
	F037N float64 //固定资产比率	DECIMAL(10,3)		固定资产/总资产*100
	F038N float64 //负债结构比	DECIMAL(15,3)		流动负债/非流动负债*100
	F039N float64 //产权比率	DECIMAL(10,3)		负债总额/股东权益*100
	F040N float64 //净资产比率	DECIMAL(10,3)		100-资产负债率；净资产/总资产*100
	F041N float64 //资产负债比率	DECIMAL(10,3)		负债总额/资产总额*100
	F042N float64 //流动比率	DECIMAL(10,3)		流动资产/流动负债
	F043N float64 //速动比率	DECIMAL(10,3)		(流动资产-存货)/流动负债
	F044N float64 //现金比率	DECIMAL(10,3)		货币资金/流动负债*100
	F045N float64 //利息保障倍数	DECIMAL(10,3)		息税前利润/利息费用=(利润总额+财务费用)/财务费用
	F046N float64 //营运资金	DECIMAL(18,2)		流动资产-流动负债
	F047N float64 //非流动负债比率	DECIMAL(10,3)		（非流动负债/负债合计）*100
	F048N float64 //流动负债比率	DECIMAL(10,3)		流动负债/总负债*100
	F049N float64 //保守速动比率	DECIMAL(10,3)		货币资金+交易性金融资产+应收票据+应收帐款+其他应收款)/流动负债，酸性比率
	F050N float64 //现金到期债务比率	DECIMAL(15,3)		经营活动现金净流量*100/本期到期的债务=经营活动现金净流量*100/（一年内到期的非流动负债＋应付票据）
	F051N float64 //有形资产净值债务率	DECIMAL(10,3)		流动负债/(股东权益-无形资产)
	F052N float64 //营业收入增长率	DECIMAL(10,3)		(本期营业收入/上期营业收入-1)*100
	F053N float64 //净利润增长率	DECIMAL(15,3)		(本期净利润/上期净利润-1)*100
	F054N float64 //净资产增长率	DECIMAL(10,3)		(期末股东权益/期初股东权益-1)*100
	F055N float64 //固定资产增长率	DECIMAL(10,3)		(期末固定资产/期初固定资产-1)*100
	F056N float64 //总资产增长率	DECIMAL(10,3)		(期末资产总额/期初资产总额-1)*100
	F057N float64 //投资收益增长率	DECIMAL(15,3)		(本期投资收益/上期投资收益-1)*100
	F058N float64 //营业利润增长率	DECIMAL(10,3)		(本期营业利润/上年同期营业利润-1)*100
	F059N float64 //每股现金流量	DECIMAL(10,3)		现金流量净额/股本
	F060N float64 //每股经营现金流量	DECIMAL(10,3)		经营现金净流量/股本
	F061N float64 //经营净现金比率（短期债务）	DECIMAL(10,3)		经营现金净流量*100/流动负债
	F062N float64 //经营净现金比率（全部债务）	DECIMAL(10,3)		经营现金净流量*100/负债总额
	F063N float64 //经营活动现金净流量与净利润比率	DECIMAL(10,3)		经营活动现金净流量*100/净利润
	F064N float64 //营业收入现金含量	DECIMAL(10,3)		经营活动现金流入量*100/营业收入
	F065N float64 //全部资产现金回收率	DECIMAL(10,3)		经营活动现金流量净额*100 /资产总额
	F066N float64 //净资产收益率(扣除非经常性损益)	DECIMAL(12,4)		单位：%
	F067N float64 //净资产收益率-加权	DECIMAL(12,4)		单位：%
	F068N float64 //净资产收益率-加权(扣除非经常性损益)	DECIMAL(12,4)		单位：%
	F069D string  //报告年度	DATE
	F070V string  //合并类型编码	VARCHAR
	F071V string  //合并类型	VARCHAR
	F076N float64 //扣除非经常性损益后的净利润	DECIMAL(18,2)		单位：元
	F077N float64 //非经常性损益合计	DECIMAL(18,2)		单位：元
	F078N float64 //毛利率	DECIMAL(18,2)		毛利率=(营业务收入-营业务成本)/营业务收入*100%
	F079N float64 //期间费用率	DECIMAL(18,2)		期间费用率=(年末期间费用/年末营业收入)*100%
	F080N float64 //现金转换周期	DECIMAL(18,2)		现金转换周期=存货转换期间+应收账款转换期间
	F081N float64 //净资产收益率	DECIMAL(18,2)		净资产收益率=年末净利润/年平均净资产
	F082N float64 //净利含金量	DECIMAL(18,2)		净利含金量=经营现金净流量/净利润
	F083N float64 //非经常性损益占比	DECIMAL(18,2)		非经常性损益占比 =非经常性损益/净利润=(资产减值损失+公允价值变动净收益+投资收益+汇兑收益+补贴收入+营业外收入+营业外支出)/净收益
	F084N float64 //期间费用增长率	DECIMAL(18,2)		期间费用增长率=((年末期间费用-上年末期间费用)/上年末期间费用)*100%
	F085N float64 //基本获利能力	DECIMAL(18,2)		基本获利能力=年化总资产息税前经常性收益率(EBIT) EBIT=净利润+所得税+利息
	F086N float64 //应收账款占比	DECIMAL(18,2)		应收账款占比=应收账款/总资产
	F087N float64 //存货占比	DECIMAL(18,2)		存货占比=存货/总资产
	F088N float64 //年化期间费用毛利比	DECIMAL(18,2)		年化期间费用毛利比=期间费用/毛利
	F089N float64 //营业收入	numeric(18,2)
	F090N float64 //营业成本	numeric(18,2)
	F091N float64 //销售费用	numeric(18,2)
	F092N float64 //管理费用	numeric(18,2)
	F093N float64 //财务费用	numeric(18,2)
	F094N float64 //三费合计	numeric(18,2)
	F095N float64 //公允价值变动净收益	numeric(18,2)
	F096N float64 //投资收益	numeric(18,2)
	F097N float64 //营业利润	numeric(18,2)
	F098N float64 //补贴收入	numeric(18,2)
	F099N float64 //营业外收支净额	numeric(18,2)
	F100N float64 //利润总额	numeric(18,2)
	F101N float64 //净利润	numeric(18,2)
	F102N float64 //归属于母公司所有者的净利润	numeric(18,2)
	F103N float64 //扣除非经常性损益后的净利润(2007版)	numeric(18,2)
	F104N float64 //非经常性损益合计(2007版)	numeric(18,2)
	F105N float64 //经营活动现金流量净额	numeric(18,2)
	F106N float64 //投资活动现金流量净额	numeric(18,2)
	F107N float64 //筹资活动现金流量净额	numeric(18,2)
	F108N float64 //现金及现金等价物净增加额	numeric(18,2)
	F109N float64 //货币资金	numeric(18,2)
	F110N float64 //交易性金融资产	numeric(18,2)
	F111N float64 //应收账款	numeric(18,2)
	F112N float64 //存货	numeric(18,2)
	F113N float64 //流动资产合计	numeric(18,2)
	F114N float64 //投资性房地产	numeric(18,2)
	F115N float64 //商誉	numeric(18,2)
	F116N float64 //固定资产	numeric(18,2)
	F117N float64 //非流动资产合计	numeric(18,2)
	F118N float64 //资产总计	numeric(18,2)
	F119N float64 //流动负债合计	numeric(18,2)
	F120N float64 //非流动负债合计	numeric(18,2)
	F121N float64 //负债合计	numeric(18,2)
	F122N float64 //股本	numeric(18,2)
	F123N float64 //资本公积	numeric(18,2)
	F124N float64 //盈余公积	numeric(18,2)
	F125N float64 //库存股	numeric(18,2)
	F126N float64 //未分配利润	numeric(18,2)
	F127N float64 //少数股东权益	numeric(18,2)
	F128N float64 //股东权益合计	numeric(18,2)
	F129N float64 //归属于母公司所有者权益	numeric(18,2)
}

func FRGetFromCNINFByScode_test() {
	fr := make([]FinancialRatios, 1, 20000)
	url := APIFR
	params := map[string]string{
		"scode": "000001",
		"sdate": cninfo.Getrdate("2021", cninfo.Q1),
		"edate": cninfo.Getrdate("2021", cninfo.Q1),
		"type":  "071001",
	}

	err := cninfo.GetInfoByScodeDate(url, params, &fr, cap(fr))
	if err != nil {
		fmt.Println(err)
	}
}
