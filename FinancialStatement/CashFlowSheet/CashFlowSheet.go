package CashFlowSheet

/**
现金流量表，数据类计算
*/

import (
	"github.com/golang/glog"
	"time"
)

//CashFlowSheet 现金流量表
type CashFlowSheet struct {
	OCF   OperatingCashFlow
	ICF   InvestingCashFlow
	FCF   FinancingCashFlow
	Total float64
	CTime time.Time //表的时间
	Unit  float64   //人民币金额单位
}

func (cs *CashFlowSheet) Check() bool {
	if !cs.OCF.Check() {
		glog.Error("CashFlowSheet Check OCF fail,year:%v", cs.CTime.String())
		return false
	}

	if !cs.ICF.Check() {
		glog.Error("CashFlowSheet Check ICF fail,year:%v", cs.CTime.String())
		return false
	}

	if !cs.FCF.Check() {
		glog.Error("CashFlowSheet Check FCF fail,year:%v", cs.CTime.String())
		return false
	}

	if cs.Total != cs.CalTotal() {
		glog.Error("CashFlowSheet Check Total fail,year:%v", cs.CTime.String())
		return false
	}
	return true
}

func (cs *CashFlowSheet) CalTotal() float64 {
	total := cs.OCF.Total + cs.ICF.Total + cs.FCF.Total
	if cs.Total > 0 {
		if total != cs.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, cs.Total)
		}
	} else {
		cs.Total = total
	}
	return cs.Total
}

//OperatingCashFlow 经营活动现金流量
type OperatingCashFlow struct {
	In    OCashFlowIn
	Out   OCashFlowOut
	Total float64
}

func (ocf *OperatingCashFlow) Check() bool {
	if !ocf.In.Check() {
		glog.Error("OperatingCashFlow Check In fail")
		return false
	}
	if !ocf.Out.Check() {
		glog.Error("OperatingCashFlow Check Out fail")
		return false
	}

	total := ocf.CalTotal()
	return ocf.Total == total
}

func (ocf *OperatingCashFlow) CalTotal() float64 {
	total := ocf.In.Subtotal - ocf.Out.Subtotal
	if ocf.Total > 0 {
		if total != ocf.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, ocf.Total)
		}
	} else {
		ocf.Total = total
	}
	return ocf.Total
}

//OCashFlowIn 经营现金流入
type OCashFlowIn struct {
	RFCL     float64 //销售商品、提供劳务收到的现金 Cash from selling commodities or offering labor
	ROT      float64 //收到的税费返还 Return On Taxes
	Others   float64 //收到其他与经营活动有关的现金 Others
	Subtotal float64 //现金流入小计
}

func (ocfi *OCashFlowIn) Check() bool {
	total := ocfi.CalSubtotal()
	return ocfi.Subtotal == total
}

func (ocfi *OCashFlowIn) CalSubtotal() float64 {
	subtotal := ocfi.RFCL + ocfi.ROT + ocfi.Others
	if ocfi.Subtotal > 0 {
		if subtotal != ocfi.Subtotal {
			glog.Error("合计错误，计算值%v 获取值%v", subtotal, ocfi.Subtotal)
		}
	} else {
		ocfi.Subtotal = subtotal
	}

	return ocfi.Subtotal
}

//OCashFlowOut 经营现金流出
type OCashFlowOut struct {
	PFCL     float64 //购买商品、接受劳务支付的现金 Cash paid for commodities or labor
	PFE      float64 //支付给职工及为职工支付的现金 Cash paid to and on behalf of employees
	PFTax    float64 //支付的各项税费 PFTax
	Others   float64 //支付其他与经营活动有关的现金 Cash paid relating to other operating activities
	Subtotal float64 //现金流出小计
}

func (ocfo *OCashFlowOut) Check() bool {
	total := ocfo.CalSubtotal()
	return ocfo.Subtotal == total
}

func (ocfo *OCashFlowOut) CalSubtotal() float64 {
	subtotal := ocfo.PFCL + ocfo.PFE + ocfo.PFTax + ocfo.Others
	if ocfo.Subtotal > 0 {
		if subtotal != ocfo.Subtotal {
			glog.Error("合计错误，计算值%v 获取值%v", subtotal, ocfo.Subtotal)
		}
	} else {
		ocfo.Subtotal = subtotal
	}
	return ocfo.Subtotal
}

//InvestingCashFlow 投资活动现金流量
type InvestingCashFlow struct {
	In    ICashFlowIn
	Out   ICashFlowOut
	Total float64
}

func (icf *InvestingCashFlow) Check() bool {
	if !icf.In.Check() {
		glog.Error("InvestingCashFlow Check In fail")
		return false
	}

	if !icf.Out.Check() {
		glog.Error("InvestingCashFlow Check Out fail")
		return false
	}
	total := icf.CalTotal()
	return icf.Total == total
}

func (icf *InvestingCashFlow) CalTotal() float64 {
	total := icf.In.Subtotal - icf.Out.Subtotal
	if icf.Total > 0 {
		if total != icf.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, icf.Total)
		}
	} else {
		icf.Total = total
	}
	return icf.Total
}

