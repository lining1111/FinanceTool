package BalanceSheet

import "github.com/golang/glog"

//DAR 资产负债率=负债总额/资产总额 标准值 0.7
func (bs *BalanceSheet) DAR() float64 {
	return bs.LEQ.Li.Total / bs.A.Total
}

//CR 流动比率=流动资产合计/流动负债合计	标准值 2.0
func (bs *BalanceSheet) CR() float64 {
	return bs.A.CA.Total / bs.LEQ.Li.CL.Total
}

//QR 2.速动比率=(流动资产合计-存货)/流动负债合计  保守值=(现金+短期投资+应收票据+应收账款)/流动负债   标准值1/0.8
func (bs *BalanceSheet) QR() (v float64, vConservative float64) {
	v = (bs.A.CA.Total - bs.A.CA.INV) / bs.LEQ.Li.CL.Total

	qaConservative := bs.A.CA.Cash + bs.A.CA.FVTOCI + bs.A.CA.AR

	glog.Info("保守流动资产值：%v", qaConservative)
	vConservative = qaConservative / bs.LEQ.Li.CL.Total
	return
}

//SQR Super Quick Ratio 超速动比率=(现金+短期债券+应收票据+应收账款净额)/流动负债
func (bs *BalanceSheet) SQR() float64 {
	quick := bs.A.CA.Cash + bs.A.CA.FVTOCI + bs.A.CA.AR
	return quick / bs.LEQ.Li.CL.Total
}

//DER 11.产权比率=负债总额/股东权益 标准值 1.2
func (bs *BalanceSheet) DER() float64 {
	return bs.LEQ.Li.Total / bs.LEQ.Eq.Total
}

//AER 权益乘数=资产总和/股东权益
func (bs *BalanceSheet) AER() float64 {
	return bs.A.Total / bs.LEQ.Eq.Total
}

//NWC 净运营资本=流动资产-流动负债
func (bs *BalanceSheet) NWC() float64 {
	return bs.A.CA.Total - bs.LEQ.Li.CL.Total
}

//CashRatio 现金比率=(现金+交易性金融资产)/流动负债
func (bs *BalanceSheet) CashRatio() float64 {
	return (bs.A.CA.Cash + bs.A.CA.FVTOCI) / bs.LEQ.Li.CL.Total
}

//LDWC Long Debt to Working Capital 长期负债与营运资金比率=长期负债/(流动资产-流动负债)
func (bs *BalanceSheet) LDWC() float64 {
	return bs.LEQ.Li.NCL.Total / (bs.A.CA.Total - bs.LEQ.Li.CL.Total)
}
