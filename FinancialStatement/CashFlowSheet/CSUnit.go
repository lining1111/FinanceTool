package CashFlowSheet

import "github.com/golang/glog"

type CSUnit struct {
	Last CashFlowSheet
	Now  CashFlowSheet
}

func (csu *CSUnit) Check() bool {
	if csu.Now.Unit != csu.Last.Unit {
		glog.Error("CSUnit Check Unit fail,year:%v", csu.Now.CTime.String())
		return false
	}

	if csu.Now.Check() {
		glog.Error("CSUnit Check Now fail,year:%v", csu.Now.CTime.String())
		return false
	}

	if csu.Last.Check() {
		glog.Error("CSUnit Check Last fail,year:%v", csu.Last.CTime.String())
		return false
	}

	return true
}

//---------------数值比率分析------------------//
