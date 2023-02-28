package IncomeSheet

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
}

//CalTotalProfit 计算利润总额
func (incs IncomeSheet) CalTotalProfit() float64 {
	incs.TotalProfit = incs.Revenue -
		incs.CostOfSales - incs.TaxesAndSurcharge - incs.SellingExpense - incs.GAExpense - incs.FinanceExpense +
		incs.AssetsImpairmentLoss + incs.IncomeFromChangesInFairValue + incs.InvestmentIncome +
		incs.OperationProfit + incs.NonOperatingIncome - incs.NonOperatingExpense + incs.SubsidyIncome +
		incs.ExchangeGainOrLoss
	return incs.TotalProfit
}

//CalNetProfit 计算净利润
func (incs IncomeSheet) CalNetProfit() float64 {
	incs.NetProfit = incs.TotalProfit - incs.IncomeTax
	return incs.NetProfit
}

func (incs *IncomeSheet) CalTotal() float64 {
	incs.Total = incs.NetProfit + incs.OCI
	return incs.Total
}
