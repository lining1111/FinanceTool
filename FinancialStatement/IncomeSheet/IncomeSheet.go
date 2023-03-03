package IncomeSheet

/**
利润表，数据类计算
*/

import (
	"github.com/golang/glog"
	"time"
)

//IncomeSheet 利润表
type IncomeSheet struct {
	Revenue                      float64          //营业收入
	CostOfSales                  float64          //减：营业成本
	TaxesAndSurcharge            float64          //减：税金及附加
	SellingExpense               float64          //减：销售费用
	GAExpense                    float64          //减：管理费用
	RDExpense                    float64          //减：研发费用 R&D expenses
	FinanceExpense               FinanceExpense   //减：财务费用 其中：利息费用+利息收入
	OtherIncome                  float64          //其他收益
	InvestmentIncome             InvestmentIncome //投资收益 损失以“-”号填列 其中：对联营企业和合营企业的投资收益 以摊余成本计量的金融资产终止确认收益（损失以“-”号填列）
	NEHI                         float64          //净敞口套期收益（损失以“-”号填列） Net exposure hedging income
	IncomeFromChangesInFairValue float64          //公允价值变动收益 损失以“-”号填列
	CreditImpairmentLoss         float64          //信用减值损失（损失以“-”号填列）
	AssetsImpairmentLoss         float64          //资产减值损失 损失以“-”号填列
	IFAD                         float64          //资产处置收益（损失以“-”号填列）Income from asset disposal
	OperationProfit              float64          //营业利润 亏损以“-”号填列
	NonOperatingIncome           float64          //加： 营业外收入
	NonOperatingExpense          float64          //减： 营业外支出
	TotalProfit                  float64          //利润总额
	IncomeTax                    float64          //企业所得税
	NetProfit                    float64          //净利润
	OCI                          float64          //其他综合收益各项目分别扣除所得税影响后的净额 Other Comprehensive Income
	Total                        float64
	EPS                          EarningsPerShare //每股收益：基本每股收益+稀释每股收益
	CTime                        time.Time        //表的时间
	Unit                         float64          //人民币金额单位
}

//FinanceExpense 财务费用 其中：利息费用+利息收入
type FinanceExpense struct {
	IE    float64 //利息费用 interest expenses
	II    float64 //利息收入 Interest income
	Total float64
}

func (fe *FinanceExpense) CalTotal() float64 {
	fe.Total = fe.IE - fe.II
	return fe.Total
}

//InvestmentIncome 投资收益 损失以“-”号填列 其中：对联营企业和合营企业的投资收益 以摊余成本计量的金融资产终止确认收益（损失以“-”号填列）
type InvestmentIncome struct {
	IIFAJV    float64 //对联营企业和合营企业的投资收益 Investment income from associates and joint ventures
	IFDOFAMAC float64 //以摊余成本计量的金融资产终止确认收益 Income from derecognition of financial assets measured at amortized cost
	Total     float64
}

func (ii *InvestmentIncome) CalTotal() float64 {
	ii.Total = ii.IIFAJV + ii.IFDOFAMAC
	return ii.Total
}

//EarningsPerShare 每股收益
type EarningsPerShare struct {
	Beps float64 //基本每股收益 Basic earnings per share
	Deps float64 //稀释每股收益 Diluted earnings per share
}

func (is *IncomeSheet) Check() bool {
	totalProfit := is.CalTotalProfit()
	if is.TotalProfit != totalProfit {
		glog.Error("IncomeSheet Check TotalProfit fail,year:%v", is.CTime.String())
		return false
	}

	netProfit := is.CalNetProfit()
	if is.NetProfit != netProfit {
		glog.Error("IncomeSheet Check NetProfit fail,year:%v", is.CTime.String())
		return false
	}

	total := is.CalTotal()
	return is.Total == total
}

//CalTotalProfit 计算利润总额
func (is *IncomeSheet) CalTotalProfit() float64 {
	is.TotalProfit = is.Revenue -
		is.CostOfSales - is.TaxesAndSurcharge - is.SellingExpense - is.GAExpense - is.FinanceExpense.CalTotal() +
		is.OtherIncome + is.InvestmentIncome.CalTotal() + is.NEHI + is.IncomeFromChangesInFairValue +
		is.CreditImpairmentLoss + is.AssetsImpairmentLoss +
		is.IFAD + is.OperationProfit + is.NonOperatingIncome - is.NonOperatingExpense
	return is.TotalProfit
}

//CalNetProfit 计算净利润
func (is *IncomeSheet) CalNetProfit() float64 {
	is.NetProfit = is.TotalProfit - is.IncomeTax
	return is.NetProfit
}

func (is *IncomeSheet) CalTotal() float64 {
	is.Total = is.NetProfit + is.OCI
	return is.Total
}

//---------------数值比率分析------------------//
