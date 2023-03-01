package BalanceSheet

import "github.com/golang/glog"

//CR 1.流动比率=流动资产合计/流动负债合计	标准值 2.0
func (bs *BalanceSheet) CR() float64 {
	return bs.A.CA.Total / bs.LEQ.Li.Cli.Total
}

//QR 2.速动比率=(流动资产合计-存货)/流动负债合计  保守值=(现金+短期投资+应收票据+应收账款)/流动负债   标准值1/0.8
func (bs BalanceSheet) QR() (v float64, vConservative float64) {
	v = (bs.A.CA.Total - bs.A.CA.Inventories) / bs.LEQ.Li.Cli.Total

	qaConservative := (bs.A.CA.Cash + bs.A.CA.FVTOCI + bs.A.CA.AccountReceivable)

	glog.Info("保守流动资产值：%v", qaConservative)
	vConservative = qaConservative / bs.LEQ.Li.Cli.Total
	return
}
