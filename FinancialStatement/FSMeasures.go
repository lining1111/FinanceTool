package FinancialStatement

/**
url http://webapi.cninfo.com.cn/api/stock/p_stock2303?

*/
type FSMeasures struct {
	ORGNAME   string  //机构名称
	SECCODE   string  //证券代码
	SECNAME   string  //证券简称
	STARTDATE string  //开始日期
	ENDDATE   string  //截止日期
	F001V     string  //数据来源编码
	F002V     string  //数据来源
	F003N     float64 //每股收益
	F004N     float64 //基本每股收益
	F005N     float64 //稀释每股收益
	F006N     float64 //扣除非经常性损益每股收益
	F007N     float64 //每股未分配利润
	F008N     float64 //每股净资产 = 净资产/股本
	F009N     float64 //调整后每股净资产
	F010N     float64 //每股资本公积金 = 资本公积/股本
	F011N     float64 //营业利润率 = 营业利润/营业收入*100
	F012N     float64 //营业税金率 = 营业税金/营业收入*100
	F013N     float64 //营业成本率 = 营业成本/营业收入*100
	F014N     float64 //净资产收益率 = 净利润/股东权益*100
	F015N     float64 //投资收益率 = 投资收益/(交易性金融资产+可供出售金融资产+持有至到期投资+长期股权投资)*100
	F016N     float64 //总资产报酬率 = 净利润/(期末资产总额+期初资产总额)/2*100
	F017N     float64 //净利润率 = 净利润/营业收入*100
	F018N     float64 //管理费用率 = 管理费用/营业收入*100
	F019N     float64 //财务费用率 = 财务费用/营业收入*100
	F020N     float64 //成本费用利润率 = 利润总额/成本费用总额×100
	F021N     float64 //三费比重 = （财务费用+管理费用+营业费用）/营业收入*100
	F022N     float64 //应收账款周转率 = 营业收入/(期初应收帐款+期末应收帐款)/2
	F023N     float64 //存货周转率 = 营业成本/(期初存货+期末存货)/2
	F024N     float64 //运营资金周转率 = 营业收入/(期初流动资产-期初流动负债)+(期末流动资产+期末流动负债)/2
	F025N     float64 //总资产周转率 = 营业收入/((期初资产总计+期末资产总计)/2)
	F026N     float64 //固定资产周转率	= 营业收入/((期初固定资产+期末固定资产)/2)
	F027N     float64 //应收账款周转天数 = 360/应收账款周转率 或（期初应收账款+期末应收账款）/2] / 营业收入
	F028N     float64 //存货周转天数 = 360/存货周转率；存货周转率=[360*（期初存货+期末存货）/2]/ 营业成本
	F029N     float64 //流动资产周转率 = 营业收入/[（期初流动资产+期末流动资产）/2]
	F030N     float64 //流动资产周转天数 = 360/流动资产周转率
	F031N     float64 //总资产周转天数 = 360/总资产周转率
	F032N     float64 //股东权益周转率 = 营业收入*2/(期初股东权益+期末股东权益)
	F033N     float64 //流动资产比率 = 流动资产/总资产*100
	F034N     float64 //货币资金比率 = 货币资金/流动资产*100
	F035N     float64 //交易性金融资产比率 = 短期投资/流动资产*100
	F036N     float64 //存货比率 = 存货/流动资产*100
	F037N     float64 //固定资产比率 = 固定资产/总资产*100
	F038N     float64 //负债结构比 = 流动负债/非流动负债*100
	F039N     float64 //产权比率 = 负债总额/股东权益*100
	F040N     float64 //净资产比率 = 100-资产负债率；净资产/总资产*100
	F041N     float64 //资产负债比率 = 负债总额/资产总额*100
	F042N     float64 //流动比率 = 流动资产/流动负债
	F043N     float64 //速动比率 = (流动资产-存货)/流动负债
	F044N     float64 //现金比率 = 货币资金/流动负债*100
	F045N     float64 //利息保障倍数 = 息税前利润/利息费用=(利润总额+财务费用)/财务费用
	F046N     float64 //营运资金 = 流动资产-流动负债
	F047N     float64 //非流动负债比率 = 非流动负债/负债合计）*100
	F048N     float64 //流动负债比率 = 流动负债/总负债*100
	F049N     float64 //保守速动比率 = 货币资金+交易性金融资产+应收票据+应收帐款+其他应收款)/流动负债，酸性比率
	F050N     float64 //现金到期债务比率 = 经营活动现金净流量*100/本期到期的债务=经营活动现金净流量*100/（一年内到期的非流动负债＋应付票据）
	F051N     float64 //有形资产净值债务率 = 流动负债/(股东权益-无形资产)
	F052N     float64 //营业收入增长率 = (本期营业收入/上期营业收入-1)*100
	F053N     float64 //净利润增长率 = (本期净利润/上期净利润-1)*100
	F054N     float64 //净资产增长率 = (期末股东权益/期初股东权益-1)*100
	F055N     float64 //固定资产增长率 = (期末固定资产/期初固定资产-1)*100
	F056N     float64 //总资产增长率 = (期末资产总额/期初资产总额-1)*100
	F057N     float64 //投资收益增长率 = (本期投资收益/上期投资收益-1)*100
	F058N     float64 //营业利润增长率 = (本期营业利润/上年同期营业利润-1)*100
	F059N     float64 //每股现金流量 = 现金流量净额/股本
	F060N     float64 //每股经营现金流量 = 经营现金净流量/股本
	F061N     float64 //经营净现金比率（短期债务）= 经营现金净流量*100/流动负债
	F062N     float64 //经营净现金比率（全部债务）= 经营现金净流量*100/负债总额
	F063N     float64 //经营活动现金净流量与净利润比率 = 经营活动现金净流量*100/净利润
	F064N     float64 //营业收入现金含量 = 经营活动现金流入量*100/营业收入
	F065N     float64 //全部资产现金回收率 = 经营活动现金流量净额*100 /资产总额
	F066N     float64 //净资产收益率(扣除非经常性损益) 单位：%
	F067N     float64 //净资产收益率-加权
	F068N     float64 //净资产收益率-加权(扣除非经常性损益)
	F069D     string  //报告年度
	F070V     string  //合并类型编码
	F071V     string  //合并类型
	F076N     float64 //扣除非经常性损益后的净利润
	F077N     float64 //非经常性损益合计
	F078N     float64 //毛利率 = (营业务收入-营业务成本)/营业务收入*100%
	F079N     float64 //期间费用率=(年末期间费用/年末营业收入)*100%
	F080N     float64 //现金转换周期=存货转换期间+应收账款转换期间
	F081N     float64 //净资产收益率=年末净利润/年平均净资产
	F082N     float64 //净利含金量=经营现金净流量/净利润
	F083N     float64 //非经常性损益占比 =非经常性损益/净利润=(资产减值损失+公允价值变动净收益+投资收益+汇兑收益+补贴收入+营业外收入+营业外支出)/净收益
	F084N     float64 //期间费用增长率=((年末期间费用-上年末期间费用)/上年末期间费用)*100%
	F085N     float64 //基本获利能力=年化总资产息税前经常性收益率(EBIT) EBIT=净利润+所得税+利息
	F086N     float64 //应收账款占比=应收账款/总资产
	F087N     float64 //存货占比=存货/总资产
	F088N     float64 //年化期间费用毛利比=期间费用/毛利
	F089N     float64 //营业收入
	F090N     float64 //营业成本
	F091N     float64 //销售费用
	F092N     float64 //管理费用
	F093N     float64 //财务费用
	F094N     float64 //三费合计
	F095N     float64 //公允价值变动净收益
	F096N     float64 //投资收益
	F097N     float64 //营业利润
	F098N     float64 //补贴收入
	F099N     float64 //营业外收支净额
	F100N     float64 //利润总额
	F101N     float64 //净利润
	F102N     float64 //归属于母公司所有者的净利润
	F103N     float64 //扣除非经常性损益后的净利润(2007版)
	F104N     float64 //非经常性损益合计(2007版)
	F105N     float64 //经营活动现金流量净额
	F106N     float64 //投资活动现金流量净额
	F107N     float64 //筹资活动现金流量净额
	F108N     float64 //现金及现金等价物净增加额
	F109N     float64 //货币资金
	F110N     float64 //交易性金融资产
	F111N     float64 //应收账款
	F112N     float64 //存货
	F113N     float64 //流动资产合计
	F114N     float64 //投资性房地产
	F115N     float64 //商誉
	F116N     float64 //固定资产
	F117N     float64 //非流动资产合计
	F118N     float64 //资产总计
	F119N     float64 //流动负债合计
	F120N     float64 //非流动负债合计
	F121N     float64 //负债合计
	F122N     float64 //股本
	F123N     float64 //资本公积
	F124N     float64 //盈余公积
	F125N     float64 //库存股
	F126N     float64 //未分配利润
	F127N     float64 //少数股东权益
	F128N     float64 //股东权益合计
	F129N     float64 //归属于母公司所有者权益
}

