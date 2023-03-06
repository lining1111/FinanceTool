package FinancialStatement

import (
	"time"
)

type FSMeasures struct {
	ComCd     string //公司代码 如果公司发行了A股，则公司代码=C+A股代码，如果仅发现B股，则公司代码=C+B股代码
	StkCd     string //股票代码
	StkCdOTrd string //交易时股票代码
	A_StkCd   string //A股代码
	B_StkCd   string //B股代码
	LComNm    string //最新公司全称
	LstFlg    string //上市标识 按A。B。H股的顺序，依次列出该公司上市的股票类型。例如：如果公司只发B股，则记为B；如果公司既发A股又发B股则记为AB；如果公司既发B股又发H股则记为BH；
	//报表信息
	EndDT         time.Time //截至日期
	ReportDT      time.Time //报表日期
	InfoSourceflg string    //信息来源标识 Q1一季度 Q2中报 Q3三季度 Q4年报 0-其他
	InfoSource    string    //信息来源 Q1一季度；中报；三季度；年报；股份报价转让说明书；上市公司报告书；增发A股招股意向书；招股说明书；招股意向书；增发新股上市公告书
	ReportType    string    //报表类型 Q1一季度，Q2中报 Q3三季度 Q4年报 0-其他
	//报表内容
	//盈利能力分析
	NPR float64 //销售净利率
	ROA float64 //资产净利率
	ROE float64 //权益净利率
	//float64//总资产报酬率
	OPM float64 //营业利润率
	// float64//成本费用利润率
	//盈利质量分析
	OCA float64 //全部资产现金回收率
	//float64//盈利现金比率
	RFCLOC float64 //销售收现比率
	//偿债能力分析
	NWC       float64 //净运营资本
	CR        float64 //流动比率
	QR        float64 //速动比率
	SQR       float64 //超速动比率
	CashRatio float64 //现金比率
	CFCL      float64 //现金流量比率
	DAR       float64 //资产负债率
	DER       float64 //产权比率
	AER       float64 //权益乘数
	ICR       float64 //利息保障倍数
	OCICR     float64 //现金流量利息保障倍数
	OCLI      float64 //经营现金流量债务比
	//营运能力分析
	ARTR float64 //应收账款周转率
	ARTD float64 //应收账款周转天数
	ITR  float64 //存货周转率
	ITD  float64 //存货周转天数
	CAR  float64 //流动资产周转率
	CAD  float64 //流动资产周转天数
	NWCR float64 //净营运资本周转率
	NWCD float64 //净营运资本周转天数
	NCAR float64 //非流动资产周转率
	NCAD float64 //非流动资产周转天数
	TAR  float64 //总资产周转率
	TAD  float64 //总资产周转天数
	//发展能力分析
	EGR  float64 //股东权益增长率
	AGR  float64 //资产增长率
	SGR  float64 //销售增长率
	SPGR float64 //营业利润增长率
	NPGR float64 //净利润增长率
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
