package IncomeSheet

//IC 13.已获利息倍数=息税前利润/利息费用=(利润总额+财务费用)/(财务费用中的利息支出+资本化利息) 标准值 2.5
func (is *IncomeSheet) IC() float64 {
	return is.EBIT() / is.FinanceExpense
}

////RC Repayment Cover 偿债倍数=EBIT/((利息+本金偿还)/(1-所得税税率))
//func (is *IncomeSheet) RC() float64 {
//	//这里所得税按一般25%
//	return is.EBIT() / ((is.FinanceExpense+) / (1 - 0.25))
//}

//EBIT 系税前利润=利润总额+利息支出=净利润+所得税+利息支出
func (is *IncomeSheet) EBIT() float64 {
	return is.TotalProfit + is.FinanceExpense
}

//NPR 14.销售净利率=净利润/销售收入 标准值 0.1
func (is *IncomeSheet) NPR() float64 {
	return is.NetProfit / is.Revenue
}

//GIR 15.销售毛利率=(销售收入-销售成本)/销售成本 标准值 0.15
func (is *IncomeSheet) GIR() float64 {
	return (is.Revenue - is.SellingExpense) / is.SellingExpense
}
