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

func (bs BalanceSheet) CheckIdentity() (bool, error) {
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
	a.Total = a.CA.Total + a.NCA.Total
	return a.Total
}

//CurrentAssets 流动资产
type CurrentAssets struct {
	Cash           float64 //货币资金
	TFA            float64 //交易性金融资产 Transactional financial assets
	DFI            float64 //衍生金融资产 Derivative Financial Instrument
	BR             float64 //应收票据 bill receivable
	AR             float64 //应收账款 Account Receivable
	ARF            float64 //应收款项融资 accounts receivable financing
	AP             float64 //预付账款 Advance Payment
	OR             float64 //其他应收款 other receivables
	INV            float64 //存货 Inventory
	ContractAssets float64 //合同资产 Contract Assets
	AHFS           float64 //持有待售资产 Assets Held for Sale
	NCAMWOY        float64 //一年内到期的非流动资产 Non-current Asset Matured within One-Year
	Others         float64 //其他
	Total          float64 //流动资产合计
}

func (ca *CurrentAssets) Check() bool {
	total := ca.CalTotal()
	return ca.Total == total
}

func (ca *CurrentAssets) CalTotal() float64 {
	ca.Total = ca.Cash + ca.TFA + ca.DFI + ca.BR + ca.AR + ca.ARF + ca.AP + ca.OR +
		ca.INV + ca.ContractAssets + ca.AHFS + +ca.NCAMWOY + ca.Others
	return ca.Total
}

//NonCurrentAssets 非流动资产
type NonCurrentAssets struct {
	DebtInvestment   float64 //债权投资 Debt investment
	ODebtInvestment  float64 //其他债权投资 Other Debt investment
	LTR              float64 //长期应收款 Long-term Receivable
	LTIOS            float64 //长期股权投资 Long-term investment on stocks
	FVTOCI           float64 //其他权益工具投资
	ONCFA            float64 //其他非流动金融资产 other non-current financial assets
	REHI             float64 //Real Estate Held for Investment投资性房地产
	FixedAssets      float64 //固定资产
	CIP              float64 //在建工程 Construction In Progress
	BiologicalAsset  float64 //生产性生物资产
	OGA              float64 //油气资产 Oil and gas assets
	ROA              float64 //使用权资产 right of assets
	IA               float64 //无形资产 Intangible Assets
	DE               float64 //开发支出 Development Expenditure
	Goodwill         float64 //商誉
	LTPE             float64 //长期待摊费用 Long-term prepaid expenses
	DeferredTaxAsset float64 //递延所得税资产
	Others           float64 //其他非流动资产 Other non-current Assets
	Total            float64 //合计
}

func (nca *NonCurrentAssets) Check() bool {
	total := nca.CalTotal()
	return nca.Total == total
}

func (nca *NonCurrentAssets) CalTotal() float64 {
	nca.Total = nca.DebtInvestment + nca.LTR + nca.LTIOS + nca.FVTOCI + nca.ONCFA + nca.REHI +
		nca.FixedAssets + nca.CIP + nca.BiologicalAsset + nca.OGA + nca.IA + nca.DE + nca.Goodwill + nca.LTPE + nca.DeferredTaxAsset + nca.Others
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
	leq.Total = leq.Li.Total + leq.Eq.Total
	return leq.Total
}

//Liability 负债
type Liability struct {
	Cli   CurrentLiability
	Ltli  NonCurrentLiability
	Total float64
}

func (li *Liability) Check() bool {
	if !li.Cli.Check() {
		glog.Error("Liability Check Cli fail")
		return false
	}

	if !li.Ltli.Check() {
		glog.Error("Liability Check Ltli fail")
		return false
	}

	total := li.CalTotal()
	return li.Total == total
}

func (li Liability) CalTotal() float64 {
	li.Total = li.Cli.Total + li.Ltli.Total
	return li.Total
}

