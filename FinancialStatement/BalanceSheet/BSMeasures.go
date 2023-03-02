package BalanceSheet

import "github.com/golang/glog"

//CR 1.流动比率=流动资产合计/流动负债合计	标准值 2.0
func (bs *BalanceSheet) CR() float64 {
	return bs.A.CA.Total / bs.LEQ.Li.Cli.Total
}

//QR 2.速动比率=(流动资产合计-存货)/流动负债合计  保守值=(现金+短期投资+应收票据+应收账款)/流动负债   标准值1/0.8
func (bs *BalanceSheet) QR() (v float64, vConservative float64) {
	v = (bs.A.CA.Total - bs.A.CA.Inventories) / bs.LEQ.Li.Cli.Total

	qaConservative := bs.A.CA.Cash + bs.A.CA.FVTOCI + bs.A.CA.AccountReceivable

	glog.Info("保守流动资产值：%v", qaConservative)
	vConservative = qaConservative / bs.LEQ.Li.Cli.Total
	return
}

//DAR 10.资产负债率=负债总额/资产总额 标准值 0.7
func (bs *BalanceSheet) DAR() float64 {
	return bs.LEQ.Li.Total / bs.A.Total
}

//DER 11.产权比率=负债总额/股东权益 标准值 1.2
func (bs *BalanceSheet) DER() float64 {
	return bs.LEQ.Li.Total / bs.LEQ.Eq.Total
}

//DTER 12.有形净值债务率=负债总额/(股东权益-无形资产净值) 标准值 1.5
func (bs *BalanceSheet) DTER() float64 {
	return bs.LEQ.Li.Total / (bs.LEQ.Eq.Total - bs.A.NCA.IAOA)
}
