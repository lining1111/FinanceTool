package BalanceSheet

/**
资产负债表，数据类计算
*/

import (
	"errors"
	"github.com/golang/glog"
	"time"
)

//BalanceSheet 资产负债表 内有恒等式(Assets==LiabilityAndShareHoldersEquity)即资产合计一定等于负债及股东权益合计
type BalanceSheet struct {
	A     Assets
	LEQ   LiabilityAndEquity
	CTime time.Time //表的时间
	Unit  float64   //人民币金额单位
}

func (bs *BalanceSheet) Check() bool {
	//单元素检查
	if !bs.A.Check() {
		glog.Error("BalanceSheet Check A fail,year:%v", bs.CTime.String())
		return false
	}
	if !bs.LEQ.Check() {
		glog.Error("BalanceSheet Check LEQ fail,year:%v", bs.CTime.String())
		return false
	}
	//联合检查
	ret, err := bs.CheckIdentity()
	if err != nil {
		glog.Error("BalanceSheet Check Identity fail,year:%v", bs.CTime.String())
	}
	return ret
}

func (bs *BalanceSheet) CheckIdentity() (bool, error) {
	if bs.A.Total > bs.LEQ.Total {
		return false, errors.New("资产多于负债")
	} else if bs.A.Total < bs.LEQ.Total {
		return false, errors.New("资产少于负债")
	} else {
		return true, nil
	}
}

//Assets 资产
type Assets struct {
	CA    CurrentAssets
	NCA   NonCurrentAssets
	Total float64 //合计
}

func (a *Assets) Check() bool {
	if !a.CA.Check() {
		glog.Error("Assets Check CA fail")
		return false
	}

	if !a.NCA.Check() {
		glog.Error("Assets Check NCA fail")
		return false
	}

	total := a.CalTotal()
	return a.Total == total
}

func (a *Assets) CalTotal() float64 {
	total := a.CA.Total + a.NCA.Total
	if a.Total > 0 {
		if total != a.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, a.Total)
		}
	} else {
		a.Total = total
	}

	return a.Total
}

//CurrentAssets 流动资产
type CurrentAssets struct {
	Cash           float64          //货币资金
	FVTOCI         float64          //交易性金融资产 FVTOCI
	DFI            float64          //衍生金融资产 Derivative Financial Instrument
	BR             float64          //应收票据 bill receivable
	AR             float64          //应收账款 Account Receivable
	ARF            float64          //应收款项融资 accounts receivable financing
	AP             float64          //预付账款 Advance Payment
	OR             OtherReceivables //其他应收款 other receivables
	INV            float64          //存货 Inventory
	ContractAssets float64          //合同资产 Contract Assets
	AHFS           float64          //持有待售资产 Assets Held for Sale
	NCAMWOY        float64          //一年内到期的非流动资产 Non-current Asset Matured within One-Year
	Others         float64          //其他
	Total          float64          //流动资产合计
}

type OtherReceivables struct {
	InterestReceivable  float64 //应收利息
	DividendsReceivable float64 //应收股利
	Others              float64 //其他应收款
	BadDebtReserves     float64 //减：坏账准备
	Total               float64 //合计
}

func (or *OtherReceivables) CalTotal() float64 {

	total := or.InterestReceivable + or.DividendsReceivable + or.Others - or.BadDebtReserves
	if or.Total > 0 {
		if total != or.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, or.Total)
		}
	} else {
		or.Total = total
	}

	return or.Total
}

func (ca *CurrentAssets) Check() bool {
	total := ca.CalTotal()
	return ca.Total == total
}

func (ca *CurrentAssets) CalTotal() float64 {
	total := ca.Cash + ca.FVTOCI + ca.DFI + ca.BR + ca.AR + ca.ARF + ca.AP + ca.OR.CalTotal() +
		ca.INV + ca.ContractAssets + ca.AHFS + +ca.NCAMWOY + ca.Others

	if ca.Total > 0 {
		if total != ca.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, ca.Total)
		}
	} else {
		ca.Total = total
	}

	return ca.Total
}

