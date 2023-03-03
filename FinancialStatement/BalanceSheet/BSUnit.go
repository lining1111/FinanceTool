package BalanceSheet

import "github.com/golang/glog"

type BSUnit struct {
	Last BalanceSheet
	Now  BalanceSheet
}

func (bsu *BSUnit) Check() bool {
	if bsu.Now.Unit != bsu.Last.Unit {
		glog.Error("BSUnit Check Unit fail,year:%v", bsu.Now.CTime.String())
		return false
	}
	if !bsu.Now.Check() {
		glog.Error("BSUnit Check Now fail,year:%v", bsu.Now.CTime.String())
		return false
	}

	if !bsu.Last.Check() {
		glog.Error("BSUnit Check Last fail,year:%v", bsu.Last.CTime.String())
		return false
	}

	return true
}

//GetSurplusReservesInc 获取盈余公积的增量值
func (bsu *BSUnit) GetSurplusReservesInc() float64 {
	return bsu.Now.LEQ.Eq.SurplusReserves - bsu.Last.LEQ.Eq.SurplusReserves
}

//---------------数值比率分析------------------//

//AGR 资产增长率=(期末资产总和-期初资产总和)/期初资产总和
func (bsu *BSUnit) AGR() float64 {
	return (bsu.Now.A.Total - bsu.Last.A.Total) / bsu.Last.A.Total
}

//EGR 股东权益增长率=(期末股东权益总和-期初股东权益总和)/期初股东权益总和
func (bsu *BSUnit) EGR() float64 {
	return (bsu.Now.LEQ.Eq.Total - bsu.Last.LEQ.Eq.Total) / bsu.Last.LEQ.Eq.Total
}

//WCAR Working Capital Assets Ratio 营运资金/资产总额=(流动资产-流动负债)/平均资产总额
func (bsu *BSUnit) WCAR() float64 {
	return bsu.Now.NWC() / ((bsu.Now.A.Total + bsu.Last.A.Total) / 2)
}
