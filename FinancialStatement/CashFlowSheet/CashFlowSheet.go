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
	cs.Total = cs.OCF.Total + cs.ICF.Total + cs.FCF.Total
	return cs.Total
}

//OperatingCashFlow 经营活动现金流量
type OperatingCashFlow struct {
	In    OCashInflow
	Out   OCashOutflow
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
	ocf.Total = ocf.In.Subtotal - ocf.Out.Subtotal
	return ocf.Total
}

//OCashInflow 经营现金流入
type OCashInflow struct {
	CFSCOL   float64 //Cash from selling commodities or offering labor销售商品、提供劳务收到的现金
	ROT      float64 //Return On Taxes收到的税费返还
	OCPRTOA  float64 //Other cash paid related to operating activities收到其他与经营活动有关的现金
	Subtotal float64 //现金流入小计
}

func (ocif *OCashInflow) Check() bool {
	total := ocif.CalSubtotal()
	return ocif.Subtotal == total
}

func (ocif *OCashInflow) CalSubtotal() float64 {
	ocif.Subtotal = ocif.CFSCOL + ocif.ROT + ocif.OCPRTOA
	return ocif.Subtotal
}

//OCashOutflow 经营现金流出
type OCashOutflow struct {
	CPFCOL      float64 //Cash paid for commodities or labor购买商品、接受劳务支付的现金
	CPAOBOE     float64 //Cash paid to and on behalf of employees支付给职工及为职工支付的现金
	TaxPayments float64 //支付的各项税费
	CPRTOOA     float64 //Cash paid relating to other operating activities支付其他与经营活动有关的现金
	Subtotal    float64 //现金流出小计
}

func (ocof *OCashOutflow) Check() bool {
	total := ocof.CalSubtotal()
	return ocof.Subtotal == total
}

func (ocof *OCashOutflow) CalSubtotal() float64 {
	ocof.Subtotal = ocof.CPFCOL + ocof.CPAOBOE + ocof.TaxPayments + ocof.CPRTOOA
	return ocof.Subtotal
}

//InvestingCashFlow 投资活动现金流量
type InvestingCashFlow struct {
	In    ICashInflow
	Out   ICashOutflow
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
	icf.Total = icf.In.Subtotal - icf.Out.Subtotal
	return icf.Total
}

//ICashInflow 投资现金流入
type ICashInflow struct {
	CFIW         float64 //Cash from investment withdrawal收回投资收到的现金
	CFII         float64 //Cash from investment income取得投资收益收到的现金
	NCFDFAIAOLTA float64 //Net cash from disposing fixed assets， intangible assets and other long-term assets 处置固定资产、无形资产和其他长期资产收回的现金净额
	NCRFDOSAOBU  float64 //Net cash received from disposal of subsidiaries and other business units处置子公司及其他营业单位收到的现金净额
	OCRRTIA      float64 //Other cash received related to investing activities收到其他与投资活动有关的现金
	Subtotal     float64 //现金流入小计
}

func (icif *ICashInflow) Check() bool {
	total := icif.CalSubtotal()
	return icif.Subtotal == total
}

func (icif *ICashInflow) CalSubtotal() float64 {
	icif.Subtotal = icif.CFIW + icif.CFII + icif.NCFDFAIAOLTA + icif.NCRFDOSAOBU + icif.OCRRTIA
	return icif.Subtotal
}

//ICashOutflow 投资现金流出
type ICashOutflow struct {
	CAFBFAIAOLTA float64 //Cash paid for buying fixed assets， intangible assets and other long-term investment购建固定资产、无形资产和其他长期资产支付的现金
	CPFI         float64 //Cash paid for investment投资支付的现金
	NCOFOPSAOBU  float64 //net cash outflows of procurement of subsidiaries and other business units取得子公司及其他营业单位支付的现金净额
	OCPRTIA      float64 //Other cash paid related to investing activities支付其他与投资活动有关的现金
	Subtotal     float64 //现金流出小计
}

func (icof *ICashOutflow) Check() bool {
	total := icof.CalSubtotal()
	return icof.Subtotal == total
}

func (icof *ICashOutflow) CalSubtotal() float64 {
	icof.Subtotal = icof.CAFBFAIAOLTA + icof.CPFI + icof.NCOFOPSAOBU + icof.OCPRTIA
	return icof.Subtotal
}

//FinancingCashFlow 筹资活动现金流量
type FinancingCashFlow struct {
	In    FCashInflow
	Out   FCashOutflow
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
	fcf.Total = fcf.In.Subtotal - fcf.Out.Subtotal
	return fcf.Total
}

//FCashInflow 筹资现金流入
type FCashInflow struct {
	CRFAI      float64 //Cash received from accepting investment吸收投资收到的现金
	Borrowings float64 //取得借款收到的现金
	OCRRTFA    float64 //Other cash received related to financing activities收到其他与筹资活动有关的现金
	Subtotal   float64 //现金流入小计
}

func (fcif *FCashInflow) Check() bool {
	total := fcif.CalSubtotal()
	return fcif.Subtotal == total
}

func (fcif *FCashInflow) CalSubtotal() float64 {
	fcif.Subtotal = fcif.CRFAI + fcif.Borrowings + fcif.OCRRTFA
	return fcif.Subtotal
}

//FCashOutflow 筹资现金流出
type FCashOutflow struct {
	CashPaidForDebt float64 //偿还债务支付的现金
	CPFDPI          float64 //Cash paid for dividend ， profit or interest分配股利、利润或偿付利息支付的现金
	OCPRTFA         float64 //Other cash paid related to financing activities支付其他与筹资活动有关的现金
	Subtotal        float64 //现金流出小计
}

func (fcof FCashOutflow) Check() bool {
	total := fcof.CalSubtotal()
	return fcof.Subtotal == total
}

func (fcof *FCashOutflow) CalSubtotal() float64 {
	fcof.Subtotal = fcof.CashPaidForDebt + fcof.CPFDPI + fcof.OCPRTFA
	return fcof.Subtotal
}

//---------------数值比率分析------------------//
