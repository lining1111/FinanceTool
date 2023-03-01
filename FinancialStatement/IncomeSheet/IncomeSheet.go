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
	Revenue                      float64 //营业收入
	CostOfSales                  float64 //营业成本
	TaxesAndSurcharge            float64 //税金及附加
	SellingExpense               float64 //销售费用
	GAExpense                    float64 //管理费用
	FinanceExpense               float64 //财务费用
	AssetsImpairmentLoss         float64 //资产减值损失 损失以“-”号填列
	IncomeFromChangesInFairValue float64 //公允价值变动收益 损失以“-”号填列
	InvestmentIncome             float64 //投资收益 损失以“-”号填列
	OperationProfit              float64 //营业利润 亏损以“-”号填列
	NonOperatingIncome           float64 //营业外收入
	NonOperatingExpense          float64 //营业外支出
	SubsidyIncome                float64 //补贴收入
	ExchangeGainOrLoss           float64 //汇兑损益 损失以“-”号填列
	TotalProfit                  float64 //利润总额
	IncomeTax                    float64 //企业所得税
	NetProfit                    float64 //净利润
	OCI                          float64 //Other Comprehensive Income其他综合收益各项目分别扣除所得税影响后的净额
	Total                        float64
	CTime                        time.Time //表的时间
	Unit                         float64   //人民币金额单位
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
		is.CostOfSales - is.TaxesAndSurcharge - is.SellingExpense - is.GAExpense - is.FinanceExpense +
		is.AssetsImpairmentLoss + is.IncomeFromChangesInFairValue + is.InvestmentIncome +
		is.OperationProfit + is.NonOperatingIncome - is.NonOperatingExpense + is.SubsidyIncome +
		is.ExchangeGainOrLoss
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
