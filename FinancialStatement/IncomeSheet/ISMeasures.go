package IncomeSheet

//IC 13.已获利息倍数=息税前利润/利息费用=(利润总额+财务费用)/(财务费用中的利息支出+资本化利息) 标准值 2.5
func (is *IncomeSheet) IC() float64 {
	return (is.TotalProfit + is.FinanceExpense) / is.FinanceExpense
}

//NPR 14.销售净利率=净利润/销售收入 标准值 0.1
func (is *IncomeSheet) NPR() float64 {
	return is.NetProfit / is.Revenue
}
