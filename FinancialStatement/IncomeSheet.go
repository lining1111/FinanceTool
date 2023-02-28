package FinancialStatement

//IncomeSheet 利润表
type IncomeSheet struct {
	Revenue                      float64 //营业收入
	CostOfSales                  float64 //营业成本
	TaxesAndSurcharge            float64 //税金及附加
	SellingExpense               float64 //销售费用
	GAndAExpense                 float64 //管理费用
	FinanceExpense               float64 //财务费用
	AssetsImpairmentLoss         float64 //资产减值损失
	IncomeFromChangesInFairValue float64 //公允价值变动收益
	InvestmentIncome             float64 //投资收益
	OperationProfit              float64 //营业利润
	NonOperatingIncome           float64 //营业外收入
	NonOperatingExpense          float64 //营业外支出
	SubsidyIncome                float64 //补贴收入
	ExchangeGainOrLoss           float64 //汇兑损益
	TotalProfit                  float64 //利润总额
	IncomeTax                    float64 //企业所得税
	NetProfit                    float64 //净利润
	OCI                          float64 //Other Comprehensive Income其他综合收益各项目分别扣除所得税影响后的净额
	Total                        float64
}
