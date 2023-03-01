package IncomeSheet

import "github.com/golang/glog"

type ISUnit struct {
	Last IncomeSheet
	Now  IncomeSheet
}

func (isu *ISUnit) Check() bool {
	if isu.Now.Unit != isu.Last.Unit {
		glog.Error("ISUnit Check Unit fail,year:%v", isu.Now.CTime.String())
		return false
	}

	if isu.Now.Check() {
		glog.Error("ISUnit Check Now fail,year:%v", isu.Now.CTime.String())
		return false
	}

	if isu.Last.Check() {
		glog.Error("ISUnit Check Last fail,year:%v", isu.Last.CTime.String())
		return false
	}
	return true
}

//---------------数值比率分析------------------//