//NonCurrentAssets 非流动资产
type NonCurrentAssets struct {
	DebtInvestment   float64 //债权投资 Debt investment
	ODebtInvestment  float64 //其他债权投资 Other Debt investment
	LTR              float64 //长期应收款 Long-term Receivable
	LTIOS            float64 //长期股权投资 Long-term investment on stocks
	OEII             float64 //其他权益工具投资 Other equity instrument investment
	ONCFA            float64 //其他非流动金融资产 other non-current financial assets
	REHI             float64 //投资性房地产 Real Estate Held for Investment
	FixedAssets      float64 //固定资产 FixedAssets
	CIP              float64 //在建工程 Construction In Progress
	BiologicalAsset  float64 //生产性生物资产 BiologicalAsset
	OGA              float64 //油气资产 Oil and gas assets
	ROA              float64 //使用权资产 right of assets
	IA               float64 //无形资产 Intangible Assets
	DE               float64 //开发支出 Development Expenditure
	Goodwill         float64 //商誉 Goodwill
	LTPE             float64 //长期待摊费用 Long-term prepaid expenses
	DeferredTaxAsset float64 //递延所得税资产 DeferredTaxAsset
	Others           float64 //其他非流动资产 Other non-current Assets
	Total            float64 //合计
}

func (nca *NonCurrentAssets) Check() bool {
	total := nca.CalTotal()
	return nca.Total == total
}

func (nca *NonCurrentAssets) CalTotal() float64 {
	total := nca.DebtInvestment + nca.LTR + nca.LTIOS + nca.OEII + nca.ONCFA + nca.REHI +
		nca.FixedAssets + nca.CIP + nca.BiologicalAsset + nca.OGA + nca.IA + nca.DE + nca.Goodwill + nca.LTPE + nca.DeferredTaxAsset + nca.Others
	if nca.Total > 0 {
		if total != nca.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, nca.Total)
		}
	} else {
		nca.Total = total
	}

	return nca.Total
}

//LiabilityAndEquity 负债及股东权益
type LiabilityAndEquity struct {
	Li    Liability
	Eq    Equity
	Total float64 //合计
}

func (leq *LiabilityAndEquity) Check() bool {
	if !leq.Li.Check() {
		glog.Error("LiabilityAndEquity Check Li fail")
		return false
	}

	if !leq.Eq.Check() {
		glog.Error("LiabilityAndEquity Check Eq fail")
		return false
	}

	total := leq.CalTotal()
	return leq.Total == total
}

func (leq *LiabilityAndEquity) CalTotal() float64 {
	total := leq.Li.Total + leq.Eq.Total
	if leq.Total > 0 {
		if total != leq.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, leq.Total)
		}
	} else {
		leq.Total = total
	}

	return leq.Total
}

//Liability 负债
type Liability struct {
	CL    CurrentLiability
	NCL   NonCurrentLiability
	Total float64
}

func (li *Liability) Check() bool {
	if !li.CL.Check() {
		glog.Error("Liability Check CL fail")
		return false
	}

	if !li.NCL.Check() {
		glog.Error("Liability Check NCL fail")
		return false
	}

	total := li.CalTotal()
	return li.Total == total
}

func (li *Liability) CalTotal() float64 {
	total := li.CL.Total + li.NCL.Total
	if li.Total > 0 {
		if total != li.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, li.Total)
		}
	} else {
		li.Total = total
	}

	return li.Total
}

//CurrentLiability 流动负债
type CurrentLiability struct {
	ShortTermLoans float64      //短期借款 ShortTermLoans
	TFL            float64      //交易性金融负债 tradable financial liabilities
	DFL            float64      //衍生金融负债 Derivative financial liabilities
	BP             float64      //应付票据 bills payable
	AP             float64      //应付账款 Accounts Payable
	DR             float64      //预收款项 Deposit Received
	CL             float64      //合同负债 Contractual liabilities
	AccruedWages   float64      //应付职工薪酬 AccruedWages
	TaxPayable     float64      //应交税费 TaxPayable
	OP             OtherPayable //其他应付款 Other payable
	LHFS           float64      //持有待售负债 Liabilities held for sale
	NCLMWOY        float64      //一年内到期的非流动负债 Non-current Liabilities Matured within One-Year
	Others         float64      //其他 Others
	Total          float64      //流动负债合计
}

type OtherPayable struct {
	InterestPayable  float64
	DividendsPayable float64
	Others           float64
	Total            float64
}

