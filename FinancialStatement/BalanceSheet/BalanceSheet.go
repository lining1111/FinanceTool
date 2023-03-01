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
	Cash              float64 //货币资金
	FVTOCI            float64 //Fair Value through Other Comprehensive Income以公允价值计量且起变动计入其他综合收益
	AccountReceivable float64 //应收账款
	AdvanceMoney      float64 //预付账款
	Inventories       float64 //存货
	NACHSADGCHS       float64 //Non-current assets classified as held for sale and assets in disposal groups classified as held for sale被划分为持有待售的非流动资产及被划分为持有待售的处置组中的资产
	AFSFA             float64 //Available-for-sale financial assets可供出售金融资产
	HTMSecurities     float64 //held-to-maturity securities持有至到期投资
	DeferredExpense   float64 //待摊费用
	Others            float64 //其他
	Total             float64 //流动资产合计
}

func (ca *CurrentAssets) Check() bool {
	total := ca.CalTotal()
	return ca.Total == total
}

func (ca *CurrentAssets) CalTotal() float64 {
	ca.Total = ca.Cash + ca.FVTOCI + ca.AccountReceivable + ca.AdvanceMoney +
		ca.Inventories + ca.NACHSADGCHS + ca.AFSFA + ca.HTMSecurities + ca.DeferredExpense + ca.Others
	return ca.Total
}

//NonCurrentAssets 非流动资产
type NonCurrentAssets struct {
	LongTermInvestment float64 //长期投资
	REHI               float64 //Real Estate Held for Investment投资性房地产
	FixedAssets        float64 //固定资产
	BiologicalAsset    float64 //生物资产
	IAOA               float64 //Intangible Assets And Other Assets无形及其他资产
	DeferredTaxAsset   float64 //递延所得税资产
	Total              float64 //合计
}

func (nca *NonCurrentAssets) Check() bool {
	total := nca.CalTotal()
	return nca.Total == total
}

func (nca *NonCurrentAssets) CalTotal() float64 {
	nca.Total = nca.LongTermInvestment + nca.REHI + nca.FixedAssets + nca.BiologicalAsset +
		nca.IAOA + nca.DeferredTaxAsset
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
	Ltli  LongTermLiability
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
	ShortTermLoans  float64 //短期借款
	FVTPL           float64 //Fair Value through Profit and Loss以公允价值计量且其变动计入当期损益的金融负债
	AccountsPayable float64 //应付账款
	DepositReceived float64 //预收款项
	AccruedWages    float64 //应付职工薪酬
	TaxPayable      float64 //应交税费
	CALHFS          float64 //Classified as liabilities held for sale被划分为持有待售的处置组中的负债
	Others          float64 //其他
	Total           float64 //流动负债合计
}

func (cli *CurrentLiability) Check() bool {
	total := cli.CalTotal()
	return cli.Total == total
}

func (cli *CurrentLiability) CalTotal() float64 {
	cli.Total = cli.ShortTermLoans + cli.FVTPL + cli.AccountsPayable + cli.DepositReceived +
		cli.AccruedWages + cli.TaxPayable + cli.CALHFS + cli.Others
	return cli.Total
}

//LongTermLiability 长期负债
type LongTermLiability struct {
	LongTermLoans                float64 //长期借款
	BondsPayable                 float64 //应付债券
	LongTermPayable              float64 //长期应付款
	EstimatedLiabilities         float64 //预计负债
	DeferredIncomeTaxLiabilities float64 //递延所得税负债
	Others                       float64 //其他
	Total                        float64 //合计
}

func (ltli *LongTermLiability) Check() bool {
	total := ltli.CalTotal()
	return ltli.Total == total
}

func (ltli *LongTermLiability) CalTotal() float64 {
	ltli.Total = ltli.LongTermLoans + ltli.BondsPayable + ltli.LongTermPayable + ltli.EstimatedLiabilities +
		ltli.DeferredIncomeTaxLiabilities + ltli.Others
	return ltli.Total
}

//Equity 股东权益
type Equity struct {
	PaidInCapital        float64 //股本
	CapitalReserves      float64 //资本公积
	SurplusReserves      float64 //盈余公积,在盈余公积未达到注册资本的50%时,最低应是税后利润的10%,作用是公司亏损之时弥补其资金上的短缺，增强公司经营的稳定性
	UndistributedProfits float64 //未分配利润
	Total                float64 //合计
}

func (eq *Equity) Check() bool {
	total := eq.CalTotal()
	return eq.Total == total
}

func (eq *Equity) CalTotal() float64 {
	eq.Total = eq.PaidInCapital + eq.CapitalReserves + eq.SurplusReserves + eq.UndistributedProfits
	return eq.Total
}
