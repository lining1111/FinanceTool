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
	SECNAME     string `excel:"证券简称"` //证券简称	varchar
	SECCODE     string `excel:"证券代码"` //证券代码	varchar
	ORGNAME     string `excel:"机构名称"` //机构名称	varchar
	DECLAREDATE string `excel:"公告日期"` //公告日期	date
	ENDDATE     string `excel:"截止日期"` //截止日期	date

	F001D string  `excel:"报告年度"`                               //报告年度
	F002V string  `excel:"合并类型编码"`                             //合并类型编码
	F003V string  `excel:"合并类型"`                               //合并类型
	F004V string  `excel:"报表来源编码"`                             //报表来源编码
	F005V string  `excel:"报表来源"`                               //报表来源
	F006N float64 `excel:"货币资金"`                               //货币资金 单位:元
	F077N float64 `excel:"结算备付金"`                              //结算备付金 单位:元
	F078N float64 `excel:"拆出资金"`                               //拆出资金 单位:元
	F007N float64 `excel:"以公允价值计量且其变动计入当期损益的金融资产(20190322弃用)"` //以公允价值计量且其变动计入当期损益的金融资产(20190322弃用) 单位:元
	F080N float64 `excel:"衍生金融资产"`                             //衍生金融资产 单位:元
	F008N float64 `excel:"应收票据"`                               //应收票据 单位:元
	F009N float64 `excel:"应收账款"`                               //应收账款 单位:元
	F010N float64 `excel:"预付款项"`                               //预付款项 单位:元
	F081N float64 `excel:"应收保费"`                               //应收保费 单位:元
	F082N float64 `excel:"应收分保账款"`                             //应收分保账款 单位:元
	F083N float64 `excel:"应收分保合同准备金"`                          //应收分保合同准备金 单位:元
	F013N float64 `excel:"其中:应收利息"`                            //其中:应收利息 单位:元
	F014N float64 `excel:"其中:应收股利"`                            //其中:应收股利 单位:元
	F011N float64 `excel:"其他应收款"`                              //其他应收款 单位:元
	F012N float64 `excel:"应收关联公司款"`                            //应收关联公司款 单位:元
	F084N float64 `excel:"买入返售金融资产"`                           //买入返售金融资产 单位:元
	F015N float64 `excel:"存货"`                                 //存货 单位:元
	F016N float64 `excel:"其中:消耗性生物资产"`                         //其中:消耗性生物资产 单位:元
	F085N float64 `excel:"划分为持有待售的资产"`                         //划分为持有待售的资产 单位:元
	F079N float64 `excel:"发放贷款及垫款-流动资产"`                       //发放贷款及垫款-流动资产 单位:元
	F017N float64 `excel:"一年内到期的非流动资产"`                        //一年内到期的非流动资产 单位:元
	F117N float64 `excel:"交易性金融资产"`                            //交易性金融资产 单位:元
	F118N float64 `excel:"应收票据及应收账款"`                          //应收票据及应收账款 单位:元
	F119N float64 `excel:"合同资产"`                               //合同资产 单位:元
	F018N float64 `excel:"其他流动资产"`                             //其他流动资产 单位:元
	F019N float64 `excel:"流动资产合计"`                             //流动资产合计 单位:元
	F086N float64 `excel:"发放贷款及垫款-非流动资产"`                      //发放贷款及垫款-非流动资产 单位:元
	F020N float64 `excel:"可供出售金融资产"`                           //可供出售金融资产 单位:元
	F021N float64 `excel:"持有至到期投资"`                            //持有至到期投资 单位:元
	F022N float64 `excel:"长期应收款"`                              //长期应收款 单位:元
	F023N float64 `excel:"长期股权投资"`                             //长期股权投资 单位:元
	F024N float64 `excel:"投资性房地产"`                             //投资性房地产 单位:元
	F025N float64 `excel:"固定资产"`                               //固定资产 单位:元
	F026N float64 `excel:"在建工程"`                               //在建工程 单位:元
	F027N float64 `excel:"工程物资"`                               //工程物资 单位:元
	F028N float64 `excel:"固定资产清理"`                             //固定资产清理 单位:元
	F029N float64 `excel:"生产性生物资产"`                            //生产性生物资产 单位:元
	F030N float64 `excel:"油气资产"`                               //油气资产 单位:元
	F031N float64 `excel:"无形资产"`                               //无形资产 单位:元
	F032N float64 `excel:"开发支出"`                               //开发支出 单位:元
	F033N float64 `excel:"商誉"`                                 //商誉 单位:元
	F034N float64 `excel:"长期待摊费用"`                             //长期待摊费用 单位:元
	F035N float64 `excel:"递延所得税资产"`                            //递延所得税资产 单位:元
	F116N float64 `excel:"债权投资"`                               //债权投资 单位:元
	F110N float64 `excel:"其他债权投资"`                             //其他债权投资 单位:元
	F111N float64 `excel:"其他权益工具投资"`                           //其他权益工具投资 单位:元
	F112N float64 `excel:"其他非流动金融资产"`                          //其他非流动金融资产 单位:元
	F036N float64 `excel:"其他非流动资产"`                            //其他非流动资产 单位:元
	F037N float64 `excel:"非流动资产合计"`                            //非流动资产合计 单位:元
	F038N float64 `excel:"资产总计"`                               //资产总计 单位:元
	F039N float64 `excel:"短期借款"`                               //短期借款 单位:元
	F087N float64 `excel:"向中央银行借款"`                            //向中央银行借款 单位:元
	F088N float64 `excel:"吸收存款及同业存放"`                          //吸收存款及同业存放 单位:元
	F089N float64 `excel:"拆入资金"`                               //拆入资金 单位:元
	F040N float64 `excel:"以公允价值计量且其变动计入当期损益的金融负债(20190322弃用)"` //以公允价值计量且其变动计入当期损益的金融负债(20190322弃用) 单位:元
	F090N float64 `excel:"衍生金融负债"`                             //衍生金融负债 单位:元
	F041N float64 `excel:"应付票据"`                               //应付票据 单位:元
	F042N float64 `excel:"应付账款"`                               //应付账款 单位:元
	F043N float64 `excel:"预收款项"`                               //预收款项 单位:元
	F091N float64 `excel:"卖出回购金融资产款"`                          //卖出回购金融资产款 单位:元
	F092N float64 `excel:"应付手续费及佣金"`                           //应付手续费及佣金 单位:元
	F044N float64 `excel:"应付职工薪酬"`                             //应付职工薪酬 单位:元
	F045N float64 `excel:"应交税费"`                               //应交税费 单位:元
	F046N float64 `excel:"其中:应付利息"`                            //其中:应付利息 单位:元
	F047N float64 `excel:"其中:应付股利"`                            //其中:应付股利 单位:元
	F048N float64 `excel:"其他应付款"`                              //其他应付款 单位:元
	F049N float64 `excel:"应付关联公司款"`                            //应付关联公司款 单位:元
	F093N float64 `excel:"应付分保账款"`                             //应付分保账款 单位:元
	F094N float64 `excel:"保险合同准备金"`                            //保险合同准备金 单位:元
	F095N float64 `excel:"代理买卖证券款"`                            //代理买卖证券款 单位:元
	F096N float64 `excel:"代理承销证券款"`                            //代理承销证券款 单位:元
	F097N float64 `excel:"划分为持有待售的负债"`                         //划分为持有待售的负债 单位:元
	F050N float64 `excel:"一年内到期的非流动负债"`                        //一年内到期的非流动负债 单位:元
	F098N float64 `excel:"预计负债-流动负债"`                          //预计负债-流动负债 单位:元
	F099N float64 `excel:"递延收益-流动负债"`                          //递延收益-流动负债 单位:元
	F113N float64 `excel:"交易性金融负债"`                            //交易性金融负债 单位:元
	F114N float64 `excel:"应付票据及应付账款"`                          //应付票据及应付账款 单位:元
	F115N float64 `excel:"合同负债"`                               //合同负债 单位:元
	F051N float64 `excel:"其他流动负债"`                             //其他流动负债 单位:元
	F052N float64 `excel:"流动负债合计"`                             //流动负债合计 单位:元
	F053N float64 `excel:"长期借款"`                               //长期借款 单位:元
	F054N float64 `excel:"应付债券"`                               //应付债券 单位:元
	F100N float64 `excel:"其中:优先股-非流动负债"`                       //其中:优先股-非流动负债 单位:元
	F101N float64 `excel:"永续债-非流动负债"`                          //永续债-非流动负债 单位:元
	F055N float64 `excel:"长期应付款"`                              //长期应付款 单位:元
	F102N float64 `excel:"长期应付职工薪酬"`                           //长期应付职工薪酬 单位:元
	F056N float64 `excel:"专项应付款"`                              //专项应付款 单位:元
	F057N float64 `excel:"预计负债"`                               //预计负债 单位:元
	F075N float64 `excel:"递延收益-非流动负债"`                         //递延收益-非流动负债 单位:元
	F058N float64 `excel:"递延所得税负债"`                            //递延所得税负债 单位:元
	F059N float64 `excel:"其他非流动负债"`                            //其他非流动负债 单位:元
	F060N float64 `excel:"非流动负债合计"`                            //非流动负债合计 单位:元
	F061N float64 `excel:"负债合计"`                               //负债合计 单位:元
	F062N float64 `excel:"实收资本(或股本)"`                          //实收资本(或股本) 单位:元
	F103N float64 `excel:"其他权益工具"`                             //其他权益工具 单位:元
	F104N float64 `excel:"其中:优先股-所有者权益"`                       //其中:优先股-所有者权益 单位:元
	F105N float64 `excel:"永续债-所有者权益"`                          //永续债-所有者权益 单位:元
	F063N float64 `excel:"资本公积"`                               //资本公积 单位:元
	F066N float64 `excel:"减:库存股"`                              //减:库存股 单位:元
	F074N float64 `excel:"其他综合收益"`                             //其他综合收益 单位:元
	F072N float64 `excel:"专项储备"`                               //专项储备 单位:元
	F064N float64 `excel:"盈余公积"`                               //盈余公积 单位:元
	F076N float64 `excel:"一般风险准备"`                             //一般风险准备 单位:元
	F065N float64 `excel:"未分配利润"`                              //未分配利润 单位:元
	F068N float64 `excel:"外币报表折算价差"`                           //外币报表折算价差 单位:元
	F073N float64 `excel:"归属于母公司所有者权益"`                        //归属于母公司所有者权益 单位:元
	F067N float64 `excel:"少数股东权益"`                             //少数股东权益 单位:元
	F069N float64 `excel:"非正常经营项目收益调整"`                        //非正常经营项目收益调整 单位:元
	F070N float64 `excel:"所有者权益(或股东权益)合计"`                     //所有者权益(或股东权益)合计 单位:元
	F071N float64 `excel:"负债和所有者(或股东权益)合计"`                    //负债和所有者(或股东权益)合计 单位:元
	MEMO  string  `excel:"备注"`                                 //备注
	F120N float64 `excel:"应收款项融资"`                             //应收款项融资 2019年8月新增
	F121N float64 `excel:"使用权资产"`                              //使用权资产 2019年8月新增
	F122N float64 `excel:"租赁负债"`                               //租赁负债 2019年8月新增
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