func (op *OtherPayable) CalTotal() float64 {

	total := op.InterestPayable + op.DividendsPayable + op.Others
	if op.Total > 0 {
		if total != op.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, op.Total)
		}
	} else {
		op.Total = total
	}

	return op.Total
}

func (cl *CurrentLiability) Check() bool {
	total := cl.CalTotal()
	return cl.Total == total
}

func (cl *CurrentLiability) CalTotal() float64 {
	total := cl.ShortTermLoans + cl.TFL + cl.DFL + cl.BP + cl.AP + cl.DR + cl.CL +
		cl.AccruedWages + cl.TaxPayable + cl.OP.CalTotal() + cl.LHFS + cl.NCLMWOY + cl.Others

	if cl.Total > 0 {
		if total != cl.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, cl.Total)
		}
	} else {
		cl.Total = total
	}

	return cl.Total
}

//NonCurrentLiability 非流动负债
type NonCurrentLiability struct {
	LongTermLoans                float64      //长期借款 LongTermLoans
	BondsPayable                 BondsPayable //应付债券 其中分为：优先股和永续债
	LeaseLiabilities             float64      //租赁负债 Lease liabilities
	LongTermPayable              float64      //长期应付款 LongTermPayable
	EstimatedLiabilities         float64      //预计负债 EstimatedLiabilities
	DeferredIncome               float64      //递延收益 Deferred Income
	DeferredIncomeTaxLiabilities float64      //递延所得税负债 DeferredIncomeTaxLiabilities
	Others                       float64      //其他 Others
	Total                        float64      //合计
}

//BondsPayable 应付债券
type BondsPayable struct {
	PFD   float64 //优先股
	PN    float64 //永续债
	Total float64
}

func (bp *BondsPayable) CalTotal() float64 {

	total := bp.PFD + bp.PN
	if bp.Total > 0 {
		if total != bp.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, bp.Total)
		}
	} else {
		bp.Total = total
	}

	return bp.Total
}

func (ncl *NonCurrentLiability) Check() bool {
	total := ncl.CalTotal()
	return ncl.Total == total
}

func (ncl *NonCurrentLiability) CalTotal() float64 {
	total := ncl.LongTermLoans + ncl.BondsPayable.CalTotal() + ncl.LeaseLiabilities +
		ncl.LongTermPayable + ncl.EstimatedLiabilities + ncl.DeferredIncome + ncl.DeferredIncomeTaxLiabilities + ncl.Others
	if ncl.Total > 0 {
		if total != ncl.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, ncl.Total)
		}
	} else {
		ncl.Total = total
	}

	return ncl.Total
}

//Equity 股东权益
type Equity struct {
	PaidInCapital        float64                //实收资本(股本) PaidInCapital
	OEI                  OtherEquityInstruments //其他权益工具 Other equity instruments 其中分为：优先股和永续债
	CapitalReserves      float64                //资本公积 CapitalReserves
	TreasuryStock        float64                //减：库存股 Treasury stock
	OCI                  float64                //其他综合收益 Other comprehensive income
	SpecialReserve       float64                //专项储备 Special reserve
	SurplusReserves      float64                //盈余公积,在盈余公积未达到注册资本的50%时,最低应是税后利润的10%,作用是公司亏损之时弥补其资金上的短缺，增强公司经营的稳定性
	UndistributedProfits float64                //未分配利润 UndistributedProfits
	Total                float64                //合计
}

type OtherEquityInstruments struct {
	PFD   float64 //优先股
	PN    float64 //永续债
	Total float64
}

func (oei *OtherEquityInstruments) CalTotal() float64 {

	total := oei.PFD + oei.PN
	if oei.Total > 0 {
		if total != oei.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, oei.Total)
		}
	} else {
		oei.Total = total
	}

	return oei.Total
}

func (eq *Equity) Check() bool {
	total := eq.CalTotal()
	return eq.Total == total
}

func (eq *Equity) CalTotal() float64 {
	total := eq.PaidInCapital + eq.OEI.CalTotal() + eq.CapitalReserves - eq.TreasuryStock +
		eq.OCI + eq.SpecialReserve + eq.SurplusReserves + eq.UndistributedProfits
	if eq.Total > 0 {
		if total != eq.Total {
			glog.Error("合计错误，计算值%v 获取值%v", total, eq.Total)
		}
	} else {
		eq.Total = total
	}

	return eq.Total
}