//ITR 存货周转率=产品销售成本/((初期存货+期末存货)/2)
func (fs *FinanceStatement) ITR() float64 {
	return fs.ISU.Now.SellingExpense / ((fs.BSU.Last.A.CA.INV + fs.BSU.Now.A.CA.INV) / 2)
}

//ITD 存货周转天数=360/存货周转率=(360(初期存货+期末存货)/2)/产品销售成本
func (fs *FinanceStatement) ITD() float64 {
	return 360 / fs.ITR()
}

//ARTR 应收账款周转率=销售收入/(期初应收账款+期末应收账款)/2 标准值 3
func (fs *FinanceStatement) ARTR() float64 {
	return fs.ISU.Now.Revenue / ((fs.BSU.Last.A.CA.AR + fs.BSU.Now.A.CA.AR) / 2)
}

//ARTD 应收账款周转天数=360/应收账款周转率 标准值 100
func (fs *FinanceStatement) ARTD() float64 {
	return 360 / fs.ARTR()
}

//CAR 流动资产周转率=销售收入/((期初流动资产+期末流动资产)/2) 标准值 1
func (fs FinanceStatement) CAR() float64 {
	return fs.ISU.Now.Revenue / ((fs.BSU.Last.A.CA.Total + fs.BSU.Now.A.CA.Total) / 2)
}

