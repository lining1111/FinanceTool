package BalanceSheet

import "github.com/golang/glog"

type BSUnit struct {
	Last BalanceSheet
	Now  BalanceSheet
}

func (bs *BSUnit) Check() bool {
	if bs.Now.Unit != bs.Last.Unit {
		glog.Error("BSUnit Check Unit fail,year:%v", bs.Now.CTime.String())
		return false
	}
	if !bs.Now.Check() {
		glog.Error("BSUnit Check Now fail,year:%v", bs.Now.CTime.String())
		return false
	}

	if !bs.Last.Check() {
		glog.Error("BSUnit Check Last fail,year:%v", bs.Last.CTime.String())
		return false
	}

	return true
}

//GetSurplusReservesInc 获取盈余公积的增量值
func (bs *BSUnit) GetSurplusReservesInc() float64 {
	return bs.Now.LEQ.Eq.SurplusReserves - bs.Last.LEQ.Eq.SurplusReserves
}

//---------------数值比率分析------------------//
