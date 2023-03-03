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

//SGR 销售增长率=(期末营业收入-期初营业收入)/期初营业收入
func (csu *CSUnit) SGR() float64 {
	return (csu.Now.OCF.In.Subtotal - csu.Last.OCF.In.Subtotal) / csu.Last.OCF.In.Subtotal
}

//SPGR 营业利润增长率=(期末营业利润-期初营业利润)/期初营业利润
func (csu *CSUnit) SPGR() float64 {
	return (csu.Now.OCF.Total - csu.Last.OCF.Total) / csu.Last.OCF.Total
}
