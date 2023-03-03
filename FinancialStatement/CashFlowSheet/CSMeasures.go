package CashFlowSheet

//OPM 营业利润率=营业利润/营业收入
func (cs *CashFlowSheet) OPM() float64 {
	return cs.OCF.Total / cs.OCF.In.Subtotal
}

//RFCLOC 销售收现比率=销售商品或提供劳务收到的现金/主营业务收入净额
func (cs *CashFlowSheet) RFCLOC() float64 {
	return cs.OCF.In.RFCL / cs.OCF.Total
}
