package IncomeSheet

//GPM 毛利率=利润总和/营业收入
func (is *IncomeSheet) GPM() float64 {
	return is.TotalProfit / is.Revenue
}

//NPM 净利率=净利润/营业收入
func (is *IncomeSheet) NPM() float64 {
	return is.NetProfit / is.Revenue
}

//NPR 销售净利率=净利润/销售收入 标准值 0.1
func (is *IncomeSheet) NPR() float64 {
	return is.NetProfit / is.Revenue
}

//ICR 利息保障倍数=(净利润+利息费用+所得税费用)/利息费用
func (is *IncomeSheet) ICR() float64 {
	return (is.NetProfit + is.FinanceExpense.IE + is.IncomeTax) / is.FinanceExpense.IE
}

//EBIT 息税前利润
func (is *IncomeSheet) EBIT() float64 {
	return is.NetProfit + is.FinanceExpense.IE + is.IncomeTax
}
