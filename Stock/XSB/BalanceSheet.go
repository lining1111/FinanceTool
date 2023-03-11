package XSB

const APIBS = "http://webapi.cninfo.com.cn/api/neeq/p_neeq6018"

//BalanceSheet 资产负债表 "http://webapi.cninfo.com.cn/api/neeq/p_neeq6018"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询报告期	string
//			edate	结束查询报告期	string
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type BalanceSheet struct {
	ORGNAME     string `excel:"机构名称"` //机构名称
	SECCODE     string `excel:"证券代码"` //证券代码
	SECNAME     string `excel:"证券简称"` //证券简称
	DECLAREDATE string `excel:"公告日期"` //公告日期
	ENDDATE     string `excel:"截止日期"` //截止日期

	F001D string  `excel:"报告年度"`                               //报告年度
	F002V string  `excel:"合并类型编码"`                             //合并类型编码 071001合并本期071002合并上期071003母公司本期071004母公司上期 	F003V string `excel:" 合并类型"` //合并类型
	F004V string  `excel:"报表来源编码"`                             //报表来源编码 033003定期报告033006转让说明书033013其他
	F005V string  `excel:"报表来源"`                               //报表来源
	F006N float64 `excel:"货币资金"`                               //货币资金
	F077N float64 `excel:"结算备付金"`                              //结算备付金
	F078N float64 `excel:"拆出资金"`                               //拆出资金
	F007N float64 `excel:"以公允价值计量且其变动计入当期损益的金融资产(20190322弃用)"` //以公允价值计量且其变动计入当期损益的金融资产(20190322弃用)
	F080N float64 `excel:"衍生金融资产"`                             //衍生金融资产
	F084N float64 `excel:"买入返售金融资产"`                           //买入返售金融资产
	F008N float64 `excel:"应收票据"`                               //应收票据
	F009N float64 `excel:"应收账款"`                               //应收账款
	F010N float64 `excel:"预付款项"`                               //预付款项
	F013N float64 `excel:"其中:应收利息"`                            //其中:应收利息
	F014N float64 `excel:"其中:应收股利"`                            //其中:应收股利
	F012N float64 `excel:"应收关联公司款"`                            //应收关联公司款
	F081N float64 `excel:"应收保费"`                               //应收保费
	F082N float64 `excel:"应收分保账款"`                             //应收分保账款
	F083N float64 `excel:"应收分保合同准备金"`                          //应收分保合同准备金
	F011N float64 `excel:"其他应收款"`                              //其他应收款
	F015N float64 `excel:"存货"`                                 //存货
	F016N float64 `excel:"其中:消耗性生物资产"`                         //其中:消耗性生物资产
	F085N float64 `excel:"划分为持有待售的资产"`                         //划分为持有待售的资产
	F017N float64 `excel:"一年内到期的非流动资产"`                        //一年内到期的非流动资产
	F079N float64 `excel:"发放贷款及垫款-流动资产"`                       //发放贷款及垫款-流动资产
	F018N float64 `excel:"其他流动资产"`                             //其他流动资产
	F019N float64 `excel:"流动资产合计"`                             //流动资产合计
	F020N float64 `excel:"可供出售金融资产"`                           //可供出售金融资产
	F021N float64 `excel:"持有至到期投资"`                            //持有至到期投资
	F022N float64 `excel:"长期应收款"`                              //长期应收款
	F023N float64 `excel:"长期股权投资"`                             //长期股权投资
	F024N float64 `excel:"投资性房地产"`                             //投资性房地产
	F025N float64 `excel:"固定资产"`                               //固定资产
	F026N float64 `excel:"在建工程"`                               //在建工程
	F027N float64 `excel:"工程物资"`                               //工程物资
	F028N float64 `excel:"固定资产清理"`                             //固定资产清理
	F029N float64 `excel:"生产性生物资产"`                            //生产性生物资产
	F030N float64 `excel:"油气资产"`                               //油气资产
	F031N float64 `excel:"无形资产"`                               //无形资产
	F032N float64 `excel:"开发支出"`                               //开发支出
	F033N float64 `excel:"商誉"`                                 //商誉
	F034N float64 `excel:"长期待摊费用"`                             //长期待摊费用
	F035N float64 `excel:"递延所得税资产"`                            //递延所得税资产
	F086N float64 `excel:"发放贷款及垫款-非流动资产"`                      //发放贷款及垫款-非流动资产
	F036N float64 `excel:"其他非流动资产"`                            //其他非流动资产
	F037N float64 `excel:"非流动资产合计"`                            //非流动资产合计
	F038N float64 `excel:"资产总计"`                               //资产总计
	F039N float64 `excel:"短期借款"`                               //短期借款
	F087N float64 `excel:"向中央银行借款"`                            //向中央银行借款
	F088N float64 `excel:"吸收存款及同业存放"`                          //吸收存款及同业存放
	F089N float64 `excel:"拆入资金"`                               //拆入资金
	F040N float64 `excel:"以公允价值计量且其变动计入当期损益的金融负债（20190322弃用)"` //以公允价值计量且其变动计入当期损益的金融负债（20190322弃用)
	F090N float64 `excel:"衍生金融负债"`                             //衍生金融负债
	F091N float64 `excel:"卖出回购金融资产款"`                          //卖出回购金融资产款
	F041N float64 `excel:"应付票据"`                               //应付票据
	F042N float64 `excel:"应付账款"`                               //应付账款
	F043N float64 `excel:"预收款项"`                               //预收款项
	F044N float64 `excel:"应付职工薪酬"`                             //应付职工薪酬
	F045N float64 `excel:"应交税费"`                               //应交税费
	F046N float64 `excel:"其中:应付利息"`                            //其中:应付利息
	F047N float64 `excel:"其中:应付股利"`                            //其中:应付股利
	F092N float64 `excel:"应付手续费及佣金"`                           //应付手续费及佣金
	F048N float64 `excel:"其他应付款"`                              //其他应付款
	F049N float64 `excel:"应付关联公司款"`                            //应付关联公司款
	F093N float64 `excel:"应付分保账款"`                             //应付分保账款
	F094N float64 `excel:"保险合同准备金"`                            //保险合同准备金
	F095N float64 `excel:"代理买卖证券款"`                            //代理买卖证券款
	F096N float64 `excel:"代理承销证券款"`                            //代理承销证券款
	F097N float64 `excel:"划分为持有待售的负债"`                         //划分为持有待售的负债
	F050N float64 `excel:"一年内到期的非流动负债"`                        //一年内到期的非流动负债
	F098N float64 `excel:"预计负债-流动负债"`                          //预计负债-流动负债
	F099N float64 `excel:"递延收益-流动负债"`                          //递延收益-流动负债
	F051N float64 `excel:"其他流动负债"`                             //其他流动负债
	F052N float64 `excel:"流动负债合计"`                             //流动负债合计
	F053N float64 `excel:"长期借款"`                               //长期借款
	F054N float64 `excel:"应付债券"`                               //应付债券
	F100N float64 `excel:"其中:优先股-非流动负债"`                       //其中:优先股-非流动负债
	F101N float64 `excel:"永续债-非流动负债"`                          //永续债-非流动负债
	F055N float64 `excel:"长期应付款"`                              //长期应付款
	F102N float64 `excel:"长期应付职工薪酬"`                           //长期应付职工薪酬
	F056N float64 `excel:"专项应付款"`                              //专项应付款
	F057N float64 `excel:"预计负债"`                               //预计负债
	F058N float64 `excel:"递延所得税负债"`                            //递延所得税负债
	F075N float64 `excel:"递延收益-非流动负债"`                         //递延收益-非流动负债
	F059N float64 `excel:"其他非流动负债"`                            //其他非流动负债
	F060N float64 `excel:"非流动负债合计"`                            //非流动负债合计
	F061N float64 `excel:"负债合计"`                               //负债合计
	F062N float64 `excel:"实收资本（或股本）"`                          //实收资本（或股本）
	F103N float64 `excel:"其他权益工具"`                             //其他权益工具
	F104N float64 `excel:"其中:优先股-所有者权益"`                       //其中:优先股-所有者权益
	F105N float64 `excel:"永续债-所有者权益"`                          //永续债-所有者权益
	F063N float64 `excel:"资本公积"`                               //资本公积
	F066N float64 `excel:"减:库存股"`                              //减:库存股
	F074N float64 `excel:"其他综合收益"`                             //其他综合收益
	F072N float64 `excel:"专项储备"`                               //专项储备
	F068N float64 `excel:"外币报表折算价差"`                           //外币报表折算价差
	F064N float64 `excel:"盈余公积"`                               //盈余公积
	F076N float64 `excel:"一般风险准备"`                             //一般风险准备
	F065N float64 `excel:"未分配利润"`                              //未分配利润
	F073N float64 `excel:"归属于母公司所有者权益"`                        //归属于母公司所有者权益
	F067N float64 `excel:"少数股东权益"`                             //少数股东权益
	F069N float64 `excel:"非正常经营项目收益调整"`                        //非正常经营项目收益调整
	F070N float64 `excel:"所有者权益（或股东权益）合计"`                     //所有者权益（或股东权益）合计
	F071N float64 `excel:"负债和所有者（或股东权益）合计"`                    //负债和所有者（或股东权益）合计
	MEMO  string  `excel:"备注"`                                 //备注
	F106N float64 `excel:"应收票据及应收账款"`                          //应收票据及应收账款 元	20210826新增
	F107N float64 `excel:"应付票据及应付账款"`                          //应付票据及应付账款 元	20210826新增
	F108N float64 `excel:"资产-其他资产"`                            //资产-其他资产 元	20210826新增
	F109N float64 `excel:"交易性金融资产"`                            //交易性金融资产_ 元	20210826新增
	F110N float64 `excel:"应收款项融资"`                             //应收款项融资 元	20210826新增
	F111N float64 `excel:"合同资产"`                               //合同资产 元	20210826新增
	F112N float64 `excel:"债权投资"`                               //债权投资 元	20210826新增
	F113N float64 `excel:"其他债权投资增"`                            //其他债权投资 元	20210826新增
	F114N float64 `excel:"其他权益工具投资"`                           //其他权益工具投资 元	20210826新增
	F115N float64 `excel:"其他非流动金融资产增"`                         //其他非流动金融资产 元	20210826新增
	F116N float64 `excel:"交易性金融负债"`                            //交易性金融负债_ 元	20210826新增
	F117N float64 `excel:"合同负债"`                               //合同负债 元	20210826新增
	F118N float64 `excel:"非流动负债-保险合同准备金"`                      //非流动负债-保险合同准备金 元	20210826新增
	F119N float64 `excel:"租赁负债"`                               //租赁负债 元	20210826新增
	F120N float64 `excel:"资产-其他负债"`                            //资产-其他负债 元	20210826新增
	F121N float64 `excel:"资产-发放贷款及垫款"`                         //资产-发放贷款及垫款 元	20210826新增
	F122N float64 `excel:"期货会员资格投资"`                           //期货会员资格投资 元	20210826新增
	F123N float64 `excel:"期货风险准备金"`                            //期货风险准备金 元	20210826新增
	F124N float64 `excel:"应收货币保证金"`                            //应收货币保证金 元	20210826新增
	F125N float64 `excel:"应付货币保证金"`                            //应付货币保证金 元	20210826新增
	F126N float64 `excel:"使用权资产"`                              //使用权资产 元	20210826新增
	F127N float64 `excel:"融出资金"`                               //融出资金 元	20210826新增
	F128N float64 `excel:"非流动资产-以公允价值计量且其变动计入当期损益的金融资产"`       //非流动资产-以公允价值计量且其变动计入当期损益的金融资产 元	20210826新增
	F129N float64 `excel:"非流动负债-以公允价值计量且其变动计入当期损益的金融负债"`       //非流动负债-以公允价值计量且其变动计入当期损益的金融负债 元	20210826新增
}