//ICashFlowIn 投资现金流入
type ICashFlowIn struct {
	RFIW      float64 //收回投资收到的现金 Cash from investment withdrawal
	RFII      float64 //取得投资收益收到的现金 Cash from investment income
	RFFAIALTA float64 //处置固定资产、无形资产和其他长期资产收回的现金净额 Net cash from disposing fixed assets， intangible assets and other long-term assets
	RFSSOBU   float64 //处置子公司及其他营业单位收到的现金净额 Net cash received from disposal of subsidiaries and other business units
	Others    float64 //收到其他与投资活动有关的现金 Other cash received related to investing activities
	Subtotal  float64 //现金流入小计
}

func (icfi *ICashFlowIn) Check() bool {
	total := icfi.CalSubtotal()
	return icfi.Subtotal == total
}

func (icfi *ICashFlowIn) CalSubtotal() float64 {
	subtotal := icfi.RFIW + icfi.RFII + icfi.RFFAIALTA + icfi.RFSSOBU + icfi.Others
	if icfi.Subtotal > 0 {
		if subtotal != icfi.Subtotal {
			glog.Error("合计错误，计算值%v 获取值%v", subtotal, icfi.Subtotal)
		}
	} else {
		icfi.Subtotal = subtotal
	}
	return icfi.Subtotal
}

//ICashFlowOut 投资现金流出
type ICashFlowOut struct {
	PFFAIALTA float64 //购建固定资产、无形资产和其他长期资产支付的现金 Cash paid for buying fixed assets， intangible assets and other long-term investment
	PFI       float64 //投资支付的现金 Cash paid for investment
	PFSSOBU   float64 //取得子公司及其他营业单位支付的现金净额 net cash outflows of procurement of subsidiaries and other business units
	Others    float64 //支付其他与投资活动有关的现金 Other cash paid related to investing activities
	Subtotal  float64 //现金流出小计
}

func (icfo *ICashFlowOut) Check() bool {
	total := icfo.CalSubtotal()
	return icfo.Subtotal == total
}

func (icfo *ICashFlowOut) CalSubtotal() float64 {
	subtotal := icfo.PFFAIALTA + icfo.PFI + icfo.PFSSOBU + icfo.Others
	if icfo.Subtotal > 0 {
		if subtotal != icfo.Subtotal {
			glog.Error("合计错误，计算值%v 获取值%v", subtotal, icfo.Subtotal)
		}
	} else {
		icfo.Subtotal = subtotal
	}
	return icfo.Subtotal
}

//FinancingCashFlow 筹资活动现金流量
type FinancingCashFlow struct {
	In    FCashFlowIn
	Out   FCashFlowOut
	Total float64
}

func (fcf *FinancingCashFlow) Check() bool {
	if !fcf.In.Check() {
		glog.Error("FinancingCashFlow Check In fail")
		return false
	}

	if !fcf.Out.Check() {
		glog.Error("FinancingCashFlow Check Out fail")
		return false
	}

	total := fcf.CalTotal()
	return fcf.Total == total
}

func (fcf *FinancingCashFlow) CalTotal() float64 {
	total := fcf.In.Subtotal - fcf.Out.Subtotal
	if fcf.Total > 0 {
		if total != fcf.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, fcf.Total)
		}
	} else {
		fcf.Total = total
	}
	return fcf.Total
}

//FCashFlowIn 筹资现金流入
type FCashFlowIn struct {
	RFI        float64 //吸收投资收到的现金 Cash received from accepting investment
	Borrowings float64 //取得借款收到的现金 Borrowings
	Others     float64 //收到其他与筹资活动有关的现金 Other cash received related to financing activities
	Subtotal   float64 //现金流入小计
}

func (fcfi *FCashFlowIn) Check() bool {
	total := fcfi.CalSubtotal()
	return fcfi.Subtotal == total
}

func (fcfi *FCashFlowIn) CalSubtotal() float64 {
	subtotal := fcfi.RFI + fcfi.Borrowings + fcfi.Others
	if fcfi.Subtotal > 0 {
		if subtotal != fcfi.Subtotal {
			glog.Error("合计错误，计算值%v 获取值%v", subtotal, fcfi.Subtotal)
		}
	} else {
		fcfi.Subtotal = subtotal
	}
	return fcfi.Subtotal
}

//FCashFlowOut 筹资现金流出
type FCashFlowOut struct {
	PFDebt   float64 //偿还债务支付的现金 PFDebt
	PFDPI    float64 //分配股利、利润或偿付利息支付的现金 Cash paid for dividend ， profit or interest
	Others   float64 //支付其他与筹资活动有关的现金 Other cash paid related to financing activities
	Subtotal float64 //现金流出小计
}

func (fcfo *FCashFlowOut) Check() bool {
	total := fcfo.CalSubtotal()
	return fcfo.Subtotal == total
}

func (fcfo *FCashFlowOut) CalSubtotal() float64 {
	subtotal := fcfo.PFDebt + fcfo.PFDPI + fcfo.Others
	if fcfo.Subtotal > 0 {
		if subtotal != fcfo.Subtotal {
			glog.Error("合计错误，计算值%v 获取值%v", subtotal, fcfo.Subtotal)
		}
	} else {
		fcfo.Subtotal = subtotal
	}
	return fcfo.Subtotal
}