//CurrentLiability 流动负债
type CurrentLiability struct {
	ShortTermLoans float64 //短期借款
	TFL            float64 //交易性金融负债 tradable financial liabilities
	DFL            float64 //衍生金融负债 Derivative financial liabilities
	BP             float64 //应付票据 bills payable
	AP             float64 //应付账款 Accounts Payable
	DR             float64 //预收款项 Deposit Received
	CL             float64 //合同负债 Contractual liabilities
	AccruedWages   float64 //应付职工薪酬
	TaxPayable     float64 //应交税费
	OP             float64 //其他应付款 Other payable
	LHFS           float64 //持有待售负债 Liabilities held for sale
	NCLMWOY        float64 //一年内到期的非流动负债 Non-current Liabilities Matured within One-Year
	Others         float64 //其他
	Total          float64 //流动负债合计
}

func (cl *CurrentLiability) Check() bool {
	total := cl.CalTotal()
	return cl.Total == total
}

func (cl *CurrentLiability) CalTotal() float64 {
	cl.Total = cl.ShortTermLoans + cl.TFL + cl.DFL + cl.BP + cl.AP + cl.DR + cl.CL +
		cl.AccruedWages + cl.TaxPayable + cl.OP + cl.LHFS + cl.NCLMWOY + cl.Others
	return cl.Total
}

//NonCurrentLiability 非流动负债
type NonCurrentLiability struct {
	LongTermLoans                float64      //长期借款
	BondsPayable                 BondsPayable //应付债券 其中分为：优先股和永续债
	LeaseLiabilities             float64      //租赁负债 Lease liabilities
	LongTermPayable              float64      //长期应付款
	EstimatedLiabilities         float64      //预计负债
	DeferredIncome               float64      //递延收益 Deferred income
	DeferredIncomeTaxLiabilities float64      //递延所得税负债
	Others                       float64      //其他
	Total                        float64      //合计
}

//BondsPayable 应付债券
type BondsPayable struct {
	PFD   float64 //优先股
	PN    float64 //永续债
	Total float64
}

func (b *BondsPayable) CalTotal() float64 {
	b.Total = b.PFD + b.PN
	return b.Total
}

func (ncl *NonCurrentLiability) Check() bool {
	total := ncl.CalTotal()
	return ncl.Total == total
}

func (ncl *NonCurrentLiability) CalTotal() float64 {
	ncl.Total = ncl.LongTermLoans + ncl.BondsPayable.CalTotal() + ncl.LeaseLiabilities +
		ncl.LongTermPayable + ncl.EstimatedLiabilities + ncl.DeferredIncome + ncl.DeferredIncomeTaxLiabilities + ncl.Others
	return ncl.Total
}

//Equity 股东权益
type Equity struct {
	PaidInCapital        float64                //实收资本(股本)
	OEI                  OtherEquityInstruments //其他权益工具 Other equity instruments 其中分为：优先股和永续债
	CapitalReserves      float64                //资本公积
	TreasuryStock        float64                //减：库存股 Treasury stock
	OCI                  float64                //其他综合收益 Other comprehensive income
	SpecialReserve       float64                //专项储备 Special reserve
	SurplusReserves      float64                //盈余公积,在盈余公积未达到注册资本的50%时,最低应是税后利润的10%,作用是公司亏损之时弥补其资金上的短缺，增强公司经营的稳定性
	UndistributedProfits float64                //未分配利润
	Total                float64                //合计
}

type OtherEquityInstruments struct {
	PFD   float64 //优先股
	PN    float64 //永续债
	Total float64
}

func (oei *OtherEquityInstruments) CalTotal() float64 {
	oei.Total = oei.PFD + oei.PN
	return oei.Total
}

func (eq *Equity) Check() bool {
	total := eq.CalTotal()
	return eq.Total == total
}

func (eq *Equity) CalTotal() float64 {
	eq.Total = eq.PaidInCapital + eq.OEI.CalTotal() + eq.CapitalReserves - eq.TreasuryStock +
		eq.OCI + eq.SpecialReserve + eq.SurplusReserves + eq.UndistributedProfits
	return eq.Total
}