//CAD 流动资产周转天数=360/(销售收入/((期初流动资产+期末流动资产)/2))
func (fs *FinanceStatement) CAD() float64 {
	return 360 / fs.CAR()
}

//NWCR 净营运资本周转率=销售收入/净营运资本
func (fs *FinanceStatement) NWCR() float64 {
	return fs.ISU.Now.Revenue / fs.BSU.Now.NWC()
}

//NWCD 净营运资本周转天数=360/(销售收入/净营运资本)
func (fs *FinanceStatement) NWCD() float64 {
	return 360 / fs.NWCR()
}

//NCAR 非流动资产周转率=销售收入/非流动资产
func (fs *FinanceStatement) NCAR() float64 {
	return fs.ISU.Now.Revenue / fs.BSU.Now.A.NCA.Total
}

//NCAD 非流动资产周转天数=360/(销售收入/非流动资产)
func (fs *FinanceStatement) NCAD() float64 {
	return 360 / fs.NCAR()
}

//TAR 总资产周转率=销售收入/((期初资产总额+期末资产总额)/2) 标准值 0.8
func (fs *FinanceStatement) TAR() float64 {
	return fs.ISU.Now.Revenue / ((fs.BSU.Last.A.Total + fs.BSU.Now.A.Total) / 2)
}

//TAD 总资产周转天数=360/(销售收入/((期初资产总额+期末资产总额)/2))
func (fs *FinanceStatement) TAD() float64 {
	return 360 / fs.TAR()
}

//OC 7.营业周期=存货周转天数+应收账款周转天数 标准值 200
func (fs *FinanceStatement) OC() float64 {
	return fs.ITD() + fs.ARTD()
}

//ROA 资产净利率=净利润/((期初资产总额+期末资产总额)/2) 无标准值
func (fs *FinanceStatement) ROA() float64 {
	return fs.ISU.Now.NetProfit / ((fs.BSU.Last.A.Total + fs.BSU.Now.A.Total) / 2)
}

//ROE 净资产收益率=净利润/((期初所有者权益合计+期末所有者权益合计)/2)
func (fs *FinanceStatement) ROE() float64 {
	return fs.ISU.Now.NetProfit / ((fs.BSU.Last.LEQ.Eq.Total + fs.BSU.Now.LEQ.Eq.Total) / 2)
}

//OCA 全部资产现金回收率=经营活动现金净流量/平均资产总额
func (fs *FinanceStatement) OCA() float64 {
	return fs.CSU.Now.OCF.Total / ((fs.BSU.Last.A.Total + fs.BSU.Now.A.Total) / 2)
}

//OCNP 盈利现金比率=经营现金净流量/净利润
func (fs *FinanceStatement) OCNP() float64 {
	return fs.CSU.Now.OCF.Total / fs.ISU.Now.NetProfit
}

//CFCL 现金流量比率=经营活动现金流量/流动负债
func (fs *FinanceStatement) CFCL() float64 {
	return fs.CSU.Now.OCF.Total / fs.BSU.Now.LEQ.Li.CL.Total
}

//OCICR 现金流量利息保障倍数=经营活动现金流量/利息费用
func (fs *FinanceStatement) OCICR() float64 {
	return fs.CSU.Now.OCF.Total / fs.ISU.Now.FinanceExpense.IE
}

//OCLI 经营现金流量债务比=经营活动现金流量/债务总额
func (fs *FinanceStatement) OCLI() float64 {
	return fs.CSU.Now.OCF.Total / fs.BSU.Now.LEQ.Li.Total
}
